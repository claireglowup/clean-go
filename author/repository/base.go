package repository

import (
	"context"
	"database/sql"
	"go-cleanv2-riky/author"
	"go-cleanv2-riky/models"

	"github.com/sirupsen/logrus"
)

type authorRepo struct {
	DB *sql.DB
}

func NewAuthorRepo(db *sql.DB) author.Repository {
	return &authorRepo{
		DB: db,
	}
}

func (ar *authorRepo) getOne(ctx context.Context, query string, args ...interface{}) (*models.Author, error) {

	stmt, err := ar.DB.PrepareContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	row := stmt.QueryRowContext(ctx, args...)
	a := &models.Author{}

	err = row.Scan(
		&a.ID,
		&a.Name,
		&a.CreatedAt,
		&a.UpdatedAt,
	)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return a, nil

}

func (ar *authorRepo) GetById(ctx context.Context, id int) (*models.Author, error) {
	query := `SELECT id, name, created_at, updated_at from author WHERE id = ?`
	return ar.getOne(ctx, query, id)
}
