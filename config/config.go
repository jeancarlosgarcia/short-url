package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/gommon/log"
)

type Config struct {
	ServerHost string      `required:"true" split_words:"true"`
	ServerPort int         `required:"true" split_words:"true"`
	Redis      RedisConfig `required:"true"`
	DB         DBConfig    `required:"true"`
}

type RedisConfig struct {
	Host string `required:"true"`
	Port string `required:"true"`
}

type DBConfig struct {
	Host     string `required:"true"`
	Database string `required:"true"`
	User     string `required:"true"`
	Password string `required:"true"`
}

var once sync.Once
var c Config

func Environments() Config {
	once.Do(func() {
		if err := envconfig.Process("", &c); err != nil {
			log.Panicf("Error parsing environment vars %#v", err)
		}
	})

	return c
}