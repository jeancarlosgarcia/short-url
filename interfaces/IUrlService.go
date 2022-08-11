package interfaces

import (
	"context"

	"short-url/models/api"
)

type IUrlService interface {
	GetOriginal(ctx context.Context, shortURL string) (string, error)
	Create(ctx context.Context, request *api.URLCreate) api.URLCreated
	Delete(ctx context.Context, shortURL string) error
}