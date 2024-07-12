package helpers

import (
	"fmt"
	"go-auth/models"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func ResponseJSON(success bool, code, message string, result interface{}) models.Response {
	response := models.Response{
		StatusCode:       code,
		Success:          success,
		Message:          message,
		ResponseDateTime: time.Now(),
		Result:           result,
	}

	return response
}

func TimeStampNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}

// Validate
var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("noSpace", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		return !regexp.MustCompile(`\s`).MatchString(password)
	})
}

func ValidateStruct(ctx echo.Context, s interface{}) error {
	// Bind the request body to the struct
	if err := ctx.Bind(s); err != nil {
		return fmt.Errorf("Invalid request body")
	}

	// Validate the struct
	err := validate.Struct(s)
	if err != nil {
		var customErrors []string

		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				customErrors = append(customErrors, fmt.Sprintf("Field '%s' is required", err.Field()))
			case "min":
				customErrors = append(customErrors, fmt.Sprintf("Field '%s' must be at least %s characters long", err.Field(), err.Param()))
			case "email":
				customErrors = append(customErrors, fmt.Sprintf("Field '%s' must be a valid email address", err.Field()))
			case "noSpace":
				customErrors = append(customErrors, fmt.Sprintf("Field '%s' cannot contain spaces", err.Field()))
			case "alphanum":
				customErrors = append(customErrors, fmt.Sprintf("Field '%s' must be alphanumeric", err.Field()))
			case "max":
				customErrors = append(customErrors, fmt.Sprintf("Field '%s' cannot be longer than %s characters", err.Field(), err.Param()))
			default:
				customErrors = append(customErrors, fmt.Sprintf("Field '%s' validation failed with tag '%s'", err.Field(), err.Tag()))
			}
		}

		return fmt.Errorf("Validation error: %s", strings.Join(customErrors, "; "))
	}
	return nil
}
