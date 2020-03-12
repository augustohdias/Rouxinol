package post

import (
	"context"
	"github.com/augustohdias/rouxinol/models"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (*models.Post, error)
	GetAllByUserID(ctx context.Context, id string) ([]*models.Post, error)
	GetAllByUserIDs(ctx context.Context, ids []string) ([]*models.Post, error)
	New(ctx context.Context, post models.Post) error
	Update(ctx context.Context, post models.Post) error
	Delete(ctx context.Context, post models.Post) error
}
