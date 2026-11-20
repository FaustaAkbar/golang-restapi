package repository

import (
	"golang-restapi/internal/database"
	"golang-restapi/internal/models"
)

type UserRepository struct{}

func (r UserRepository) FindAll() ([]models.User, error) {

	var users []models.User

	err := database.DB.Find(&users).Error

	return users, err
}

func (r UserRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}
