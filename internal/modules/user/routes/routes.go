package routes

import (
	"log"

	"github.com/The-EpaG/Go-BasicServer/internal/modules/user/controller"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func Routes(router *gin.Engine, container *dig.Container) {
	userRouter := router.Group("/users")
	err := container.Invoke(func(UserController *controller.UserController) {
		userRouter.POST("/", UserController.RegisterUser)
		userRouter.GET("/:id", UserController.User)
		userRouter.GET("/", UserController.Index)
		userRouter.GET("/external", UserController.ExternalUsers)
	})

	if err != nil {
		log.Fatalf("failed to invoke UserController: %v", err)
	}
}
