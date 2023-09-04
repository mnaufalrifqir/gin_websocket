package db

import (
	"errors"
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

func (config *Config) ConnectDB() *gorm.DB {
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

	return DB
}

func InitialMigration() {
	err := DB.AutoMigrate(&models.Student{})

	if err != nil {
		log.Printf("Error when migrating the database: %v", err)
	}
}

func SeedStudent(db *gorm.DB) (models.Student, error) {
	var student models.Student = models.Student{
		Nim:     "12345",
		Name:    "John Doe",
		Jurusan: "Computer Science",
	}

	result := db.Create(&student)

	if err := result.Error; err != nil {
		return models.Student{}, err
	}

	if err := db.Last(&student).Error; err != nil {
		return models.Student{}, err
	}

	return student, nil
}

func CleanSeeders(db *gorm.DB) error {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

	studentErr := db.Exec("DELETE FROM students").Error

	if studentErr != nil {
		return errors.New("cleaning failed")
	}

	return nil
}
