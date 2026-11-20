package main

import (
	"github.com/gin-gonic/gin"

	"golang-restapi/internal/config"
	"golang-restapi/internal/database"
	"golang-restapi/internal/models"
	"golang-restapi/internal/routes"
)

func main() {

	config.LoadEnv()

	database.ConnectMySQL()
	database.ConnectRedis()

	database.DB.AutoMigrate(&models.User{})

	router := gin.Default()

	routes.Setup(router)

	router.Run(":8080")
}
