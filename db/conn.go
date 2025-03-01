package db

import (
	"fmt"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"log"

)

var(
	DB *gorm.DB
	err error
)

func ConnectDB() *gorm.DB{
	strCon := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	DB, err = gorm.Open(postgres.Open(strCon))

	if err != nil{
		log.Fatal(err.Error())
	}

	return DB
}