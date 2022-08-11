package repository

import (
	"context"
	"fmt"
	"time"

	"short-url/interfaces"
	"short-url/models"

	"github.com/labstack/gommon/log"
)

type URLRepository struct {
	storage interfaces.IStorage
	cache   interfaces.ICache
}

func NewUrlRepository(storage interfaces.IStorage, cache interfaces.ICache) interfaces.IUrlRepository {
	return &URLRepository{storage, cache }
}

func (u URLRepository) Save(ctx context.Context, shortURL *models.ShortURL) error {
	err := u.storage.Insert(ctx, shortURL)
	if err != nil {
		log.Error("error insert storage: "+ err.Error()+ " original: "+shortURL.Original)

		return err
	}

	err = u.cache.Set(ctx, composeKeys(shortURL.Short), shortURL.Original, 10 * time.Minute)
	if err != nil {
		log.Error("error set cache: "+ err.Error()+ " short: "+shortURL.Short)

		return err
	}

	return nil
}

func (u URLRepository) Get(ctx context.Context, shortURL string) (string, error) {
	value, err := u.cache.Get(ctx, composeKeys(shortURL))
	if err != nil {
		log.Error("error: "+ err.Error() +" getting cache url: "+ shortURL)
		return "", err
	}

	// get DB con circuit breaker

	return value, nil
}

func (u URLRepository) Delete(ctx context.Context, shortURL string) error {
	return u.cache.Del(ctx, composeKeys(shortURL))
}

func composeKeys(key string) string {
	return fmt.Sprintf("short:%s", key)
}