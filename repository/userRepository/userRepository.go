package userrepository

import (
	"errors"
	"go-auth/helpers"
	"go-auth/models"
	"go-auth/repository"
	"log"
)

type UserRepository struct {
	repo repository.Repository
}

func NewUserRepository(repo repository.Repository) UserRepository {
	return UserRepository{
		repo: repo,
	}
}

func (r UserRepository) Register(req models.UserModels) (int64, error) {
	var ID int64
	query := `
		INSERT INTO users (username, email, password, status, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?)
		RETURNING id`

	query = helpers.ReplaceSQL(query, "?")
	err := r.repo.DB.QueryRow(query, req.Username, req.Email, req.Password, req.Status, req.CreatedAt, req.UpdatedAt).Scan(&ID)
	if err != nil {
		log.Println("Error querying register: ", err)
		return ID, errors.New("error query")
	}

	return ID, nil
}

func (r UserRepository) FindUserByID(id int64) (models.UserModels, error) {
	var user models.UserModels
	query := `SELECT * FROM users WHERE id = ? AND status = ''`

	query = helpers.ReplaceSQL(query, "?")

	rows, err := r.repo.DB.Query(query, id)
	if err != nil {
		log.Println("Error querying find product: ", err)
		return user, errors.New("error query")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Status, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Println("Error scanning row: ", err)
			return user, errors.New("error scanning row")
		}
	}
	return user, nil
}

func (r UserRepository) Login(email string) (models.UserModels, error) {
	var user models.UserModels
	query := `
		SELECT 
			id, 
			username, 
			email, 
			password
		FROM users 
		WHERE email =?
	`

	query = helpers.ReplaceSQL(query, "?")

	rows, err := r.repo.DB.Query(query, email)
	if err != nil {
		log.Println("Error querying login: ", err)
		return user, errors.New("error query")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			log.Println("Error scanning row: ", err)
			return user, errors.New("error scanning row")
		}
	}

	return user, nil
}

func (r UserRepository) FindListUserPermissions(userID int64) ([]models.UserPermissionModels, error) {
	var permissions []models.UserPermissionModels
	query := `
		SELECT 
			p.id AS permissions_id,
			p.groups AS permission_group,
			p.name AS permission_name,
			rp.status
		FROM 
			users u
		JOIN 
			user_role ur ON u.id = ur.user_id
		JOIN 
			roles r ON ur.role_id = r.id
		JOIN 
			role_permissions rp ON r.id = rp.role_id
		JOIN 
			permissions p ON rp.permission_id = p.id
		WHERE 
			u.id = ?

    `
	query = helpers.ReplaceSQL(query, "?")

	rows, err := r.repo.DB.Query(query, userID)
	if err != nil {
		log.Println("Error querying find user permission: ", err)
		return permissions, errors.New("error query")
	}
	defer rows.Close()
	for rows.Next() {
		var permission models.UserPermissionModels
		err := rows.Scan(&permission.ID, &permission.Group, &permission.Name, &permission.Status)
		if err != nil {
			log.Println("Error scanning row: ", err)
			return permissions, errors.New("error scanning row")
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func (r UserRepository) AssignRoleToUserRequest(req models.AssignRoleToUserRequest) error {

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

func (r UserRepository) FindUserRole(userID int64) (models.FindUserRoleResponse, error) {
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

func (r UserRepository) FindUserPermissions(userID int64, permissionGroup, permissionName string) (models.UserPermissionModels, error) {
	var permission models.UserPermissionModels
	query := `
        SELECT 
            p.id AS permissions_id,
            p.groups AS permission_group,
            p.name AS permission_name,
            rp.status
        FROM 
            users u
        JOIN 
            user_role ur ON u.id = ur.user_id
        JOIN 
            roles r ON ur.role_id = r.id
        JOIN 
            role_permissions rp ON r.id = rp.role_id
        JOIN 
            permissions p ON rp.permission_id = p.id
        WHERE 
            u.id =? AND p.groups =? AND p.name =?
    `

	query = helpers.ReplaceSQL(query, "?")

	rows, err := r.repo.DB.Query(query, userID, permissionGroup, permissionName)
	if err != nil {
		log.Println("Error querying find user permission: ", err)
		return permission, errors.New("error query")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&permission.ID, &permission.Group, &permission.Name, &permission.Status)
		if err != nil {
			log.Println("Error scanning row: ", err)
			return permission, errors.New("error scanning row")
		}
	}
	return permission, nil
}
