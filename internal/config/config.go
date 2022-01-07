package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"sync"
)

type Config struct {
	Environment string    `yaml:"environment"`
	AppName     string    `yaml:"appName"`
	Server      ServerCfg `yaml:"application"`
	Db          DbCfg     `yaml:"storage"`
	Log         LogCfg    `yaml:"log"`
	Auth        AuthCfg   `yaml:"auth"`
}

type ServerCfg struct {
	Port     string `yaml:"port"`
	BasePath string `yaml:"basePath"`
	TimeOut  int    `yaml:"timeOut"`
}

type DbCfg struct {
	PostgresConnection string `yaml:"postgresConnection"`
	MaxIdleConnections int    `yaml:"maxIdleConnections"`
	MaxOpenConnections int    `yaml:"maxOpenConnections"`
}

type LogCfg struct {
	Title     string `yaml:"title"`
	Type      string `yaml:"type"`
	Level     string `yaml:"level"`
	Formatter string `yaml:"formatter"`
}

type AuthCfg struct {
	Secret string `yaml:"secret"`
}

var cfg *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		cfg = &Config{}
		if err := cleanenv.ReadConfig("config.yml", cfg); err != nil {
			help, _ := cleanenv.GetDescription(cfg, nil)
			logrus.Info(help)
			logrus.Fatal(err)
		}
	})

	return cfg
}
