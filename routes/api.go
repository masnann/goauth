package routes

import (
	"go-auth/handler"
	userhandler "go-auth/handler/userHandler"
	"go-auth/helpers"
	"net/http"

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
	private.Use(helpers.JWTMiddleware)

	userGroup := private.Group("/user")
	userGroup.POST("/findbyid", userHandler.FindUserByID)
	userGroup.POST("/delete", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, "Delete user functionality is not implemented yet.")
	})
}
