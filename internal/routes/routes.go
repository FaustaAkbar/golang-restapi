package routes

import (
	"github.com/gin-gonic/gin"

	"golang-restapi/internal/handler"
	"golang-restapi/internal/repository"
	"golang-restapi/internal/service"
)

func Setup(router *gin.Engine) {

	userHandler := handler.UserHandler{
		Service: service.UserService{
			Repo: repository.UserRepository{},
		},
	}

	api := router.Group("/api/v1")
	{
		api.GET("/users", userHandler.GetUsers)
		api.POST("/users", userHandler.CreateUser)
	}
}
