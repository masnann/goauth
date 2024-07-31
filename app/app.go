package app

import (
	"go-auth/handler"
	"go-auth/helpers"
	"go-auth/repository"
	permissionrepository "go-auth/repository/permissionRepository"
	userrepository "go-auth/repository/userRepository"
	"go-auth/service"
	permissionservice "go-auth/service/permissionService"
	userservice "go-auth/service/userService"
)

func SetupApp(repo repository.Repository) handler.Handler {
	generator := helpers.NewGenerator()
	userRepo := userrepository.NewUserRepository(repo)
	permissionRepo := permissionrepository.NewPermissionRepository(repo)

	service := service.NewService(generator, userRepo, permissionRepo)

	userService := userservice.NewUserService(service)
	permissionService := permissionservice.NewPermissionService(service)

	handler := handler.NewHandler(userService, permissionService)

	return handler
}
