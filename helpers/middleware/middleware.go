package middleware

import (
	"go-auth/config"
	"go-auth/constants"
	"go-auth/handler"
	"go-auth/helpers"
	"go-auth/models"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// JWTMiddleware validates the JWT token and sets the user in the context
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var result models.Response
		// Extract the JWT token from the request header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			result = helpers.ResponseJSON(false, constants.VALIDATE_ERROR_CODE, "Missing authorization header", nil)
			return c.JSON(http.StatusBadRequest, result)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWTSecret), nil
		})
		if err != nil {
			result = helpers.ResponseJSON(false, constants.UNAUTHORIZED_CODE, err.Error(), nil)
			return c.JSON(http.StatusUnauthorized, result)
		}

		// Extract claims and create a User struct
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["userID"].(float64)
			userRole := claims["role"].(string)
			userEmail := claims["email"].(string)

			user := models.CurrentUserModels{
				ID:    int64(userID),
				Role:  userRole,
				Email: userEmail,
			}

			// Set the user struct in the context
			c.Set("user", user)
		} else {
			result = helpers.ResponseJSON(false, constants.UNAUTHORIZED_CODE, "Invalid token", nil)
			return c.JSON(http.StatusUnauthorized, result)
		}

		return next(c)
	}
}

// Permission Middleware to check user permissions
func PermissionMiddleware(handler handler.Handler, permissionGroup, permissionName string, next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var result models.Response
		currentUser, ok := c.Get("user").(models.CurrentUserModels)
		if !ok {
			result = helpers.ResponseJSON(false, constants.UNAUTHORIZED_CODE, "Failed to get user from context", nil)
			return c.JSON(http.StatusInternalServerError, result)
		}

		// Check user-specific permissions
		userHavePermission, err := handler.PermissionService.IsUserHavePermission(currentUser.ID, permissionGroup, permissionName)
		if err != nil {
			log.Printf("Error checking user-specific permissions for user %d: %v", currentUser.ID, err)
			result = helpers.ResponseJSON(false, constants.SYSTEM_ERROR_CODE, err.Error(), nil)
			return c.JSON(http.StatusInternalServerError, result)
		}

		if userHavePermission {
			return next(c)
		}

		// Check role permissions if no user-specific permission found
		roleHavePermission, err := handler.PermissionService.IsRoleHavePermission(currentUser.ID, permissionGroup, permissionName)
		if err != nil {
			log.Printf("Error checking role permissions for user %d: %v", currentUser.ID, err)
			result = helpers.ResponseJSON(false, constants.SYSTEM_ERROR_CODE, err.Error(), nil)
			return c.JSON(http.StatusInternalServerError, result)
		}

		if roleHavePermission {
			return next(c)
		}

		// If no role permission found, return 403 Forbidden
		log.Printf("Access denied for user %d. Permission Group: %s, Permission Name: %s", currentUser.ID, permissionGroup, permissionName)
		result = helpers.ResponseJSON(false, constants.UNAUTHORIZED_CODE, "Access denied. You don't have permission", nil)
		return c.JSON(http.StatusUnauthorized, result)
	}
}
