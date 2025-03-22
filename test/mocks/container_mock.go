package mocks

import (
	"fmt"

	"github.com/The-EpaG/Go-BasicServer/pkg/bootstrap/di"
	"github.com/The-EpaG/Go-BasicServer/pkg/config"
	"github.com/The-EpaG/Go-BasicServer/pkg/database"
	"github.com/The-EpaG/Go-BasicServer/pkg/http"
	"go.uber.org/dig"
)

func BuildTestContainer() *dig.Container {
	container := dig.New()

	if err := container.Provide(config.Get); err != nil {
		fmt.Println("Failed to provide config:", err)
	}

	if err := container.Provide(database.Connection); err != nil {
		fmt.Println("Failed to provide database connection:", err)
	}

	if err := container.Provide(func() http.HttpClient {
		return &MockHttpClient{}
	}); err != nil {
		fmt.Println("Failed to provide http client:", err)
	}

	di.BuildApplicationContainer(container)

	return container
}
