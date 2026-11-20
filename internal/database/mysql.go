package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"golang-restapi/internal/config"
)

var DB *gorm.DB

func ConnectMySQL() {

	var err error

	DB, err = gorm.Open(
		mysql.Open(config.GetEnv("MYSQL_DSN")),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal(err)
	}
}
