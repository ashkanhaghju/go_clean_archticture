package postgres

import (
	"fmt"
	"go_web_boilerplate/internal/config"
	"go_web_boilerplate/internal/entity"
	"go_web_boilerplate/internal/pkg/logger"
	ormPostgres "gorm.io/driver/postgres"
	ormLogger "gorm.io/gorm/logger"

	"gorm.io/gorm"
)

type DB struct {
	Db     *gorm.DB
	Logger logger.Logger
}

func NewGorm(cfg config.Postgres, logger logger.Logger) (*DB, error) {
	db, err := gorm.Open(ormPostgres.Open(dsn(cfg)), &gorm.Config{
		Logger: ormLogger.Default.LogMode(ormLogger.Error),
	})

	if err != nil {
		return nil, err
	}

	migrationError := db.AutoMigrate(&entity.User{})

	if migrationError != nil {
		logger.Error(migrationError.Error())
	}
	return &DB{Db: db, Logger: logger}, nil
}

func (dbc DB) Disconnect() error {
	db, err := dbc.Db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func dsn(cfg config.Postgres) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran", cfg.Host, cfg.Username, cfg.Password, cfg.DBName, cfg.Port)
}
