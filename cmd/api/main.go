package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)
import "net/http"

func main() {
	app := gin.New()

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	app.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["images"]
		dst := "./images/"

		for _, file := range files {
			log.Println(file.Filename)

			// Upload the file to specific dst.
			err := c.SaveUploadedFile(file, dst)
			if err != nil {
				log.Println("error on save image ", file.Filename)
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	err := app.Run(":8000")
	if err != nil {
		log.Fatal("Server error", err)
	}
}
