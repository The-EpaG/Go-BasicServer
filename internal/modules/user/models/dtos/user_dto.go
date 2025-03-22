package dtos

import "time"

type PermissionDTO struct {
	Uuid     string `json:"uuid"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

type RoleDTO struct {
	Uuid        string          `json:"uuid"`
	RoleName    string          `json:"role_name"`
	Description string          `json:"description"`
	Permissions []PermissionDTO `json:"permissions"`
}

type CreateUserDTO struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserDTO struct {
	Uuid      string    `json:"uuid"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	Roles     []RoleDTO `json:"roles"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
