package main

import (
	"flag"
	"log"
	"os"
	"path"
	"time"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/server"
	"xiaoniuds.com/cid/pkg/mylog"
	"xiaoniuds.com/cid/vars"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Version string
	conf    string
)

func init() {
	if err := loadSysPath(); err != nil {
		log.Fatal("系统路径加载失败", err.Error())
	}
	log.Println("系统路径加载成功", vars.BasePath)
	flag.StringVar(&conf, "conf", path.Join(vars.BasePath, "config/config.yaml"), "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()

	// 配置文件加载
	c, err := config.LoadConfig(conf)
	if err != nil {
		log.Fatal("配置文件加载失败", err.Error())
	}
	sysLogPath := path.Join(vars.BasePath, "log", "syslog", time.Now().Format("20060102"))
	vars.SysLog = mylog.NewLog(sysLogPath)

	_ = server.NewServer(c).Run()
}

func loadSysPath() (err error) {
	// 通过系统库设置应用根目录变量 BasePath
	vars.BasePath, err = os.Getwd()
	return
}
