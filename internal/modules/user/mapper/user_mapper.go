package mapper

import (
	"github.com/The-EpaG/Go-BasicServer/internal/modules/user/models"
	"github.com/The-EpaG/Go-BasicServer/internal/modules/user/models/dtos"
)

func MapPermissionsToDTO(permissions []models.Permission) []dtos.PermissionDTO {
	permissionDTOs := make([]dtos.PermissionDTO, len(permissions))
	for i, permission := range permissions {
		permissionDTOs[i] = dtos.PermissionDTO{
			Uuid:     permission.Uuid.String(),
			Resource: permission.Resource,
			Action:   permission.Action,
		}
	}
	return permissionDTOs
}

func MapRolesToDTO(roles []models.Role) []dtos.RoleDTO {
	roleDTOs := make([]dtos.RoleDTO, len(roles))
	for i, role := range roles {
		roleDTOs[i] = dtos.RoleDTO{
			Uuid:        role.Uuid.String(),
			RoleName:    role.RoleName,
			Description: role.Description,
			Permissions: MapPermissionsToDTO(role.Permissions), // Mappa i permessi
		}
	}
	return roleDTOs
}

func MapCreateUserDTOToModel(dto dtos.CreateUserDTO, hashedPassword string) models.User {

	return models.User{
		Username:     &dto.Username,
		Name:         dto.Name,
		Surname:      dto.Surname,
		Email:        dto.Email,
		PasswordHash: hashedPassword, // Salviamo l'hash della password
	}
}

func MapUserModelToDTO(user models.User) dtos.UserDTO {
	return dtos.UserDTO{
		Uuid:      user.Uuid.String(),
		Username:  getValue(user.Username),
		Name:      user.Name,
		Surname:   user.Surname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Roles:     MapRolesToDTO(user.Roles),
	}
}

func MapUserExtDTOToModel(dto dtos.UserExtDTO) models.User {
	return models.User{
		Username:     &dto.Email,
		Name:         dto.FirstName,
		Surname:      dto.LastName,
		Email:        dto.Email,
		PasswordHash: "",
	}
}

func getValue(str *string) string {
	if str != nil {
		return *str
	}
	return ""
}
