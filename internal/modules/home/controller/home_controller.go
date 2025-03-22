package controller

import (
	"github.com/gin-gonic/gin"

	env "github.com/The-EpaG/Go-BasicServer/config"
	config "github.com/The-EpaG/Go-BasicServer/pkg/config"
)

type HomeController struct {
	cfg *env.Config
}

func NewHomeController() *HomeController {
	cnf := config.Get()
	return &HomeController{cfg: &cnf}
}

func (controller *HomeController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"AppName": controller.cfg.App.Name,
		"Env":     controller.cfg.App.Env,
		"Version": controller.cfg.App.Version,
	})
}

func (controller *HomeController) Health(c *gin.Context) {
	c.Done()
}

func (controller *HomeController) Version(c *gin.Context) {
	c.String(200, controller.cfg.App.Version)
}
