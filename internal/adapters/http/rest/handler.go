package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vnniciusg/microservice-upload-image/config"
	"github.com/vnniciusg/microservice-upload-image/internal/application"
	"github.com/vnniciusg/microservice-upload-image/internal/domain"
)

type ImageHandler struct {
	ImageService application.ImageService
}

func New(imageService application.ImageService) *ImageHandler {
	return &ImageHandler{
		ImageService: imageService,
	}
}

func (i *ImageHandler) UploadImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")

	if err != nil {
		restErr := config.NewBadRequestErr("Missing or invalid file")
		c.JSON(restErr.Code, restErr)
		return
	}

	ctx := c.Request.Context()

	image := &domain.Image{
		Ctx:      ctx,
		File:     file,
		PublicID: "yqy46hiz",
	}

	uploadedURL, err := i.ImageService.UploadImage.UploadImage(image)

	if err != nil {
		restErr := config.NewInternalServerError(fmt.Sprintf("Failed to upload image : \n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Image uploaded successfully",
		"secure_url":    uploadedURL,
		"original_name": header.Filename,
	})

}
