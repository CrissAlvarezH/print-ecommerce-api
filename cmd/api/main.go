package main

import (
	"github.com/gin-gonic/gin"
	"log"
)
import "net/http"

func main() {
	app := gin.New()

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	err := app.Run(":8000")
	if err != nil {
		log.Fatal("Server error", err)
	}
}
