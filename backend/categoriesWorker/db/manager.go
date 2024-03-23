package db

import (
	"categoriesWorker/util"
	"context"
	"protos/parser"
)

func (d *Database) DBGetMarketsWithoutParseDate() ([]parser.Markets, error) {
	var markets []parser.Markets
	sqlStatement := `SELECT marketName FROM Marketplaces WHERE marketParseDate IS NULL;`
	rows, err := d.conn.Query(context.Background(), sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var marketStr string
		if err := rows.Scan(&marketStr); err != nil {
			return nil, err
		}
		// Convert the market string to a parser.Markets enum value
		marketEnum, err := util.StringToMarket(marketStr)
		if err != nil {
			return nil, err
		}
		markets = append(markets, marketEnum)
	}
	return markets, nil
}

func (d *Database) DBSetMarketParseDate(market parser.Markets) error {
	// Convert the market to a string
	marketStr := util.MarketToString(market)
	// Prepare the SQL statement to update the market's parseDate
	sqlStatement := `UPDATE Marketplaces SET marketParseDate=NOW() WHERE marketName=$1;`
	// Execute the SQL statement
	_, err := d.conn.Exec(context.Background(), sqlStatement, marketStr)
	if err != nil {
		return err
	}
	// Return nil if no errors occurred
	return nil
}

func (d *Database) DBGetCategoriesWithoutParseDate(market parser.Markets) ([]string, error) {
	var categories []string
	sqlStatement := `SELECT categoryURL FROM Categories WHERE categoryParseDate IS NULL AND Marketplaces_marketName=$1;`
	rows, err := d.conn.Query(context.Background(), sqlStatement, market.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var categoryUrl string
		if err := rows.Scan(&categoryUrl); err != nil {
			return nil, err
		}
		categories = append(categories, categoryUrl)
	}
	return categories, nil
}

func (d *Database) DBSetCategoryParseDate(categoryUrl string) error {
	// Prepare the SQL statement to update the category's parseDate
	sqlStatement := `UPDATE Categories SET categoryParseDate=NOW() WHERE categoryURL=$1;`
	// Execute the SQL statement
	_, err := d.conn.Exec(context.Background(), sqlStatement, categoryUrl)
	if err != nil {
		return err
	}
	// Return nil if no errors occurred
	return nil
}
