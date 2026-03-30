package main

import (
	"log"
	"skyfox_backend/internal/config"
	"skyfox_backend/internal/database"
	"skyfox_backend/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	handler := database.NewDBHandler(cfg)

	database := handler.NewDatabase()

	log.Println("database connected successfully")

	err = database.AutoMigrate(&user.User{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	userRepo := user.NewUserRepository(database)
	userService := user.NewUserService(userRepo)
	userController := user.NewUserController(userService)

	server := gin.Default()
	server.POST("/user",userController.CreateUser)

	server.GET("/ping",func (context *gin.Context){
		context.JSON(200,gin.H{"message":"pong"})
	})
	server.Run(":8000")
}
