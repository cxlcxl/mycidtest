package config

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"xiaoniuds.com/cid/pkg/errs"
)

type Config struct {
	Port       int        `yaml:"port"`
	Mode       string     `yaml:"mode"`
	MainDomain string     `yaml:"main_domain"`
	Auth       AuthModule `yaml:"auth"`
	Database   Database   `yaml:"database"`
	DuoIds     []int64    `yaml:"duo_ids"`
}

type Database struct {
	Mysql        []MysqlHost   `yaml:"mysql"`
	MysqlConnect MysqlConnect  `yaml:"mysql_connect"`
	Redis        Redis         `yaml:"redis"`
	Ssh          bool          `yaml:"ssh"`
	SshHost      []ConnectHost `yaml:"ssh_host"`
}

type MysqlConnect struct {
	MaxIdle int `yaml:"max_idle"`
	MaxOpen int `yaml:"max_open"`
	MaxLife int `yaml:"max_life"`
}

type Redis struct {
	Common RedisHost `yaml:"common"`
}

type AuthModule struct {
	Login       Auth         `yaml:"login"`
	OpenApi     Auth         `yaml:"open_api"`
	OpenApiApps []OpenApiApp `yaml:"open_api_apps"`
}

type MysqlHost struct {
	HostKey string `yaml:"host_key"`
	Dsn     string `yaml:"dsn"`
}
type RedisHost struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Pass string `yaml:"pass"`
	Db   string `yaml:"db"`
}
type ConnectHost struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Pass string `yaml:"pass"`
	User string `yaml:"user"`
}
type Auth struct {
	Id      string `yaml:"id"`
	SignKey string `yaml:"sign_key"`
	Exp     int64  `yaml:"exp"`
}
type OpenApiApp struct {
	AppId      string `yaml:"app_id"`
	AppSecret  string `yaml:"app_secret"`
	MainUserId int64  `yaml:"main_user_id"`
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
