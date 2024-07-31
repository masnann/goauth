package permissionrepository

import (
	"database/sql"
	"errors"
	"go-auth/helpers"
	"go-auth/models"
	"go-auth/repository"
	"log"
)

type PermissionRepository struct {
	repo repository.Repository
}

func NewPermissionRepository(repo repository.Repository) PermissionRepository {
	return PermissionRepository{
		repo: repo,
	}
}

func (r PermissionRepository) AssignRoleToUserRequest(req models.AssignRoleToUserRequest) error {

	query := `
        INSERT INTO user_role (user_id, role_id) 
        VALUES (?,?)
	`

	query = helpers.ReplaceSQL(query, "?")
	_, err := r.repo.DB.Exec(query, req.UserID, req.RoleID)
	if err != nil {
		log.Println("Error querying create user role: ", err)
		return errors.New("error query")
	}

	return nil
}

func (r PermissionRepository) FindUserRole(userID int64) (models.FindUserRoleResponse, error) {
	var userRole models.FindUserRoleResponse
	query := ` 
		SELECT
			u.username,
			u.email,
			r.id as role_id, 
			r.name as role_name
		FROM 
			users u
		JOIN 
			user_role ur ON u.id = ur.user_id
		JOIN 
			roles r ON ur.role_id = r.id 
		WHERE
			u.id = ?
		`
	query = helpers.ReplaceSQL(query, "?")
	rows, err := r.repo.DB.Query(query, userID)
	if err != nil {
		log.Println("Error querying find user role: ", err)
		return userRole, errors.New("error query")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&userRole.Username, &userRole.Email, &userRole.RoleID, &userRole.RoleName)
		if err != nil {
			log.Println("Error scanning row: ", err)
			return userRole, errors.New("error scanning row")
		}
	}
	return userRole, nil
}

func (r PermissionRepository) CreateRole(req models.RolesModels) (int64, error) {
	var ID int64

	query := `
		INSERT INTO roles
            (name, is_active, created_at, updated_at)
        VALUES (?,?,?,?)
        RETURNING id`

	query = helpers.ReplaceSQL(query, "?")
	err := r.repo.DB.QueryRow(query, req.Name, req.IsActive, req.CreatedAt, req.UpdatedAt).Scan(&ID)
	if err != nil {
		log.Println("Error querying create role: ", err)
		return ID, errors.New("error query")
	}

	return ID, nil
}

func (r PermissionRepository) CreatePermission(req models.PermissionModels) (int64, error) {
	var ID int64

	query := `
        INSERT INTO permissions
            (groups, name, created_at, updated_at)
        VALUES (?,?, ?, ?)
        RETURNING id`

	query = helpers.ReplaceSQL(query, "?")
	err := r.repo.DB.QueryRow(query, req.Groups, req.Name, req.CreatedAt, req.UpdatedAt).Scan(&ID)
	if err != nil {
		log.Println("Error querying create permission: ", err)
		return ID, errors.New("error query")
	}

	return ID, nil
}

func (r PermissionRepository) FindListRole() ([]models.RolesModels, error) {
	var roles []models.RolesModels

	query := `
		SELECT 
			id, name, is_active, created_at, updated_at
		FROM 
			roles
	`
	query = helpers.ReplaceSQL(query, "?")
	rows, err := r.repo.DB.Query(query)
	if err != nil {
		log.Println("Error querying find list role: ", err)
		return roles, errors.New("error query")
	}
	defer rows.Close()
	for rows.Next() {
		var row models.RolesModels
		err := rows.Scan(&row.ID, &row.Name, &row.IsActive, &row.CreatedAt, &row.UpdatedAt)
		if err != nil {
			log.Println("Error scanning row: ", err)
			return roles, errors.New("error scanning row")
		}
		roles = append(roles, row)
	}
	return roles, nil

}

func (r PermissionRepository) FindListPermission() ([]models.PermissionModels, error) {
	var permissions []models.PermissionModels

	query := `
		SELECT 
			id, groups, name, created_at, updated_at
		FROM 
			permissions
	`
	query = helpers.ReplaceSQL(query, "?")
	rows, err := r.repo.DB.Query(query)
	if err != nil {
		log.Println("Error querying find list permission: ", err)
		return permissions, errors.New("error query")
	}
	defer rows.Close()
	for rows.Next() {
		var row models.PermissionModels
		err := rows.Scan(&row.ID, &row.Groups, &row.Name, &row.CreatedAt, &row.UpdatedAt)
		if err != nil {
			log.Println("Error scanning row: ", err)
			return permissions, errors.New("error scanning row")
		}
		permissions = append(permissions, row)
	}
	return permissions, nil
}

