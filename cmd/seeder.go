package cmd

import (
	"fmt"
	"log"

	"github.com/The-EpaG/Go-BasicServer/internal/modules/user/models"
	"github.com/The-EpaG/Go-BasicServer/pkg/config"
	"github.com/The-EpaG/Go-BasicServer/pkg/database"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var seederCmd = &cobra.Command{
	Use:   "seeder",
	Short: "Seed database",
	Long:  "Seed database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start seeding roles and permissions...")
		// Retrieve the database from the repository
		config.Init()
		db, err := database.CreateConnection()
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
		// Execute seeding
		err = SeedDb(db)
		if err != nil {
			log.Fatalf("Error during seeding: %v", err)
		}

		fmt.Println("Seeding completed successfully!")
	},
}

func SeedDb(db *gorm.DB) error {
	err := SeedRolesAndPermissions(db)
	if err != nil {
		return err
	}
	return nil
}

// Seed roles and permissions
func SeedRolesAndPermissions(db *gorm.DB) error {
	// Define roles
	roles := []models.Role{
		{RoleName: "Admin", Description: "Administrator with full access"},
		{RoleName: "Vendor", Description: "Vendor with limited access"},
		{RoleName: "User", Description: "Basic user"},
	}

	// Define permissions for user management
	permissions := []models.Permission{
		{Resource: "users", Action: "create"},
		{Resource: "users", Action: "read"},
		{Resource: "users", Action: "update"},
		{Resource: "users", Action: "delete"},
	}

	users := []models.User{
		{
			Username:     nil,
			Name:         "admin",
			Surname:      "admin",
			Email:        "admin@example.com",
			PasswordHash: "admin",
			Roles:        []models.Role{roles[0]}},
	}

	// Save roles to the database
	for _, role := range roles {
		var existingRole models.Role
		if err := db.Where("role_name = ?", role.RoleName).First(&existingRole).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&role).Error; err != nil {
					return fmt.Errorf("error creating role %s: %w", role.RoleName, err)
				}
				fmt.Printf("Role created: %s\n", role.RoleName)
			} else {
				return err
			}
		}
	}

	// Save permissions to the database
	for _, perm := range permissions {
		var existingPerm models.Permission
		if err := db.Where("resource = ? AND action = ?", perm.Resource, perm.Action).First(&existingPerm).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&perm).Error; err != nil {
					return fmt.Errorf("error creating permission %s-%s: %w", perm.Resource, perm.Action, err)
				}
				fmt.Printf("Permission created: %s-%s\n", perm.Resource, perm.Action)
			} else {
				return err
			}
		}
	}

	// Associate permissions only with Admin
	var adminRole models.Role
	if err := db.Where("role_name = ?", "Admin").First(&adminRole).Error; err != nil {
		return fmt.Errorf("error retrieving Admin role: %w", err)
	}

	for _, perm := range permissions {
		var permission models.Permission
		if err := db.Where("resource = ? AND action = ?", perm.Resource, perm.Action).First(&permission).Error; err == nil {
			if err := db.Model(&adminRole).Association("Permissions").Append(&permission); err != nil {
				return fmt.Errorf("error associating permission %s-%s with Admin role: %w", perm.Resource, perm.Action, err)
			}
		}
	}

	fmt.Println("Roles and permissions associated successfully!")

	for _, user := range users {
		var existingUser models.User
		if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&user).Error; err != nil {
					return fmt.Errorf("error creating user %s: %w", user.Email, err)
				}
				fmt.Printf("User created: %s\n", user.Email)
			} else {
				return err
			}
		}
	}

	fmt.Println("Users created successfully!")

	return nil
}

// Register the Cobra command
func init() {
	rootCmd.AddCommand(seederCmd)
}
