package tests

import (
	"rest/app/helpers/env"
	"os"
)

var testApp *App

type App struct {
	Environment string
	Url         string
	Port        string
	Secret      string
}

func InitApp() {
	testApp = &App{
		Environment: os.Getenv("ENV"),
		Secret:      env.MustGet("APP_KEY"),
		Port:        env.MustGet("APP_PORT"),
	}
}

func GetTestApp() *App {
	return testApp
}
