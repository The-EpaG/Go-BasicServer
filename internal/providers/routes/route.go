package routes

import (
	homeRoutes "github.com/The-EpaG/Go-BasicServer/internal/modules/home/routes"
	userRoutes "github.com/The-EpaG/Go-BasicServer/internal/modules/user/routes"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func RegisterRoutes(router *gin.Engine, container *dig.Container) {
	homeRoutes.Routes(router)
	userRoutes.Routes(router, container)
}
