package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vnniciusg/microservice-upload-image/config"
	"github.com/vnniciusg/microservice-upload-image/internal/adapters/cloudinary"
	"github.com/vnniciusg/microservice-upload-image/internal/adapters/http/rest"
	"github.com/vnniciusg/microservice-upload-image/internal/application"
)

func main() {

	config, err := config.LoadConfig()

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	clourdinaryAdapter := cloudinary.New(config.CloudName, config.CloudAPIKey, config.CloudSecret)

	imageService := application.New(clourdinaryAdapter)

	imageHandler := rest.New(*imageService)

	router := r.Group("/api/v1")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.POST("/upload", imageHandler.UploadImage)

	r.Run(":8082")

}
