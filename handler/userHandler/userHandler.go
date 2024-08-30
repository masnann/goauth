package userhandler

import (
	"go-auth/constants"
	"go-auth/handler"
	"go-auth/helpers"
	logerror "go-auth/helpers/logError"
	"go-auth/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	handler handler.Handler
}

func NewUserHandler(handler handler.Handler) UserHandler {
	return UserHandler{
		handler: handler,
	}
}

func (h UserHandler) Register(ctx echo.Context) error {
	var result models.Response

	req := new(models.UserRegisterRequest)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATION_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var (
		serviceName string = "Register"
		errMsg      string = "An unexpected error occurred"
		userEmail   string = req.Email
	)

	userID, err := h.handler.UserService.Register(*req)
	if err != nil {
		if myErr, ok := err.(logerror.LogError); ok {
			log.Printf("Error finding user: %v", err)
			return h.handler.Generator.HandlerErr(ctx, myErr.ErrorType, serviceName, userEmail, myErr.Msg, err)
		}
		log.Printf("Error: Unexpected error finding user (%v)", err)
		return h.handler.Generator.HandlerErr(ctx, "unknown", serviceName, userEmail, errMsg, err)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, userID)
	return ctx.JSON(http.StatusCreated, result)

}

func (h UserHandler) FindUserByID(ctx echo.Context) error {
	var result models.Response

	currentUser, err := helpers.ValidateUserAndRole(ctx, []string{"Admin"})
	if err != nil {
		log.Printf("Error Permission: %v", err)
		result := helpers.ResponseJSON(false, constants.FORBIDDEN_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusForbidden, result)
	}

	req := new(models.RequestID)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATION_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var (
		serviceName string = "FindUserByID"
		errMsg      string = "An unexpected error occurred"
		userEmail   string = currentUser.Email
	)

	user, err := h.handler.UserService.FindUserByID(*req)
	if err != nil {
		if myErr, ok := err.(logerror.LogError); ok {
			log.Printf("Error finding user: %v", err)
			return h.handler.Generator.HandlerErr(ctx, myErr.ErrorType, serviceName, userEmail, myErr.Msg, err)
		}
		log.Printf("Error: Unexpected error finding user (%v)", err)
		return h.handler.Generator.HandlerErr(ctx, "unknown", serviceName, userEmail, errMsg, err)
	}

	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, user)
	return ctx.JSON(http.StatusOK, result)
}

func (h UserHandler) Login(ctx echo.Context) error {
	var result models.Response

	req := new(models.UserLoginRequest)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATION_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	token, err := h.handler.UserService.Login(*req)
	if err != nil {
		log.Printf("Error Login: %v", err)
		result = helpers.ResponseJSON(false, constants.INTERNAL_SERVER_ERROR, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, token)
	return ctx.JSON(http.StatusOK, result)
}

func (h UserHandler) RefreshToken(ctx echo.Context) error {
	var result models.Response
	req := new(models.RefreshTokenRequest)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATION_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	token, err := h.handler.UserService.RefreshToken(req.RefreshToken)
	if err != nil {
		log.Printf("Error RefreshToken: %v", err)
		result = helpers.ResponseJSON(false, constants.INTERNAL_SERVER_ERROR, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, token)
	return ctx.JSON(http.StatusOK, result)
}

func (h UserHandler) DeleteUser(ctx echo.Context) error {
	var result models.Response

	response := "Success"

	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, response)
	return ctx.JSON(http.StatusOK, result)
}

func (h UserHandler) GenerateOTP(ctx echo.Context) error {
	var result models.Response

	req := new(models.UserGenerateOTPRequest)
	otp, err := h.handler.UserService.GenerateOTP(*req)
	if err != nil {
		log.Printf("Error GenerateOTP: %v", err)
		result = helpers.ResponseJSON(false, constants.INTERNAL_SERVER_ERROR, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, otp)
	return ctx.JSON(http.StatusOK, result)
}

func (h UserHandler) VerifyOTP(ctx echo.Context) error {
	var result models.Response

	req := new(models.UserValidateOtpRequest)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATION_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	verified, err := h.handler.UserService.ValidateOtp(*req)
	if err != nil {
		log.Printf("Error VerifyOTP: %v", err)
		result = helpers.ResponseJSON(false, constants.INTERNAL_SERVER_ERROR, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, verified)
	return ctx.JSON(http.StatusOK, result)
}