func (r PermissionRepository) CreateRolePermission(req models.RolePermissionModels) (int64, error) {
	var ID int64

	query := `
		INSERT INTO role_permissions
			(role_id, permission_id, created_at, updated_at)
		VALUES (?, ?, ?, ?)
		RETURNING id`

	query = helpers.ReplaceSQL(query, "?")
	err := r.repo.DB.QueryRow(query, req.RoleID, req.PermissionID, req.CreatedAt, req.UpdatedAt).Scan(&ID)
	if err != nil {
		log.Println("Error querying create permission: ", err)
		return ID, errors.New("error query")
	}

	return ID, nil
}

func (r PermissionRepository) CreateUserPermission(req models.UserPermissionModels) (int64, error) {
	var ID int64
	query := `
		INSERT INTO user_permissions
		    (user_id, permission_id, status, granted_by, granted_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
		RETURNING id`

	query = helpers.ReplaceSQL(query, "?")
	err := r.repo.DB.QueryRow(query, req.UserID, req.PermissionID, req.Status, req.GrantedBy, req.GrantedAt, req.UpdatedAt).Scan(&ID)
	if err != nil {
		log.Println("Error querying create permission: ", err)
		return ID, errors.New("error query")
	}
	return ID, nil
}

func (r PermissionRepository) IsUserHavePermission(userID int64, permissionGroup, permissionName string) (bool, error) {
	var permission models.UserRolePermissionModels
	query := `
        SELECT 
			p.id, 
			p.groups AS permission_group, 
			p.name AS permission_name, 
			up.status AS permission_status
        FROM 
			user_permissions up
        JOIN 
			permissions p ON up.permission_id = p.id
        WHERE 
			up.user_id = ? AND p.groups = ? AND p.name = ? AND up.status = 'true'
    `
	query = helpers.ReplaceSQL(query, "?")
	row := r.repo.DB.QueryRow(query, userID, permissionGroup, permissionName)
	err := row.Scan(&permission.ID, &permission.Group, &permission.Name, &permission.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r PermissionRepository) IsRoleHavePermission(userID int64, permissionGroup, permissionName string) (bool, error) {
	query := `
        SELECT 
            COUNT(*) > 0
        FROM 
            user_role ur
        JOIN 
            role_permissions rp ON ur.role_id = rp.role_id
        JOIN 
            permissions p ON rp.permission_id = p.id
        WHERE 
            ur.user_id = ? AND p.groups = ? AND p.name = ?
    `

	query = helpers.ReplaceSQL(query, "?")

	var hasPermission bool
	err := r.repo.DB.QueryRow(query, userID, permissionGroup, permissionName).Scan(&hasPermission)
	if err != nil {
		log.Println("Error querying role permissions: ", err)
		return false, err
	}

	return hasPermission, nil
}

func (r PermissionRepository) FindPermissionsForUser(userID int64) ([]models.UserRolePermissionModels, error) {
	permissionsMap := make(map[string]models.UserRolePermissionModels)

	rolePermissionsQuery := `
    SELECT
		p.id,
		p.groups, 
		p.name, 
		rp.status 
    FROM 
		role_permissions rp
    JOIN 
		permissions p ON rp.permission_id = p.id
    JOIN 
		user_role ur ON ur.role_id = rp.role_id
    WHERE 
		ur.user_id = $1`

	roleRows, err := r.repo.DB.Query(rolePermissionsQuery, userID)
	if err != nil {
		return nil, err
	}
	defer roleRows.Close()

	for roleRows.Next() {
		var row models.UserRolePermissionModels
		if err := roleRows.Scan(&row.ID, &row.Group, &row.Name, &row.Status); err != nil {
			return nil, err
		}
		key := row.Group + "_" + row.Name
		permissionsMap[key] = row
	}

	// Get user-specific permissions
	userPermissionsQuery := `
    SELECT 
		p.id, 
		p.groups, 
		p.name, 
		up.status 
    FROM 
		user_permissions up
    JOIN 
		permissions p ON up.permission_id = p.id
    WHERE 
		up.user_id = $1`

	userRows, err := r.repo.DB.Query(userPermissionsQuery, userID)
	if err != nil {
		return nil, err
	}
	defer userRows.Close()

	for userRows.Next() {
		var row models.UserRolePermissionModels
		if err := userRows.Scan(&row.ID, &row.Group, &row.Name, &row.Status); err != nil {
			return nil, err
		}
		key := row.Group + "_" + row.Name
		permissionsMap[key] = row
	}

	var permissions []models.UserRolePermissionModels
	for _, permission := range permissionsMap {
		permissions = append(permissions, permission)
	}

	return permissions, nil
}
