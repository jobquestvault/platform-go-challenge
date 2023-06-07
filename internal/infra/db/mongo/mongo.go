package mongo

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/jobquestvault/platform-go-challenge/internal/sys"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/cfg"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/errors"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/log"
)

type (
	DB struct {
		sys.Core
		client *mongo.Client
		db     *mongo.Database
	}
)

const (
	name = "mongo-db"
)

func NewDB(log log.Logger, cfg *cfg.Config) *DB {
	return &DB{
		Core:   sys.NewCore(log, cfg),
		client: nil,
	}
}

func (db *DB) Start(ctx context.Context) error {
	return db.Connect()
}

func (db *DB) Connect() (err error) {
	connString := db.connString()
	db.client, err = mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		return errors.Wrap("MongoDB connect client error", err)
	}

	ctx := context.TODO()

	err = db.client.Connect(ctx)
	if err != nil {
		return errors.Wrap("MongoDB connect error", err)
	}

	dbName := db.Cfg().DB.Name
	db.db = db.client.Database(dbName)

	db.Log().Info("database connected:", db.Cfg().DB.Name)
	return nil
}

func (db *DB) DB() any {
	return db.db
}

func (db *DB) mongoDB() (sqlDB *sqlx.DB, ok bool) {
	sqlDB, ok = db.DB().(*sqlx.DB)
	if !ok {
		return sqlDB, false
	}

	return sqlDB, true
}

func (db *DB) connString() (connString string) {
	user := db.Cfg().DB.Username
	pass := db.Cfg().DB.Password
	name := db.Cfg().DB.Name
	host := db.Cfg().DB.Host
	port := db.Cfg().DB.Port
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", user, pass, host, port, name)
}
