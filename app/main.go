package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mickaelyoshua7674/finances-app/handlers"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("templates/*")
	server.GET("/", handlers.DefaultHandler)

	server.Run(":8080")
}