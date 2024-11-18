package services

import (
	"context"
	"educore-api/internal/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func (u *UserService) CreateUser(user *models.User) error {
	return nil
}

func (u *UserService) GetUser(name *string) (*models.User, error) {
	return nil, nil
}

func (u *UserService) GetAllUsers() ([]*models.User, error) {
	return nil, nil
}

func (u *UserService) UpdateUser(user *models.User) error {
	return nil
}

func (u *UserService) DeleteUser(name *string) error {
	return nil
}
