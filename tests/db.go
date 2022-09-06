package tests

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

type Database struct {
	Host     string
	Port     string
	User     string
	DB       string
	Password string
	Timezone string
}

func ConnectDb(database PostgresDatabase) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		database.GetDatabase().Host, database.GetDatabase().Port, database.GetDatabase().DB, database.GetDatabase().User,
		database.GetDatabase().Password)
	return gorm.Open(postgres.Open(connStr), &gorm.Config{})
}

func LoadDB(database PostgresDatabase) {
	var err error
	db, err = ConnectDb(database)
	if err != nil {
		log.Fatal(err)
	}
}

type PostgresDatabase interface {
	GetDatabase() *Database
}

func CloseDb(database *gorm.DB) error {
	db, err := database.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (d *Database) GetDatabase() *Database {
	return &Database{
		Host: os.Getenv("DATABASE_HOST"),
		//Host: "localhost",
		Port: os.Getenv("DATABASE_PORT"),
		//Port: "5411",
		User:     os.Getenv("DATABASE_USER"),
		DB:       os.Getenv("DATABASE_DB_TEST"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		//Password: "secret",
		Timezone: os.Getenv("DATABASE_TIMEZONE"),
	}
}

func GetDb() *gorm.DB {
	return db
}

func Migrate(database PostgresDatabase) {
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
