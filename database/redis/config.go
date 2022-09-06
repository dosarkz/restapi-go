package redis

import (
	"rest/app/helpers/env"
	"os"
)

type Config struct {
	Host      string
	Port      string
	Password  string
	CacheTime string
}

type Redis interface {
	Get() *Config
}

func (d *Config) Get() *Config {
	return &Config{
		Host:     env.MustGet("REDIS_HOST"),
		Port:     env.MustGet("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}
}
