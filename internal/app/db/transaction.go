package db

import (
	"context"
	"go-server/internal/pkg/contextx"

	"gorm.io/gorm"
)

type Trans struct {
	DB *gorm.DB
}

func (t *Trans) Exec(ctx context.Context, fn func(context.Context) error) error {
	if _, ok := contextx.FromTrans(ctx); ok {
		return fn(ctx)
	}

	return t.DB.Transaction(func(db *gorm.DB) error {
		return fn(contextx.NewTrans(ctx, db))
	})
}
