package permissionhandler

import (
	"go-auth/constants"
	"go-auth/handler"
	"go-auth/helpers"
	"go-auth/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PermissionHandler struct {
	handler handler.Handler
}

func NewPermissionHandler(handler handler.Handler) PermissionHandler {
	return PermissionHandler{
		handler: handler,
	}
}

func (h PermissionHandler) CreateRole(ctx echo.Context) error {
	var result models.Response

	req := new(models.RoleCreateRequest)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATION_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	roleID, err := h.handler.PermissionService.CreateRole(*req)
	if err != nil {
		log.Printf("Error CreateRole: %v", err)
		result = helpers.ResponseJSON(false, constants.INTERNAL_SERVER_ERROR, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, roleID)
	return ctx.JSON(http.StatusCreated, result)
}

func (h PermissionHandler) CreatePermission(ctx echo.Context) error {
	var result models.Response

	req := new(models.PermissionCreateRequest)
	if err := helpers.ValidateStruct(ctx, req); err != nil {
		log.Printf("Error Failed to validate request: %v", err)
		result = helpers.ResponseJSON(false, constants.VALIDATION_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	permissionID, err := h.handler.PermissionService.CreatePermission(*req)
	if err != nil {
		log.Printf("Error CreatePermission: %v", err)
		result = helpers.ResponseJSON(false, constants.INTERNAL_SERVER_ERROR, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, permissionID)
	return ctx.JSON(http.StatusCreated, result)
}

func (h PermissionHandler) FindListRole(ctx echo.Context) error {
	var result models.Response

	_, err := helpers.ValidateUserAndRole(ctx, []string{"Admin"})
	if err != nil {
		log.Printf("Error Permission: %v", err)
		result := helpers.ResponseJSON(false, constants.FORBIDDEN_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusForbidden, result)
	}
	roles, err := h.handler.PermissionService.FindListRole()
	if err != nil {
		log.Printf("Error FindListRole: %v", err)
		result = helpers.ResponseJSON(false, constants.INTERNAL_SERVER_ERROR, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, roles)
	return ctx.JSON(http.StatusOK, result)

}

func (h PermissionHandler) FindListPermission(ctx echo.Context) error {
	var result models.Response
	_, err := helpers.ValidateUserAndRole(ctx, []string{"Admin"})
	if err != nil {
		log.Printf("Error Permission: %v", err)
		result := helpers.ResponseJSON(false, constants.FORBIDDEN_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusForbidden, result)
	}
	permissions, err := h.handler.PermissionService.FindListPermission()
	if err != nil {
		log.Printf("Error FindListPermission: %v", err)
		result = helpers.ResponseJSON(false, constants.INTERNAL_SERVER_ERROR, err.Error(), nil)
		return ctx.JSON(http.StatusInternalServerError, result)
	}
	result = helpers.ResponseJSON(true, constants.SUCCESS_CODE, constants.EMPTY_VALUE, permissions)
	return ctx.JSON(http.StatusOK, result)
}
