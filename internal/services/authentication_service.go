package services

import "educore-api/data/request"

type AuthenticationService interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.CreateUsersRequest)
}
