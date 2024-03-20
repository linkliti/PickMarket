package db

import (
	"context"
	"fmt"
	"pmutils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	conn *pgxpool.Pool
}

func (d *Database) NewConnection() error {
	var err error
	DB_NAME := pmutils.GetEnv("DB_NAME", "")
	DB_ADDR := pmutils.GetEnv("DB_ADDR", "")
	DB_ITEMS_USER := pmutils.GetEnv("DB_ITEMS_USER", "")
	DB_ITEMS_PASS := pmutils.GetEnv("DB_ITEMS_PASS", "")
	if pmutils.ContainsEmptyString(DB_NAME, DB_ADDR, DB_ITEMS_USER, DB_ITEMS_PASS) {
		return fmt.Errorf("db params cannot be empty")
	}
	url := "postgresql://" + DB_ITEMS_USER + ":" + DB_ITEMS_PASS + "@" + DB_ADDR + "/" + DB_NAME
	d.conn, err = pgxpool.New(context.Background(), url)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer d.conn.Close()
	return nil
}
