package service

import (
	"go-auth/helpers"
	"go-auth/repository"
)

type Service struct {
	Generator helpers.GeneratorInterface
	UserRepo  repository.UserRepositoryInterface
}

func NewService(
	generator helpers.GeneratorInterface,
	userRepo repository.UserRepositoryInterface,
) Service {
	return Service{
		Generator: generator,
		UserRepo:  userRepo,
	}
}
