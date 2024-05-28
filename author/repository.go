package author

import (
	"context"
	"go-cleanv2-riky/models"
)

type Repository interface {
	GetById(ctx context.Context, id int) (*models.Author, error)
}
