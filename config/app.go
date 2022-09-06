package config

import (
	"rest/app/helpers/env"
	"os"
)

type App struct {
	Environment      string
	Url              string
	Port             string
	Secret           string
	Storage          string
	IPRequestLimiter RequestLimiterConfig
}

type RequestLimiterConfig struct {
	RequestPerSecond int
	LimiterCapacity  int
}

var app *App

func Init() {
	app = &App{
		Environment: os.Getenv("ENV"),
		Secret:      env.MustGet("APP_KEY"),
		Port:        env.MustGet("APP_PORT"),
		Url:         env.MustGet("APP_URL"),
		Storage:     env.MustGet("APP_MEDIA_STORAGE"),
		IPRequestLimiter: RequestLimiterConfig{
			RequestPerSecond: env.MustGetInt("APP_IP_REQUEST_PER_SECOND"),
			LimiterCapacity:  env.MustGetInt("APP_REQUEST_LIMITER_CAPACITY"),
		},
	}
}

func GetApp() *App {
	return app
}

func (a App) GetUrlAddr() string {
	return a.Url + ":" + a.Port
}
