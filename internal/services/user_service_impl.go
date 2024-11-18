package services

import (
	"educore-api/internal/models"
	"educore-api/internal/repositories"
)

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo: userRepo,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	return u.userRepo.Create(user)
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	return u.userRepo.GetByName(*name)
}

func (u *UserServiceImpl) GetAllUsers() ([]*models.User, error) {
	return u.userRepo.GetAll()
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	return u.userRepo.Update(user)
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	return u.userRepo.Delete(*name)
}
