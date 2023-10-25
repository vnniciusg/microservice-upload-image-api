package application

import (
	"github.com/vnniciusg/microservice-upload-image/internal/service"
)

type ImageService struct {
	UploadImage service.UploaderImage
}

func New(uploadImage service.UploaderImage) *ImageService {
	return &ImageService{
		UploadImage: uploadImage,
	}
}
