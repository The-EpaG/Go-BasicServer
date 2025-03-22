package models

import (
	"github.com/The-EpaG/Go-BasicServer/internal/modules/home/models"
)

type User struct {
	models.BaseModel
	Username     *string
	Name         string
	Surname      string
	Email        string
	PasswordHash string `json:"password"`

	Roles []Role `gorm:"many2many:user_roles;"`
}

// Error implements error.
func (u User) Error() string {
	panic("unimplemented")
}

type Role struct {
	models.BaseModel
	RoleName    string
	Description string
	Users       []User       `gorm:"many2many:user_roles;"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type Permission struct {
	models.BaseModel
	Resource string
	Action   string
	Roles    []Role `gorm:"many2many:role_permissions;"`
}

type UserRole struct {
	UserUuid string `gorm:"index"`
	RoleUuid string `gorm:"index"`
}
