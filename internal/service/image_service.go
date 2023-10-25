package service

import "github.com/vnniciusg/microservice-upload-image/internal/domain"

type UploaderImage interface {
	UploadImage(image *domain.Image) (string, error)
}
