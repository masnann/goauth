package service

import (
	"go-auth/helpers"
	"go-auth/repository"
)

type Service struct {
	Generator      helpers.GeneratorInterface
	UserRepo       repository.UserRepositoryInterface
	PermissionRepo repository.PermissionRepositoryInterface
}

func NewService(
	generator helpers.GeneratorInterface,
	userRepo repository.UserRepositoryInterface,
	permissionRepo repository.PermissionRepositoryInterface,
) Service {
	return Service{
		Generator:      generator,
		UserRepo:       userRepo,
		PermissionRepo: permissionRepo,
	}
}
