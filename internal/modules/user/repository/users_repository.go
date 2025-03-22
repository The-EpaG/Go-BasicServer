package repository

import (
	"context"

	"github.com/The-EpaG/Go-BasicServer/internal/modules/user/models"
)

type UsersRepository interface {
	FindAll() []models.User
	GetUserById(ctx context.Context, uuid string) (models.User, error)
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	SoftDeleteUser(ctx context.Context, uuid string) error
}
