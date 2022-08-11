package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"short-url/interfaces/mocks"
	"short-url/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateURL(t *testing.T) {
	shortURL := &models.ShortURL{ Short: "short", Original: "original"}

	storageMock := new(mocks.IStorage)
	cacheMock := new(mocks.ICache)

	storageMock.On("Insert", context.Background(), shortURL).Return(nil)
	cacheMock.On("Set", context.Background(), "short:short", "original", 10 * time.Minute).Return(nil)

	repository := NewUrlRepository(storageMock, cacheMock)

	err := repository.Save(context.Background(), shortURL)

	assert.NoError(t, err)
}

func TestCreateURLErrorStorage(t *testing.T) {
	errorExpected := "error storage"
	shortURL := &models.ShortURL{ Short: "short", Original: "original"}

	storageMock := new(mocks.IStorage)
	cacheMock := new(mocks.ICache)

	storageMock.On("Insert", context.Background(), shortURL).Return(errors.New(errorExpected))
	cacheMock.On("Set", context.Background(), "short:short", "original", 10 * time.Minute).Return(nil)

	repository := NewUrlRepository(storageMock, cacheMock)

	err := repository.Save(context.Background(), shortURL)

	assert.Equal(t, err.Error(), errorExpected)
}

func TestCreateURLErrorCache(t *testing.T) {
	errorExpected := "error cache"
	shortURL := &models.ShortURL{ Short: "short", Original: "original"}

	storageMock := new(mocks.IStorage)
	cacheMock := new(mocks.ICache)

	storageMock.On("Insert", context.Background(), shortURL).Return(nil)
	cacheMock.On("Set", context.Background(), "short:short", "original", 10 * time.Minute).
		Return(errors.New(errorExpected))

	repository := NewUrlRepository(storageMock, cacheMock)

	err := repository.Save(context.Background(), shortURL)

	assert.Equal(t, err.Error(), errorExpected)
}

func TestGetOriginal(t *testing.T) {
	longURLExpected := "longURL"
	cacheMock := new(mocks.ICache)
	cacheMock.On("Get", context.Background(), "short:shortUrl").Return(longURLExpected, nil)

	repository := NewUrlRepository(new(mocks.IStorage), cacheMock)

	longUrl, err := repository.Get(context.Background(), "shortUrl")

	assert.NoError(t, err)
	assert.Equal(t, longUrl, longURLExpected)
}

func TestGetOriginalError(t *testing.T) {
	errorExpected := "any"
	cacheMock := new(mocks.ICache)
	cacheMock.On("Get", context.Background(), "short:shortUrl").
		Return("", errors.New(errorExpected))

	repository := NewUrlRepository(new(mocks.IStorage), cacheMock)

	_, err := repository.Get(context.Background(), "shortUrl")

	assert.Equal(t, err.Error(), errorExpected)
}

func TestDelete(t *testing.T) {
	cacheMock := new(mocks.ICache)
	cacheMock.On("Del", context.Background(), "short:shortUrl").Return(nil)

	repository := NewUrlRepository(new(mocks.IStorage), cacheMock)

	err := repository.Delete(context.Background(), "shortUrl")

	assert.NoError(t, err)
}

func TestDeleteError(t *testing.T) {
	errorExpected := "any"
	cacheMock := new(mocks.ICache)
	cacheMock.On("Del", context.Background(), "short:shortUrl").Return(errors.New(errorExpected))

	repository := NewUrlRepository(new(mocks.IStorage), cacheMock)

	err := repository.Delete(context.Background(), "shortUrl")

	assert.Equal(t, err.Error(), errorExpected)
}