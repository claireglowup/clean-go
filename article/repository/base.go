package repository

import (
	"context"
	"database/sql"
	"go-cleanv2-riky/article"
	"go-cleanv2-riky/models"

	"github.com/sirupsen/logrus"
)

type articleRepo struct {
	DB *sql.DB
}

func NewArticleRepo(db *sql.DB) article.Repository {
	return &articleRepo{
		DB: db,
	}
}

func (ar *articleRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Article, error) {
	rows, err := ar.DB.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			logrus.Error(err)
		}
	}()

	result := make([]*models.Article, 0)
	for rows.Next() {
		t := new(models.Article)
		authorId := 0
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&authorId,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		t.Author = models.Author{
			ID: authorId,
		}

		result = append(result, t)
	}

	return result, nil

}
