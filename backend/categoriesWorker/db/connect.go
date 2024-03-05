package db

import (
	"context"
	"pmutils"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	conn *pgxpool.Pool
}

func (d *Database) NewConnection() {
	var err error
	DB_NAME := pmutils.GetEnv("DB_NAME", "")
	DB_ADDR := pmutils.GetEnv("DB_ADDR", "")
	DB_CATEGORIES_USER := pmutils.GetEnv("DB_CATEGORIES_USER", "")
	DB_CATEGORIES_PASS := pmutils.GetEnv("DB_CATEGORIES_PASS", "")
	if pmutils.ContainsEmptyString(DB_NAME, DB_ADDR, DB_CATEGORIES_USER, DB_CATEGORIES_PASS) {
		panic("DB params cannot be empty")
	}
	url := "postgresql://" + DB_CATEGORIES_USER + ":" + DB_CATEGORIES_PASS + "@" + DB_ADDR + "/" + DB_NAME
	d.conn, err = pgxpool.Connect(context.Background(), url)
	if err != nil {
		panic(err)
	}
	defer d.conn.Close()
}
