package userrepository

import (
	"database/sql"
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
	query := `
		SELECT 
			id, username, email, password, status, created_at, updated_at
		FROM 
			users WHERE id = ? AND status = 'active'`

	query = helpers.ReplaceSQL(query, "?")

	row := r.repo.DB.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Status, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found")
		}
		log.Println("Error scanning row: ", err)
		return user, errors.New("error scanning row")
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
