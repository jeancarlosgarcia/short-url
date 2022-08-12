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

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const prefix = "api/short-url"

var (
	serverHost = config.Environments().ServerHost
	serverPort = config.Environments().ServerPort
)

func main() {
	e := echo.New()

	router := e.Group(prefix)

	// Enable metrics middleware
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

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

	log.Infof(fmt.Sprintf("server is ready to handle requests %s:%d", serverHost, serverPort))
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", serverHost, serverPort)))
}