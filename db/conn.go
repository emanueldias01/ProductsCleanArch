package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"log"

)

var(
	DB *gorm.DB
	err error
)

func ConnectDB(){
	strCon := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(strCon))

	if err != nil{
		log.Fatal(err.Error())
	}
}