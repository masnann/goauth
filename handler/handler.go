package handler

import "go-auth/service"

type Handler struct {
	UserService       service.UserServiceInterface
	PermissionService service.PermissionServiceInterface
}

func NewHandler(
	userService service.UserServiceInterface,
	permissionService service.PermissionServiceInterface,
) Handler {
	return Handler{
		UserService:       userService,
		PermissionService: permissionService,
	}
}
