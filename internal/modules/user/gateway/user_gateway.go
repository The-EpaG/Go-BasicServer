package gateway

import (
	"context"

	"github.com/The-EpaG/Go-BasicServer/internal/modules/user/models"
)

type UserGateway interface {
	FindAll(ctx context.Context) ([]models.User, error)
}
