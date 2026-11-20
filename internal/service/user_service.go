package service

import (
	"encoding/json"
	"time"

	"golang-restapi/internal/database"
	"golang-restapi/internal/models"
	"golang-restapi/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s UserService) GetUsers() ([]models.User, error) {

	cacheKey := "users"

	val, err := database.RedisClient.Get(
		database.Ctx,
		cacheKey,
	).Result()

	if err == nil {

		var users []models.User

		json.Unmarshal([]byte(val), &users)

		return users, nil
	}

	users, err := s.Repo.FindAll()

	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(users)

	database.RedisClient.Set(
		database.Ctx,
		cacheKey,
		data,
		5*time.Minute,
	)

	return users, nil
}

func (s UserService) CreateUser(user *models.User) error {

	err := s.Repo.Create(user)

	if err != nil {
		return err
	}

	database.RedisClient.Del(
		database.Ctx,
		"users",
	)

	return nil
}
