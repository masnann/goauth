package routes

import (
	"go-auth/handler"
	permissionhandler "go-auth/handler/permissionHandler"
	userhandler "go-auth/handler/userHandler"
	"go-auth/helpers/middleware"

	"github.com/labstack/echo/v4"
)

func ApiRoutes(e *echo.Echo, handler handler.Handler) {

	public := e.Group("/api/v1/public")
	userHandler := userhandler.NewUserHandler(handler)
	permissionHandler := permissionhandler.NewPermissionHandler(handler)

	authGroup := public.Group("/auth")
	authGroup.POST("/register", userHandler.Register)
	authGroup.POST("/login", userHandler.Login)
	authGroup.POST("/token/refresh", userHandler.RefreshToken)
	authGroup.POST("/generate-otp", userHandler.GenerateOTP)
	authGroup.POST("/validate-otp", userHandler.VerifyOTP)

	private := e.Group("/api/v1/private")
	private.Use(middleware.JWTMiddleware)

	userGroup := private.Group("/user")
	userGroup.POST("/findbyid", userHandler.FindUserByID)
	userGroup.POST("/delete", middleware.PermissionMiddleware(handler, "USER", "DELETE", userHandler.DeleteUser))

	permissionGroup := private.Group("/permission")
	permissionGroup.POST("/create", middleware.PermissionMiddleware(handler, "PERMISSION", "CREATE", permissionHandler.CreatePermission))
	permissionGroup.POST("/role/create", middleware.PermissionMiddleware(handler, "PERMISSION", "CREATE", permissionHandler.CreateRole))
	permissionGroup.POST("/list", permissionHandler.FindListPermission)
	permissionGroup.POST("/role/list", permissionHandler.FindListRole)
}
