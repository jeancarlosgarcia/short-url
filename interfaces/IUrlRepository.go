package interfaces

import (
	"context"

	"short-url/models"
)

type IUrlRepository interface {
	Get(ctx context.Context, shortURL string) (string, error)
	Save(ctx context.Context, shortURL *models.ShortURL) error
	Delete(ctx context.Context, shortURL string) error
}