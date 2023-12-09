package repository

import (
	"context"
	"database/sql"
	"log"
)

type Repository interface {
}

type RepositoryImpl struct {
	sqlDatabase *sql.DB
}

func New(ctx context.Context) (Repository, error) {

	// ================================================================= FIREBASE & FIRESTORE

	connectionString := "postgres://username:password@your-database-url:5432/your-database-name"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return &RepositoryImpl{
		sqlDatabase: db,
	}, nil
}
