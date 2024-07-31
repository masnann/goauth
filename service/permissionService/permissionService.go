package permissionservice

import (
	"go-auth/helpers"
	"go-auth/models"
	"go-auth/service"
	"log"
)

type PermissionService struct {
	service service.Service
}

func NewPermissionService(service service.Service) PermissionService {
	return PermissionService{
		service: service,
	}
}

func (s PermissionService) IsUserHavePermission(userID int64, permissionGroup, permissionName string) (bool, error) {
	result, err := s.service.PermissionRepo.IsUserHavePermission(userID, permissionGroup, permissionName)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (s PermissionService) IsRoleHavePermission(userID int64, permissionGroup, permissionName string) (bool, error) {
	result, err := s.service.PermissionRepo.IsRoleHavePermission(userID, permissionGroup, permissionName)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (s PermissionService) CreateRole(req models.RoleCreateRequest) (int64, error) {
	newData := models.RolesModels{
		Name:      req.Name,
		IsActive:  true,
		CreatedAt: helpers.TimeStampNow(),
		UpdatedAt: "",
	}

	result, err := s.service.PermissionRepo.CreateRole(newData)
	if err != nil {
		log.Println("Error creating role: ", err)
		return 0, err
	}
	return result, nil
}

func (s PermissionService) CreatePermission(req models.PermissionCreateRequest) (int64, error) {
	newData := models.PermissionModels{
		Groups:    req.Groups,
		Name:      req.Name,
		CreatedAt: helpers.TimeStampNow(),
		UpdatedAt: "",
	}
	result, err := s.service.PermissionRepo.CreatePermission(newData)
	if err != nil {
		log.Println("Error creating permission: ", err)
		return 0, err
	}
	return result, nil
}

func (s PermissionService) FindListRole() ([]models.RolesModels, error) {
	result, err := s.service.PermissionRepo.FindListRole()
	if err != nil {
		log.Println("Error finding list role: ", err)
		return nil, err
	}
	return result, nil
}

func (s PermissionService) FindListPermission() ([]models.PermissionModels, error) {
	result, err := s.service.PermissionRepo.FindListPermission()
	if err != nil {
		log.Println("Error finding list permission: ", err)
		return nil, err
	}
	return result, nil
}
