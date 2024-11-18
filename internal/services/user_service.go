package services

import "educore-api/internal/models"

type UserService interface {
	CreateUser(user *models.User) error
	GetUser(name *string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(name *string) error
}
