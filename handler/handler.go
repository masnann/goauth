package handler

import (
	"go-auth/helpers"
	"go-auth/service"
)

type Handler struct {
	UserService       service.UserServiceInterface
	PermissionService service.PermissionServiceInterface
	Generator         helpers.GeneratorInterface
}

func NewHandler(
	userService service.UserServiceInterface,
	permissionService service.PermissionServiceInterface,
	generator helpers.GeneratorInterface,
) Handler {
	return Handler{
		UserService:       userService,
		PermissionService: permissionService,
		Generator:         generator,
	}
}
