package db

import (
	"context"
	"fmt"
)

func (d *Database) DBGetMarketURLs() (map[string]string, error) {
	// SQL
	rows, err := d.Conn.Query(context.Background(), "SELECT marketName, marketURL FROM Marketplaces")
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()
	// Saving to map
	marketLinks := make(map[string]string)
	for rows.Next() {
		var marketName, marketURL string
		err := rows.Scan(&marketName, &marketURL)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		marketLinks[marketName] = marketURL
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("error reading rows: %w", rows.Err())
	}
	return marketLinks, nil
}
