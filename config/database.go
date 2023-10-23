package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	dsn := "host=trumpet.db.elephantsql.com user=knxjvuxb password=JyXrohYCO_a3FaMAJRiC3Y6pib9rQ0XG dbname=knxjvuxb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Default().Output(1, "Connected")
	return db
}
