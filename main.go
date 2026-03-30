package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/ping",func (context *gin.Context){
		context.String(200,"pong")
	})
	server.Run(":8000")
}