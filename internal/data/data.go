package data

import (
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/pkg/cache"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type QueryBuilder func(db *gorm.DB) *gorm.DB
type Data struct {
	DbConnects map[string]*gorm.DB
}

func NewDB() (data *Data) {
	data = &Data{
		DbConnects: make(map[string]*gorm.DB),
	}
	if vars.Config.Database.Ssh {
		for _, sshHost := range vars.Config.Database.SshHost {
			registerSsh(&Ssh{
				Host:     sshHost.Host,
				User:     sshHost.User,
				Port:     sshHost.Port,
				Password: sshHost.Pass,
			})
		}
	}
	for _, host := range vars.Config.Database.Mysql {
		if host.Dsn == "" {
			continue
		}
		db, err := connectDb(host)
		if err != nil {
			log.Fatalf("[%s]failed opening connection to mysql: %v", host.HostKey, err)
		}

		s, err := db.DB()
		if err != nil {
			log.Fatalf("[%s]failed opening mysql: %v", host.HostKey, err)
		}
		s.SetMaxIdleConns(vars.Config.Database.MysqlConnect.MaxIdle)
		s.SetMaxOpenConns(vars.Config.Database.MysqlConnect.MaxOpen)
		s.SetConnMaxLifetime(time.Minute * time.Duration(vars.Config.Database.MysqlConnect.MaxLife))

		data.DbConnects[host.HostKey] = db
	}

	return
}

func connectDb(host config.MysqlHost) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(host.Dsn), &gorm.Config{
		Logger: logger.New(vars.SysLog, logger.Config{
			SlowThreshold:             1 * time.Second,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Warn,
		}),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// QueryCallFunc 取不到缓存时的回调执行函数
type QueryCallFunc func(interface{}, ...interface{}) *errs.MyErr

// QueryCallWithCache 带缓存的查询
func QueryCallWithCache(
	cache cache.Cache, keySuffix string, value interface{}, exp time.Duration, fn QueryCallFunc, parameters ...interface{},
) (err *errs.MyErr) {
	v := cache.Get(keySuffix)
	if v != "" {
		e := json.Unmarshal([]byte(v), value)
		if e != nil {
			err = errs.Err(errs.SysError, e)
		}
	} else {
		err = fn(value, parameters...)
		if err == nil {
			_ = cache.Set(keySuffix, value, exp)
		}
	}
	return
}
