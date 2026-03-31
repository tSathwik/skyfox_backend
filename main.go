package main

import (
	"log"
	"skyfox_backend/internal/config"
	"skyfox_backend/internal/database"
	"skyfox_backend/internal/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	handler := database.NewDBHandler(cfg)

	database := handler.NewDatabase()
	logger,_ := zap.NewProduction()

	log.Println("database connected successfully")

	err = database.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}
	userRepo := user.NewUserRepository(database,logger)
	userService := user.NewUserService(userRepo,logger)
	userController := user.NewUserController(userService)

	server := gin.Default()
	userGroup :=server.Group("/user")
	userGroup.POST("/",userController.CreateUser)
	userGroup.GET("/:id",userController.GetUserById)
	server.GET("/ping",func (context *gin.Context){
		context.JSON(200,gin.H{"message":"pong"})
	})
	server.Run(":8000")
}
