package repository

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	DB      *sql.DB
	MongoDB *mongo.Database
}

func NewRepository(db *sql.DB, MongoDB *mongo.Database) Repository {
	return Repository{
		DB:      db,
		MongoDB: MongoDB,
	}
}
