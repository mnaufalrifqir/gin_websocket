package main

import (
	"gin-framework/app/db"
	"gin-framework/app/router"
	"gin-framework/app/utils"
)

// @title CRUD Students API
// @version 1.0
// @description API for CRUD Students
// @host localhost:8000
// @BasePath /
func main() {
	config := db.Config{
		DB_Username: utils.GetConfig("DB_USERNAME"),
		DB_Password: utils.GetConfig("DB_PASSWORD"),
		DB_Host:     utils.GetConfig("DB_HOST"),
		DB_Port:     utils.GetConfig("DB_PORT"),
		DB_Name:     utils.GetConfig("DB_NAME"),
	}

	config.ConnectDB()
	db.InitialMigration()

	route := router.SetupRouter()

	route.Run(":" + utils.GetConfig("PORT"))
}
