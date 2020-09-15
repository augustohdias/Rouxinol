package post

import (
	"context"

	"github.com/augustohdias/Rouxinol/models"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (*models.Post, error)
	GetAllByUsername(ctx context.Context, uername string) ([]*models.Post, error)
	GetAllByUsernames(ctx context.Context, usernames []string) ([]*models.Post, error)
	New(ctx context.Context, post models.Post) error
	Update(ctx context.Context, post models.Post) error
	Delete(ctx context.Context, post models.Post) error
}
