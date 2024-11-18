package services

import (
	"context"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func (u *UserServiceImpl) CreateUser(user *User) error {
	return nil
}

func (u *UserServiceImpl) GetUser(name *string) (*model.User, error) {
	return nil, nil
}
