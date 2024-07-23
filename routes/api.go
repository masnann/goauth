package routes

import (
	"go-auth/handler"
	userhandler "go-auth/handler/userHandler"
	"go-auth/helpers/middleware"

	"github.com/labstack/echo/v4"
)

func ApiRoutes(e *echo.Echo, handler handler.Handler) {

	public := e.Group("/api/v1/public")
	userHandler := userhandler.NewUserHandler(handler)

	authGroup := public.Group("/auth")
	authGroup.POST("/register", userHandler.Register)
	authGroup.POST("/login", userHandler.Login)
	authGroup.POST("/token/refresh", userHandler.RefreshToken)

	private := e.Group("/api/v1/private")
	private.Use(middleware.JWTMiddleware)

	userGroup := private.Group("/user")
	userGroup.POST("/findbyid", userHandler.FindUserByID)
	userGroup.POST("/delete", middleware.PermissionMiddleware(handler, "USER", "DELETE")(userHandler.DeleteUser))
}
