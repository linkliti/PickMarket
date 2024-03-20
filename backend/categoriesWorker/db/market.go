package db

import (
	"context"
)

func (d *Database) DBGetMarketsWithEmptyUpdateTime() ([]string, error) {
	var markets []string
	sqlStatement := `SELECT marketName FROM Marketplaces WHERE marketParseDate IS NULL;`
	rows, err := d.conn.Query(context.Background(), sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var marketName string
		if err := rows.Scan(&marketName); err != nil {
			return nil, err
		}
		markets = append(markets, marketName)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return markets, nil
}

func (d *Database) DBUpdateMarketUpdateTime(marketName string) error {
	sqlStatement := `UPDATE Marketplaces SET marketParseDate = NOW() WHERE marketName = $1;`
	_, err := d.conn.Exec(context.Background(), sqlStatement, marketName)
	if err != nil {
		return err
	}
	return nil
}
