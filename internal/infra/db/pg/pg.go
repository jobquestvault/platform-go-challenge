package pg

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jobquestvault/platform-go-challenge/internal/sys"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/cfg"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/errors"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/log"
)

type (
	DB struct {
		sys.Core
		db *sql.DB
	}
)

func NewDB(log log.Logger, cfg *cfg.Config) *DB {
	return &DB{
		Core: sys.NewCore(log, cfg),
		db:   nil,
	}
}

func (db *DB) Start(ctx context.Context) error {
	return db.Connect()
}

func (db *DB) Connect() error {
	pgdb, err := sql.Open("postgres", db.connString())
	if err != nil {
		msg := fmt.Sprintf("%s connection error", db.Cfg().DB.Name)
		return errors.Wrap(msg, err)
	}

	err = pgdb.Ping()
	if err != nil {
		msg := fmt.Sprintf("%s ping connection error", db.Cfg().DB.Name)
		return errors.Wrap(msg, err)
	}

	db.Log().Info("Database connected:", db.Cfg().DB.Name)

	db.db = pgdb

	return nil
}

func (db *DB) DB() any {
	return db.db
}

func (db *DB) connString() (connString string) {
	user := db.Cfg().DB.Username
	pass := db.Cfg().DB.Password
	name := db.Cfg().DB.Name
	host := db.Cfg().DB.Host
	port := db.Cfg().DB.Port
	schema := db.Cfg().DB.Schema

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d search_path=%s", user, pass, name, host, port, schema)

	if db.Cfg().DB.SSL {
		connStr = connStr + " sslmode=require"
	} else {
		connStr = connStr + " sslmode=disable"
	}

	return connStr
}
