package main

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"rest/config/initial"
	"rest/database/seeders"
	"rest/router"
	"rest/server"
)

func main() {
	server.LoadDependencies()
	router.Init()
	seeders.Run(initial.GetRepositories().UserRepo)
	server.Run()
}
