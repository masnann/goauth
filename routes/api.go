package routes

import (
	"go-auth/handler"
	userhandler "go-auth/handler/userHandler"
	"go-auth/helpers"

	"github.com/labstack/echo/v4"
)

func ApiRoutes(e *echo.Echo, handler handler.Handler) {

	e.Use(helpers.JWTMiddleware)

	public := e.Group("/api/v1/public")

	userHandler := userhandler.NewUserHandler(handler)
	userGroup := public.Group("/user")
	userGroup.POST("/register", userHandler.Register)
	userGroup.POST("/findbyid", helpers.AdminMiddleware(userHandler.FindUserByID))
	userGroup.POST("/login", userHandler.Login)
	userGroup.POST("/token/refresh", userHandler.RefreshToken)

}
