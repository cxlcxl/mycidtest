package bootstrap

import (
	"log"
	"os"
	"path"
	"time"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/pkg/mylog"
	"xiaoniuds.com/cid/pkg/util/validator"
	"xiaoniuds.com/cid/vars"
)

func init() {
	if err := LoadSysPath(); err != nil {
		log.Fatal("系统路径加载失败", err.Error())
	}

	// 配置文件加载
	var err *errs.MyErr
	vars.Config, err = config.LoadConfig(path.Join(vars.BasePath, "config/config.yaml"))
	if err != nil {
		log.Fatal("配置文件加载失败", err.Error())
	}
	sysLogPath := path.Join(vars.BasePath, "log", "syslog", time.Now().Format("20060102"))
	vars.SysLog = mylog.NewLog(sysLogPath)

	validator.RegisterValidators()
}

func LoadSysPath() (err error) {
	// 通过系统库设置应用根目录变量 BasePath
	vars.BasePath, err = os.Getwd()
	return
}
