package gateway

import (
	"context"
	"encoding/json"

	UserHttp "github.com/The-EpaG/Go-BasicServer/internal/modules/user/gateway/http"
	"github.com/The-EpaG/Go-BasicServer/internal/modules/user/mapper"
	"github.com/The-EpaG/Go-BasicServer/internal/modules/user/models"
	"github.com/The-EpaG/Go-BasicServer/internal/modules/user/models/dtos"
	"github.com/The-EpaG/Go-BasicServer/pkg/exceptions"
	"github.com/The-EpaG/Go-BasicServer/pkg/logging"
)

type UserGatewayImpl struct {
	userHttpClient UserHttp.UserHttpClient
}

func NewUserGatewayImpl(userHttpClient UserHttp.UserHttpClient) UserGateway {
	return &UserGatewayImpl{
		userHttpClient: userHttpClient,
	}
}

func (gateway *UserGatewayImpl) FindAll(ctx context.Context) ([]models.User, error) {
	body, err := gateway.userHttpClient.Get(ctx, "/users?page=2")

	if err != nil {
		return nil, exceptions.ErrUserExternalService
	}

	var userPage dtos.UserExtPageDTO
	if err := json.Unmarshal(body, &userPage); err != nil {
		logging.Error(ctx, "Error while unmarshalling json: "+err.Error())
		return nil, exceptions.ErrUserExternalService
	}
	var users []models.User
	for _, user := range userPage.Data {
		users = append(users, mapper.MapUserExtDTOToModel(user))
	}

	return users, nil
}
