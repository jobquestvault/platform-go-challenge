package pg

import (
	"context"
	"errors"

	"github.com/jobquestvault/platform-go-challenge/internal/infra/db"
)

type (
	ctxKey string
)

const (
	TxKey = ctxKey("tx")
)

var (
	TxNotFoundError = errors.New("no transaction found in context")
)

func (ar *AssetRepo) Tx(ctx context.Context) (db.Tx, error) {
	tx, ok := ctx.Value(TxKey).(db.Tx)
	if !ok {
		return nil, TxNotFoundError
	}

	return tx, nil
}
