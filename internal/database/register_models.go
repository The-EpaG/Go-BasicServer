package migration

import (
	"github.com/The-EpaG/Go-BasicServer/internal/modules/user/models"
)

// RegisterModels contains the models to be migrated
func RegisterModels() []interface{} {
	return []interface{}{
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.UserRole{},
	}
}
