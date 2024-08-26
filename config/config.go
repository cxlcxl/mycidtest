package config

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"xiaoniuds.com/cid/pkg/errs"
)

type Config struct {
	Port     int      `yaml:"port"`
	Mode     string   `yaml:"mode"`
	Database Database `yaml:"database"`
}

type Database struct {
	Mysql []MysqlHost `yaml:"mysql"`
	Redis Redis       `yaml:"redis"`
}

type Redis struct {
	Common RedisHost `yaml:"common"`
}

type MysqlHost struct {
	HostKey string `yaml:"host_key"`
	Dsn     string `yaml:"dsn"`
}
type RedisHost struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Pass string `yaml:"password"`
	Db   string `yaml:"db"`
}

func LoadConfig(configPath string) (c *Config, err *errs.MyErr) {
	f, e := os.Open(configPath)
	if e != nil {
		return nil, errs.Err(errs.ConfigLoadError, e)
	}
	defer f.Close()

	all, e := io.ReadAll(f)
	if e != nil {
		return nil, errs.Err(errs.ConfigLoadError, e)
	}

	e = yaml.Unmarshal(all, &c)
	if e != nil {
		return nil, errs.Err(errs.ConfigLoadError, e)
	}

	return c, nil
}
