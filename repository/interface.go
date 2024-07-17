package repository

import "go-auth/models"

type UserRepositoryInterface interface {
	Register(req models.UserModels) (int64, error)
	FindUserByID(id int64) (models.UserModels, error)
	Login(email string) (models.UserModels, error)
}
