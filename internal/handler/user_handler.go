package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"golang-restapi/internal/models"
	"golang-restapi/internal/service"
)

type UserHandler struct {
	Service service.UserService
}

func (h UserHandler) GetUsers(c *gin.Context) {

	users, err := h.Service.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h UserHandler) CreateUser(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateUser(&user); err != nil {

		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
