package service

import (
	"context"

	UserDTO "github.com/The-EpaG/Go-BasicServer/internal/modules/user/models/dtos"
)

type UserService interface {
	FindAll() []UserDTO.UserDTO
	User(ctx context.Context, uuid string) (UserDTO.UserDTO, error)
	CreateUser(ctx context.Context, dto UserDTO.CreateUserDTO) (UserDTO.UserDTO, error)
	GetExternalUsers(ctx context.Context) ([]UserDTO.UserDTO, error)
}
