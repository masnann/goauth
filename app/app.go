package app

import (
	"go-auth/handler"
	"go-auth/helpers"
	"go-auth/repository"
	userrepository "go-auth/repository/userRepository"
	"go-auth/service"
	userservice "go-auth/service/userService"
)

func SetupApp(repo repository.Repository) handler.Handler {
	generator := helpers.NewGenerator()
	userRepo := userrepository.NewUserRepository(repo)

	service := service.NewService(generator, userRepo)

	userService := userservice.NewUserService(service)

	handler := handler.NewHandler(userService)

	return handler
}
