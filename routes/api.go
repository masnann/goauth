package routes

import (
	"go-auth/handler"
	userhandler "go-auth/handler/userHandler"

	"github.com/labstack/echo/v4"
)

func ApiRoutes(e *echo.Echo, handler handler.Handler) {
	public := e.Group("/api/v1/public")

	userHandler := userhandler.NewUserHandler(handler)
	userGroup := public.Group("/user")
	userGroup.POST("/register", userHandler.Register)
	userGroup.POST("/findbyid", userHandler.FindUserByID)

}
