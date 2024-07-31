package repository

import "go-auth/models"

type UserRepositoryInterface interface {
	Register(req models.UserModels) (int64, error)
	FindUserByID(id int64) (models.UserModels, error)
	Login(email string) (models.UserModels, error)
}

type PermissionRepositoryInterface interface {
	AssignRoleToUserRequest(req models.AssignRoleToUserRequest) error
	FindUserRole(userID int64) (models.FindUserRoleResponse, error)
	CreateRole(req models.RolesModels) (int64, error)
	CreatePermission(req models.PermissionModels) (int64, error)
	FindListRole() ([]models.RolesModels, error)
	FindListPermission() ([]models.PermissionModels, error)
	CreateRolePermission(req models.RolePermissionModels) (int64, error)
	CreateUserPermission(req models.UserPermissionModels) (int64, error)
	IsUserHavePermission(userID int64, permissionGroup, permissionName string) (bool, error)
	IsRoleHavePermission(userID int64, permissionGroup, permissionName string) (bool, error)
	FindPermissionsForUser(userID int64) ([]models.UserRolePermissionModels, error)
}
