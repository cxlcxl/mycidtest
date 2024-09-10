package mylog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"path"
	"time"
	"xiaoniuds.com/cid/pkg/util"
	"xiaoniuds.com/cid/vars"
)

type Log struct {
	*logrus.Logger
}

func NewLog() *Log {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.DateTime,
	})

	sysLogPath := path.Join(vars.BasePath, "log", "syslog", time.Now().Format("2006010215"))
	util.Mkdir(sysLogPath, 1)
	l := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", sysLogPath, "syslog.log"),
		MaxSize:    500,
		MaxBackups: 10,
		MaxAge:     15,
		Compress:   true,
	}
	log.SetOutput(l)

	return &Log{log}
}
