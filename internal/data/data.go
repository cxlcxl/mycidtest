package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
	"xiaoniuds.com/cid/config"
)

type Data struct {
	DbConnects map[string]*gorm.DB
}

func NewDB(c *config.Config) (data *Data) {
	data = &Data{
		DbConnects: make(map[string]*gorm.DB),
	}
	for _, host := range c.Database.Mysql {
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
		s.SetMaxIdleConns(c.Database.MysqlConnect.MaxIdle)
		s.SetMaxOpenConns(c.Database.MysqlConnect.MaxOpen)
		s.SetConnMaxLifetime(time.Minute * time.Duration(c.Database.MysqlConnect.MaxLife))

		data.DbConnects[host.HostKey] = db
	}

	return
}

func connectDb(host config.MysqlHost) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(host.Dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
