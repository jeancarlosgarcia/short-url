package service

import (
	"context"
	"errors"
	"testing"

	"short-url/interfaces/mocks"
	"short-url/models/api"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOK(t *testing.T) {
	repositoryMock := new(mocks.IUrlRepository)

	repositoryMock.On("Save", context.Background(), mock.AnythingOfType("*models.ShortURL")).
		Return(nil)

	service := NewURLService(repositoryMock)

	response := service.Create(context.Background(), &api.URLCreate{ Url: "any" })

	assert.Equal(t, response.Error, "")
}

func TestCreateError(t *testing.T) {
	repositoryMock := new(mocks.IUrlRepository)

	repositoryMock.On("Save", context.Background(), mock.AnythingOfType("*models.ShortURL")).
		Return(errors.New("any"))

	service := NewURLService(repositoryMock)

	response := service.Create(context.Background(), &api.URLCreate{ Url: "any" })

	assert.Equal(t, response.Error, "any")
}

func TestGetOriginalOK(t *testing.T) {
	repositoryMock := new(mocks.IUrlRepository)

	repositoryMock.On("Get", context.Background(), "shortURL").Return("longURL", nil)

	service := NewURLService(repositoryMock)

	url, err := service.GetOriginal(context.Background(), "shortURL")

	assert.NoError(t, err)
	assert.Equal(t, url, "longURL")
}

func TestGetOriginalError(t *testing.T) {
	repositoryMock := new(mocks.IUrlRepository)

	repositoryMock.On("Get", context.Background(), "shortURL").
		Return("", errors.New("any"))

	service := NewURLService(repositoryMock)

	_, err := service.GetOriginal(context.Background(), "shortURL")

	assert.Equal(t, err.Error(), "any")
}

func TestDeleteOK(t *testing.T) {
	repositoryMock := new(mocks.IUrlRepository)

	repositoryMock.On("Delete", context.Background(), "shortURL").Return(nil)

	service := NewURLService(repositoryMock)

	err := service.Delete(context.Background(), "shortURL")

	assert.NoError(t, err)
}

func TestDeleteError(t *testing.T) {
	repositoryMock := new(mocks.IUrlRepository)

	repositoryMock.On("Delete", context.Background(), "shortURL").Return(errors.New("any"))

	service := NewURLService(repositoryMock)

	err := service.Delete(context.Background(), "shortURL")

	assert.Equal(t, err.Error(), "any")
}