package global

import (
	"context"
	"github.com/TimIrwing/nyashka-butler/internal/mongodb"
)

type Global struct {
	Context context.Context
	DB      *mongodb.DB
}

func New(ctx context.Context, db *mongodb.DB) *Global {
	g := &Global{
		ctx,
		db,
	}
	return g
}
