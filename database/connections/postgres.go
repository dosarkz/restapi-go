package connections

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rest/config/initial"
	"log"
	"time"
)

var db *gorm.DB

func Connect(database initial.PostgresDatabase) (*gorm.DB, error) {
	maxRetries := 3
	waitTime := 5

	for i := 1; i <= maxRetries; i++ {
		log.Printf("Opening Connection; Attempt %d of %d...\n", i, maxRetries)
		connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
			database.GetDatabase().Host, database.GetDatabase().Port, database.GetDatabase().DB,
			database.GetDatabase().User, database.GetDatabase().Password)

		db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err != nil {
			if i != maxRetries {
				log.Printf("Cannot open connection (retrying in %ds): %v\n", waitTime, err)
				time.Sleep(time.Duration(waitTime) * time.Second)
			}
			continue
		}
		return db, err
	}
	return nil, errors.New("could not connect to database")
}

func GetDB() *gorm.DB {
	return db
}

func LoadDB(database initial.PostgresDatabase) {
	var err error
	if db != nil {
		return
	}

	db, err = Connect(database)
	if err != nil {
		log.Fatal(err)
	}
}
