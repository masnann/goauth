package models

type PermissionModels struct {
	ID    int64  `json:"id:"`
	Group string `json:"group"`
	Name  string `json:"name"`
}

type RoleModels struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type RolePermissionModels struct {
	ID           int64 `json:"id:"`
	RoleID       int64 `json:"roleID"`
	PermissionID int64 `json:"permissionID"`
	Status       bool  `json:"status"`
}

type UserPermissionModels struct {
	ID     int64  `json:"id:"`
	Group  string `json:"group"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type AssignRoleToUserRequest struct {
	UserID int64 `json:"userID"`
	RoleID int64 `json:"roleID"`
}
