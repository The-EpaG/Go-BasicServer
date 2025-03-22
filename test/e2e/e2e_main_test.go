package e2e_test

import (
	"os"
	"testing"
	"time"

	"github.com/The-EpaG/Go-BasicServer/cmd"
	migration "github.com/The-EpaG/Go-BasicServer/internal/database"
	"github.com/The-EpaG/Go-BasicServer/pkg/config"
	"github.com/The-EpaG/Go-BasicServer/pkg/database"
	"github.com/The-EpaG/Go-BasicServer/pkg/routing"
	"github.com/The-EpaG/Go-BasicServer/test/mocks"
	"go.uber.org/dig"
)

var (
	container *dig.Container
)

func TestMain(m *testing.M) {
	container = mocks.BuildTestContainer()
	cfg, err := config.ReadConfig("../../config", "test", "yaml")
	if err != nil {
		panic(err)
	}
	config.Set(cfg)
	// will use test config (:memory: sqlite)
	database.Connect()

	migration.MigrateUp()
	err = cmd.SeedDb(database.Connection())
	if err != nil {
		panic(err)
	}
	routing.Init()

	routing.RegisterRoutes(container)

	go routing.Serve()
	time.Sleep(time.Second * 1)
	os.Exit(m.Run())
}
