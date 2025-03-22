package bootstrap

import (
	migration "github.com/The-EpaG/Go-BasicServer/internal/database"
	"github.com/The-EpaG/Go-BasicServer/pkg/bootstrap/di"
	"github.com/The-EpaG/Go-BasicServer/pkg/config"
	"github.com/The-EpaG/Go-BasicServer/pkg/database"
	"github.com/The-EpaG/Go-BasicServer/pkg/routing"
)

func Serve() {

	// Create dependency container
	container := di.BuildContainer()

	config.Init()
	database.Connect()

	cfg := config.Get()
	if cfg.App.Env == "dev" {
		migration.MigrateUp()
	}

	// Start the server
	routing.Init()

	routing.RegisterRoutes(container)

	routing.Serve()

}
