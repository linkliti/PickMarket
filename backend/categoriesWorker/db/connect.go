package db

import (
	"context"
	"fmt"
	"pmutils"
	"protos/parser"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Conn        *pgxpool.Pool
	MarketLinks map[string]string
}

func NewDBConnection() (*Database, error) {
	var err error
	// ENV
	DB_NAME := pmutils.GetEnv("DB_NAME", "")
	DB_ADDR := pmutils.GetEnv("DB_ADDR", "")
	DB_CATEGORIES_USER := pmutils.GetEnv("DB_CATEGORIES_USER", "")
	DB_CATEGORIES_PASS := pmutils.GetEnv("DB_CATEGORIES_PASS", "")
	if pmutils.ContainsEmptyString(DB_NAME, DB_ADDR, DB_CATEGORIES_USER, DB_CATEGORIES_PASS) {
		return nil, fmt.Errorf("db params cannot be empty")
	}
	url := "postgresql://" + DB_CATEGORIES_USER + ":" + DB_CATEGORIES_PASS + "@" + DB_ADDR + "/" + DB_NAME
	// Database struct
	d := &Database{}
	d.Conn, err = pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	d.MarketLinks, err = d.DBGetMarketURLs()
	if err != nil {
		return nil, fmt.Errorf("failed to get market URLs: %w", err)
	}
	return d, nil
}

func (d *Database) MakeFullURL(url string, market parser.Markets) (string, error) {
	marketURL, ok := d.MarketLinks[market.String()]
	if !ok {
		return "", fmt.Errorf("market not found in MarketLinks")
	}
	return marketURL + url, nil
}

func (d *Database) MakeShortURL(url string, market parser.Markets) (string, error) {
	marketURL, ok := d.MarketLinks[market.String()]
	if !ok {
		return "", fmt.Errorf("market not found in MarketLinks")
	}
	return strings.TrimPrefix(url, marketURL), nil
}
