package interfaces

import (
	"context"
	"time"
)

type ICache interface {
	Set(ctx context.Context, key string, value interface{}, exp time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}