package controller

import (
	"net/http"

	"github.com/The-EpaG/Go-BasicServer/internal/modules/user/models/dtos"
	"github.com/The-EpaG/Go-BasicServer/internal/modules/user/service"
	"github.com/The-EpaG/Go-BasicServer/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{userService: service}
}

func (controller *UserController) Index(ctx *gin.Context) {
	users := controller.userService.FindAll()
	ctx.JSON(http.StatusOK, users)
}

func (controller *UserController) ExternalUsers(ctx *gin.Context) {
	users, err := controller.userService.GetExternalUsers(ctx.Request.Context())
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (controller *UserController) User(ctx *gin.Context) {
	idParam := ctx.Param("id")

	_, err := uuid.Parse(idParam)
	if err != nil {
		_ = ctx.Error(exceptions.ErrUserBadId)
		return
	}

	user, err := controller.userService.User(ctx.Request.Context(), idParam)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (controller *UserController) RegisterUser(ctx *gin.Context) {
	var dto dtos.CreateUserDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		_ = ctx.Error(exceptions.ErrInvalidRequestPayload)
		return
	}

	createdUser, err := controller.userService.CreateUser(ctx.Request.Context(), dto)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}
