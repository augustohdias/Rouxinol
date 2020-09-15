package user

import (
	"context"
	"github.com/augustohdias/Rouxinol/models"
)

type Repository interface {
	GetById(ctx context.Context, id string) (*models.User, error)
}
