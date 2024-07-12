package handler

import "go-auth/service"

type Handler struct {
	UserService service.UserServiceInterface
}

func NewHandler(
	userService service.UserServiceInterface,
) Handler {
	return Handler{
		UserService: userService,
	}
}
