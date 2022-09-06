package initial

import "os"

type Database struct {
	Host     string
	Port     string
	User     string
	DB       string
	Password string
	Timezone string
}

type PostgresDatabase interface {
	GetDatabase() *Database
}

func (d *Database) GetDatabase() *Database {
	return &Database{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		User:     os.Getenv("DATABASE_USER"),
		DB:       os.Getenv("DATABASE_DB"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Timezone: os.Getenv("DATABASE_TIMEZONE"),
	}
}
