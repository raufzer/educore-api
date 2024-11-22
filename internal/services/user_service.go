package services

import (
	"educore-api/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUser(name *string) (*models.User, error)
	GetUserID(Id *primitive.ObjectID) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	UpdateUser(users *models.User) error
	DeleteUser(name *string) error
}
