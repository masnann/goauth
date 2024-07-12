package service

import "go-auth/models"

type UserServiceInterface interface {
	Register(req models.UserRegisterRequest) (int64, error)
	FindUserByID(req models.RequestID) (models.UserModels, error)
}
