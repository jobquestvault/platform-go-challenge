package pg

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/jobquestvault/platform-go-challenge/internal/infra/db"
	"github.com/jobquestvault/platform-go-challenge/internal/sys"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/cfg"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/errors"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/log"
)

type (
	AssetRepo struct {
		*sys.BaseCore
		db db.DB
	}
)

func NewAssetRepo(db db.DB, log log.Logger, cfg *cfg.Config) *AssetRepo {
	return &AssetRepo{
		BaseCore: sys.NewCore(log, cfg),
		db:       db,
	}
}

func (ar *AssetRepo) Start(ctx context.Context) error {
	err := ar.db.Start(ctx)
	if err != nil {
		msg := fmt.Sprintf("%s setup error", err)
		return errors.Wrap(msg, err)
	}

	return nil
}

func (ar *AssetRepo) DB() (db any) {
	return ar.db.DB()
}

func (ar *AssetRepo) PgDB() (db *sql.DB, ok bool) {
	db, ok = ar.DB().(*sql.DB)
	if !ok {
		return db, false
	}

	return db, true
}
