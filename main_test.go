package main

import (
	"gorm.io/gorm"
	"rest/app/helpers/env"
	"rest/config/initial"
	"rest/tests"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	dbConf := new(tests.Database)
	env.HasEnvironment()
	tests.InitApp()
	tests.LoadDB(dbConf)
	tests.Migrate(dbConf)
	initial.LoadRepos(tests.GetDb())
	m.Run()
	defer func(database *gorm.DB) {
		db, err := database.DB()
		if err != nil {
			log.Fatal(err)
		}
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}(tests.GetDb())
}
