package main

import (
	"fmt"

	"short-url/component"
	"short-url/config"
	"short-url/connection/postgres"
	"short-url/connection/redis"
	"short-url/handler"
	"short-url/repository"
	"short-url/service"

	"github.com/labstack/echo/v4"
)

const prefix = "api/short-url"

var (
	serverHost = config.Environments().ServerHost
	serverPort = config.Environments().ServerPort
)

func main() {
	e := echo.New()
	router := e.Group(prefix)

	health := component.NewHealth(handler.NewHealthCheckHandler(), router)
	health.Routes()

	postgres := postgres.NewPostgresConnection()
	redis := redis.NewRedisConnection()

	cache := repository.NewCacheRepository(redis)
	storage := repository.NewStorageRepository(postgres)
	urlRepository := repository.NewUrlRepository(storage, cache)

	urlService := service.NewURLService(urlRepository)
	url := component.NewUrl(handler.NewURLHandler(urlService), router)
	url.Routes()

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", serverHost, serverPort)))
}