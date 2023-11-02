package config

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgresqlDB(connStr string, timeout time.Duration) (*sql.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	conn, err := db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	if conn.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}

func ClosePostgresqlDB(db *sql.DB) {
	db.Close()
}
