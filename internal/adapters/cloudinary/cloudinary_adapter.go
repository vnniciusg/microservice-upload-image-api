package cloudinary

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/vnniciusg/microservice-upload-image/internal/domain"
)

type ImageUploaderAdapter struct {
	CloudName   string
	CloudAPIKey string
	CloudSecret string
}

func New(cloudName, cloudApiKey, cloudSecret string) *ImageUploaderAdapter {
	return &ImageUploaderAdapter{
		CloudName:   cloudName,
		CloudAPIKey: cloudApiKey,
		CloudSecret: cloudSecret,
	}
}

func (c *ImageUploaderAdapter) UploadImage(image *domain.Image) (string, error) {
	config, _ := cloudinary.NewFromParams(c.CloudName, c.CloudAPIKey, c.CloudSecret)

	uploadParams := uploader.UploadParams{
		PublicID: image.PublicID,
	}

	result, err := config.Upload.Upload(image.Ctx, image.File, uploadParams)

	if err != nil {
		return "", err
	}

	return result.URL, nil
}
