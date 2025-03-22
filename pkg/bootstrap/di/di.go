package di

import (
	"fmt"

	userController "github.com/The-EpaG/Go-BasicServer/internal/modules/user/controller"
	userGateway "github.com/The-EpaG/Go-BasicServer/internal/modules/user/gateway"
	userHttp "github.com/The-EpaG/Go-BasicServer/internal/modules/user/gateway/http"
	userRepository "github.com/The-EpaG/Go-BasicServer/internal/modules/user/repository"
	userService "github.com/The-EpaG/Go-BasicServer/internal/modules/user/service"
	"github.com/The-EpaG/Go-BasicServer/pkg/config"
	"github.com/The-EpaG/Go-BasicServer/pkg/database"

	// "github.com/The-EpaG/Go-BasicServer/pkg/http"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	container = BuildExternalContainer(container)
	container = BuildApplicationContainer(container)

	return container
}

// Used for external dependencies
func BuildExternalContainer(container *dig.Container) *dig.Container {
	if err := container.Provide(config.Get); err != nil {
		fmt.Println("Failed to provide config:", err)
	}

	if err := container.Provide(database.Connection); err != nil {
		fmt.Println("Failed to provide database connection:", err)
	}

	// if err := container.Provide(http.NewHttpClientImpl); err != nil {
	// 	fmt.Println("Failed to provide http client:", err)
	// }

	return container
}

// Used for application dependencies
func BuildApplicationContainer(container *dig.Container) *dig.Container {
	// User module
	if err := container.Provide(userHttp.NewUserHttpClientImpl); err != nil {
		fmt.Println("Failed to provide user gateway:", err)
	}

	if err := container.Provide(userGateway.NewUserGatewayImpl); err != nil {
		fmt.Println("Failed to provide user gateway:", err)
	}

	if err := container.Provide(userRepository.NewUsersRepositoryImpl); err != nil {
		fmt.Println("Failed to provide user repository:", err)
	}

	if err := container.Provide(userService.NewUserServiceImpl); err != nil {
		fmt.Println("Failed to provide user service:", err)
	}

	if err := container.Provide(userController.NewUserController); err != nil {
		fmt.Println("Failed to provide user controller:", err)
	}

	// Other modules

	return container
}
