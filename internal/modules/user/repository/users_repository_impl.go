package repository

import (
	"context"
	"errors"

	"github.com/The-EpaG/Go-BasicServer/internal/modules/user/models"
	"github.com/The-EpaG/Go-BasicServer/pkg/exceptions"
	"github.com/The-EpaG/Go-BasicServer/pkg/logging"

	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func (repository *UsersRepositoryImpl) GetUserById(ctx context.Context, uuid string) (models.User, error) {
	var user models.User
	result := repository.Db.WithContext(ctx).Preload("Roles.Permissions").Where("uuid = ?", uuid).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.User{}, exceptions.ErrUserNotFound
	}
	return user, nil
}

func (repository *UsersRepositoryImpl) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	var existingUser models.User
	existingResult := repository.Db.WithContext(ctx).Where("deleted_at IS NULL").Where("email = ?", user.Email).First(&existingUser)
	if existingResult.Error == nil {
		return models.User{}, exceptions.ErrUserAlreadyExists
	} else if !errors.Is(existingResult.Error, gorm.ErrRecordNotFound) {
		// An unexpected error occurred
		logging.Error(ctx, existingResult.Error.Error())
		return models.User{}, exceptions.ErrGenericDb
	}

	result := repository.Db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		logging.Error(ctx, existingResult.Error.Error())
		return models.User{}, exceptions.ErrGenericDb
	}
	return user, nil
}

func (repository *UsersRepositoryImpl) SoftDeleteUser(ctx context.Context, uuid string) error {
	// GORM supports soft delete by default for entities with a gorm.DeletedAt field
	result := repository.Db.WithContext(ctx).Model(&models.User{}).Where("id = ?", uuid).Delete(&models.User{})

	if result.Error != nil {
		logging.Error(ctx, result.Error.Error())
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (repository *UsersRepositoryImpl) FindAll() []models.User {
	var users []models.User
	repository.Db.Where("deleted_at IS NULL").Find(&users)
	return users
}

func NewUsersRepositoryImpl(db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: db}
}
