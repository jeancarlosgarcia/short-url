package service

import (
	"context"

	"short-url/interfaces"
	"short-url/models"
	"short-url/models/api"
	"short-url/util"
)

type urlService struct {
	repo interfaces.IUrlRepository
}

func NewURLService(repo interfaces.IUrlRepository) interfaces.IUrlService {
	return &urlService{ repo }
}

func (u urlService) Create(ctx context.Context, request *api.URLCreate) api.URLCreated {
	short := util.GenerateShortLink(request.Url)

	shortURL := models.ShortURL{
		Short: short,
		Original: request.Url,
	}

	var error string

	err := u.repo.Save(ctx, &shortURL)
	if err != nil {
		short = ""
		error = err.Error()
	}

	return api.URLCreated{
		ShortURL: short,
		Error:    error,
	}
}

func (u urlService) GetOriginal(ctx context.Context, shortURL string) (string, error) {
	return u.repo.Get(ctx, shortURL)
}

func (u urlService) Delete(ctx context.Context, shortURL string) error {
	return u.repo.Delete(ctx, shortURL)
}