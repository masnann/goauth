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
		INSERT INTO users (username, email, password, role, status, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?)
		RETURNING id`

	query = helpers.ReplaceSQL(query, "?")
	err := r.repo.DB.QueryRow(query, req.Username, req.Email, req.Password, req.Role, req.Status, req.CreatedAt, req.UpdatedAt).Scan(&ID)
	if err != nil {
		log.Println("Error querying: ", err)
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
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.Status, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Println("Error scanning row: ", err)
			return user, errors.New("error scanning row")
		}
	}
	return user, nil
}
