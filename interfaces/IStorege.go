package interfaces

import "context"

type IStorage interface {
	Insert(ctx context.Context, model interface{}) error
}
