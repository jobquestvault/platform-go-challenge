package db

import (
	"context"

	"github.com/jobquestvault/platform-go-challenge/internal/sys"
)

type (
	DB interface {
		sys.Core
		DB() any
		Tx
	}

	Tx interface {
		Begin(ctx context.Context) (TxContext, error)
	}

	TxContext interface {
		Commit(ctx context.Context) error
		Rollback(ctx context.Context) error
	}
)
