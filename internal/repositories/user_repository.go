package repository

import (
	"educore-api/internal/models"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByName(name string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Update(user *models.User) error
	Delete(name string) error
}
