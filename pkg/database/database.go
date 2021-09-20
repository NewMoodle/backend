package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgreConnectionPool(host string, port int, username, password, database, sslmode string, pools int) (*pgxpool.Pool, error) {
	var err error
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s pool_max_conns=%d",
		username, password, host, port, database, sslmode, pools,
	)
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return &pgxpool.Pool{}, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), cfg)

	return pool, err
}
