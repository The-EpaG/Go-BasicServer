package routing

import (
	"github.com/The-EpaG/Go-BasicServer/internal/middleware"
	"github.com/The-EpaG/Go-BasicServer/internal/providers/routes"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func Init() {
	router = gin.Default()
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.HTTPLogger())
	router.Use(middleware.ErrorHandler())
}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRoutes(container *dig.Container) {
	routes.RegisterRoutes(GetRouter(), container)
}
