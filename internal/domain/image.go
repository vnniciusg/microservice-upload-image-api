package domain

import (
	"context"
	"mime/multipart"
)

type Image struct {
	Ctx      context.Context
	File     multipart.File
	PublicID string
}
