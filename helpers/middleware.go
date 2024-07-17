package helpers

import (
	"go-auth/config"
	"go-auth/constants"
	"go-auth/models"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// Role defines the user role
type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

// JWTMiddleware validates the JWT token and sets the user role in the context
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var result models.Response
		// Extract the JWT token from the request header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			result = ResponseJSON(false, constants.VALIDATE_ERROR_CODE, "Missing authorization header", nil)
			return c.JSON(http.StatusBadRequest, result)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWTSecret), nil
		})
		if err != nil {
			result = ResponseJSON(false, constants.UNAUTHORIZED_CODE, err.Error(), nil)
			return c.JSON(http.StatusUnauthorized, result)
		}

		// Extract claims and set user role in the context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userRole := Role(claims["role"].(string))
			c.Set("role", userRole)
		} else {
			result = ResponseJSON(false, constants.UNAUTHORIZED_CODE, "Invalid token", nil)
			return c.JSON(http.StatusUnauthorized, result)
		}

		return next(c)
	}
}

// AdminMiddleware restricts access to routes only for admin users
func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var result models.Response
		// Get the user role from the context
		userRole, ok := c.Get("role").(Role)
		if !ok || userRole != RoleAdmin {
			result = ResponseJSON(false, constants.FORBIDDEN_CODE, "Access denied. You don't have permission", nil)
			return c.JSON(http.StatusForbidden, result)
		}

		return next(c)
	}
}

// UserMiddleware restricts access to routes only for regular users
func UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var result models.Response
		// Get the user role from the context
		userRole, ok := c.Get("role").(Role)
		if !ok || userRole != RoleUser {
			result = ResponseJSON(false, constants.FORBIDDEN_CODE, "Access denied. You don't have permission", nil)
			return c.JSON(http.StatusForbidden, result)
		}

		return next(c)
	}
}
