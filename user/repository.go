package user

import (
	"context"
	"github.com/augustohdias/rouxinol/models"
)

type Repository interface {
	GetById(ctx context.Context, id string) (*models.User, error)
}
