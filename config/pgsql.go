package config

import (
	"main/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB
var dns string = "host=localhost user=root password=root dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func Connect() error {
	var err error

	Database, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.Dog{})

	return nil
}
