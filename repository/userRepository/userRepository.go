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
		INSERT INTO users (username, email, password, role_id, status, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?)
		RETURNING id`

	query = helpers.ReplaceSQL(query, "?")
	err := r.repo.DB.QueryRow(query, req.Username, req.Email, req.Password, req.RoleID, req.Status, req.CreatedAt, req.UpdatedAt).Scan(&ID)
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
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.RoleID, &user.Status, &user.CreatedAt, &user.UpdatedAt)
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
			password,
			role_id
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
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.RoleID)
		if err != nil {
			log.Println("Error scanning row: ", err)
			return user, errors.New("error scanning row")
		}
	}

	return user, nil
}

func (r UserRepository) FindUserPermissions(userID int64) ([]models.UserPermissionModels, error) {
	var permissions []models.UserPermissionModels
	query := `
		SELECT 
			p.id, 
			p.groups,
			p.name,
			rp.status
		FROM 
			permissions p
		JOIN 
			role_permissions rp ON p.id = rp.permission_id
		JOIN 
			users u ON rp.role_id = u.role_id
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
