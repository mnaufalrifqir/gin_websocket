package db

import (
	"fmt"
	"log"

	"gin-framework/app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func (config *Config) ConnectDB() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Printf("Database connected!")
}

func InitialMigration() {
	err := DB.AutoMigrate(&models.Student{})

	if err != nil {
		log.Printf("Error when migrating the database: %v", err)
	}

	log.Printf("Database migrated!")
}
