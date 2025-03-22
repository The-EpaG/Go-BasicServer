package migration

import (
	"log"

	"github.com/The-EpaG/Go-BasicServer/pkg/database"
)

func MigrateUp() {
	db := database.Connection()

	// Migrate the schema
	if err := db.AutoMigrate(RegisterModels()...); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("Database migrated")

}
