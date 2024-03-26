package db

import (
	"context"
	"fmt"
	"pmutils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Conn *pgxpool.Pool
}

func NewDBConnection() (*Database, error) {
	var err error
	DB_NAME := pmutils.GetEnv("DB_NAME", "")
	DB_ADDR := pmutils.GetEnv("DB_ADDR", "")
	DB_CATEGORIES_USER := pmutils.GetEnv("DB_CATEGORIES_USER", "")
	DB_CATEGORIES_PASS := pmutils.GetEnv("DB_CATEGORIES_PASS", "")
	if pmutils.ContainsEmptyString(DB_NAME, DB_ADDR, DB_CATEGORIES_USER, DB_CATEGORIES_PASS) {
		return nil, fmt.Errorf("db params cannot be empty")
	}
	url := "postgresql://" + DB_CATEGORIES_USER + ":" + DB_CATEGORIES_PASS + "@" + DB_ADDR + "/" + DB_NAME
	d := &Database{}
	d.Conn, err = pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return d, nil
}
