package service

import (
	"context"

	"github.com/The-EpaG/Go-BasicServer/internal/modules/user/gateway"
	"github.com/The-EpaG/Go-BasicServer/internal/modules/user/mapper"
	UserDTO "github.com/The-EpaG/Go-BasicServer/internal/modules/user/models/dtos"
	"github.com/The-EpaG/Go-BasicServer/internal/modules/user/repository"
	"github.com/The-EpaG/Go-BasicServer/pkg/logging"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UsersRepository
	UserGateway    gateway.UserGateway
}

func (service *UserServiceImpl) User(ctx context.Context, uuid string) (UserDTO.UserDTO, error) {
	user, err := service.UserRepository.GetUserById(ctx, uuid)
	if err != nil {
		return UserDTO.UserDTO{}, err
	}

	return mapper.MapUserModelToDTO(user), nil
}

func NewUserServiceImpl(userRepository repository.UsersRepository, userGateway gateway.UserGateway) UserService {
	return &UserServiceImpl{UserRepository: userRepository, UserGateway: userGateway}
}

func (service *UserServiceImpl) FindAll() []UserDTO.UserDTO {
	users := service.UserRepository.FindAll()
	var dtos []UserDTO.UserDTO
	for _, user := range users {
		dtos = append(dtos, mapper.MapUserModelToDTO(user))
	}
	return dtos
}

func (service *UserServiceImpl) GetExternalUsers(ctx context.Context) ([]UserDTO.UserDTO, error) {
	users, err := service.UserGateway.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var dtos []UserDTO.UserDTO
	for _, user := range users {
		dtos = append(dtos, mapper.MapUserModelToDTO(user))
	}

	return dtos, nil
}

func (service *UserServiceImpl) CreateUser(ctx context.Context, dto UserDTO.CreateUserDTO) (UserDTO.UserDTO, error) {

	hashPassword, err := hashPassword(dto.Password)
	if err != nil {
		logging.Error(ctx, err.Error())
		return UserDTO.UserDTO{}, err
	}

	userModel := mapper.MapCreateUserDTOToModel(dto, hashPassword)

	// Salva l'utente nel database
	createdUser, err := service.UserRepository.CreateUser(ctx, userModel)
	if err != nil {
		logging.Error(ctx, err.Error())
		return UserDTO.UserDTO{}, err
	}

	// Mappa il model nella risposta DTO
	return mapper.MapUserModelToDTO(createdUser), nil
}

// Funzione per l'hashing della password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
