package database

import (
	"FGA_Hacktiv8/jwt/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	HOST     = "localhost"
	USER     = "postgres"
	PASSWORD = "root"
	PORT     = "5432"
	DBNAME   = "simple_api"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", HOST, USER, PASSWORD, DBNAME, PORT)

	db, err = gorm.Open(postgres.Open(config))
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("database connected!")
	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
