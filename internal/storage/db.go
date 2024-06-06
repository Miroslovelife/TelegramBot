package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"goTgExample/internal/config"
)

func ConnectDB(cfgDb config.Database) (*pgxpool.Pool, error) {

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfgDb.User, cfgDb.Password, cfgDb.Host, cfgDb.Port, cfgDb.DbName, cfgDb.SslMode)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %w", err)

	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return pool, nil

}
