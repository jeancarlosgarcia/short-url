package postgres

import (
	"short-url/config"

	"github.com/go-pg/pg/v10"
)

func NewPostgresConnection() *pg.DB {
	dbConfig := config.Environments().DB

	db := pg.Connect(&pg.Options{
		Addr: dbConfig.Host,
		Database: dbConfig.Database,
		User: dbConfig.User,
		Password: dbConfig.Password,
	})

	return db
}