package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go_web_boilerplate/internal/config"
	"go_web_boilerplate/internal/pkg/logger"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBClient struct {
	Db     mongo.Database
	Logger logger.Logger
}

func NewClient(cfg config.Mongo, logger logger.Logger) (*DBClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn(cfg)))

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = db.Ping(ctx, readpref.Primary())

	if err != nil {
		return nil, err
	}

	return &DBClient{Db: *db.Database(cfg.DBName), Logger: logger}, nil
}

func (dbc DBClient) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := dbc.Db.Client().Disconnect(ctx); err != nil {
		return err
	}
	return nil
}

func dsn(cfg config.Mongo) string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port)
}
