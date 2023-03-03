package util

import (
	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitializeGorm() *gorm.DB {
	dbLink := os.Getenv("DATABASE_URL")

	postgreDriver, err := sql.Open("postgres", dbLink)
	if err != nil {
		log.Fatal(err.Error())
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: postgreDriver,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
