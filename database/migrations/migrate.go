package migrations

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"rest/config/initial"
	"log"
)

func Migrate(database initial.PostgresDatabase) {
	cnf := database.GetDatabase()
	m, err := migrate.New(
		"file://database/migrations/sources",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			cnf.User, cnf.Password, cnf.Host, cnf.Port, cnf.DB),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Println(err)
	}
}
