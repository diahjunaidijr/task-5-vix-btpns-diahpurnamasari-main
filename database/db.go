package database

import (
	"fmt"
	"log"
	"task-5-vix-btpns-Diah_Purnama_Sari/helpers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// on package go.dev
func ConnDb() *gorm.DB {
	connectString := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v", helpers.GetEnv("DB_HOST"), helpers.GetEnv("DB_USER"), helpers.GetEnv("DB_PASSWORD"), helpers.GetEnv("DB_NAME"), helpers.GetEnv("DB_PORT"), helpers.GetEnv("DB_SSL_MODE"), helpers.GetEnv("TIMEZONE"))

	db, err := gorm.Open(postgres.Open(connectString), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
