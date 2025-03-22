package migrate

import (
	migration "github.com/The-EpaG/Go-BasicServer/internal/database"
	"github.com/The-EpaG/Go-BasicServer/pkg/config"
	"github.com/The-EpaG/Go-BasicServer/pkg/database"
)

func Up() {
	config.Init()

	database.Connect()

	migration.MigrateUp()
}
