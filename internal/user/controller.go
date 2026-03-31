package user

import (
	"net/http"
	"skyfox_backend/internal/user/dto"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
}

type userController struct {
	service UserService
}

func NewUserController(service UserService) UserController {
	return &userController{service: service}
}

func (c *userController) CreateUser(ctx *gin.Context){
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error(),
		})
		return 
	}
	user, err := c.service.CreateUser(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error": err.Error(),
		})
		return 
	}
	ctx.JSON(http.StatusCreated,user)
}


func (c *userController) GetUserById(ctx *gin.Context) {
	userId := ctx.Params.ByName("id")
	user, err := c.service.GetUserById(userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}