package repository

import (
	"context"

	"short-url/interfaces"

	"github.com/go-pg/pg/v10"
)

type StorageRepository struct {
	db *pg.DB
}

func NewStorageRepository(db *pg.DB) interfaces.IStorage {
	return StorageRepository{db}
}

func (s StorageRepository) Insert(_ context.Context, model interface{}) error {
	_, err := s.db.Model(model).Insert()

	return err
}