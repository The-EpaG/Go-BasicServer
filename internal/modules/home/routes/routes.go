package routes

import (
	"github.com/The-EpaG/Go-BasicServer/internal/modules/home/controller"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	homeController := controller.NewHomeController()

	router.GET("/", homeController.Index)
	router.GET("/health", homeController.Health)
	router.GET("/version", homeController.Version)
}
