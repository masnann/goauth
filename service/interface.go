package service

import "go-auth/models"

type UserServiceInterface interface {
	Register(req models.UserRegisterRequest) (int64, error)
	FindUserByID(req models.RequestID) (models.UserModels, error)
	Login(req models.UserLoginRequest) (models.UserLoginResponse, error)
	RefreshToken(accessToken string) (models.UserLoginResponse, error)
}

type PermissionServiceInterface interface {
	IsUserHavePermission(userID int64, permissionGroup, permissionName string) (bool, error)
	IsRoleHavePermission(userID int64, permissionGroup, permissionName string) (bool, error)
	CreateRole(req models.RoleCreateRequest) (int64, error) 
	CreatePermission(req models.PermissionCreateRequest) (int64, error)
	FindListRole() ([]models.RolesModels, error)
	FindListPermission() ([]models.PermissionModels, error)
}