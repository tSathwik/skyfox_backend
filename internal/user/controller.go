package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
}

type userController struct {
	service UserService
}

func NewUserController(service UserService) UserController {
	return &userController{service: service}
}

func (c *userController) CreateUser(ctx *gin.Context){
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error(),
		})
		return 
	}
	err := c.service.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"error": err.Error(),
		})
		return 
	}
	ctx.JSON(http.StatusCreated,user)
}