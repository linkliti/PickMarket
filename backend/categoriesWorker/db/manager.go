package db

import (
	"categoriesWorker/util"
	"context"
	"protos/parser"
)

func (d *Database) DBGetMarketsWithCategoriesWithoutParseDate() ([]parser.Markets, error) {
	// SQL
	var markets []parser.Markets
	sqlStatement := `SELECT DISTINCT Marketplaces_marketName FROM Categories WHERE categoryParseDate IS NULL;`
	rows, err := d.Conn.Query(context.Background(), sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Saving to parser.Markets
	for rows.Next() {
		var marketStr string
		if err := rows.Scan(&marketStr); err != nil {
			return nil, err
		}
		marketEnum, err := util.StringToMarket(marketStr)
		if err != nil {
			return nil, err
		}
		markets = append(markets, marketEnum)
	}
	return markets, nil
}

func (d *Database) DBGetMarketsWithoutParseDate() ([]parser.Markets, error) {
	// SQL
	var markets []parser.Markets
	sqlStatement := `SELECT marketName FROM Marketplaces WHERE marketParseDate IS NULL;`
	rows, err := d.Conn.Query(context.Background(), sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Saving to parser.Markets
	for rows.Next() {
		var marketStr string
		if err := rows.Scan(&marketStr); err != nil {
			return nil, err
		}
		marketEnum, err := util.StringToMarket(marketStr)
		if err != nil {
			return nil, err
		}
		markets = append(markets, marketEnum)
	}
	return markets, nil
}

func (d *Database) DBSetMarketParseDate(market parser.Markets) error {
	// SQL
	marketStr := util.MarketToString(market)
	sqlStatement := `UPDATE Marketplaces SET marketParseDate=NOW() WHERE marketName=$1;`
	_, err := d.Conn.Exec(context.Background(), sqlStatement, marketStr)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) DBGetCategoriesWithoutParseDate(market parser.Markets) ([]string, error) {
	// SQL
	var categories []string
	sqlStatement := `SELECT categoryURL FROM Categories WHERE categoryParseDate IS NULL AND Marketplaces_marketName=$1;`
	rows, err := d.Conn.Query(context.Background(), sqlStatement, market.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Saving category urls to []string
	for rows.Next() {
		var categoryUrl string
		if err := rows.Scan(&categoryUrl); err != nil {
			return nil, err
		}
		shortURL, err := d.MakeShortURL(categoryUrl, market)
		if err != nil {
			return nil, err
		}
		categories = append(categories, shortURL)
	}
	return categories, nil
}

func (d *Database) DBSetCategoryParseDate(categoryUrl string, market parser.Markets) error {
	// SQL
	longURL, err := d.MakeFullURL(categoryUrl, market)
	if err != nil {
		return err
	}
	sqlStatement := `UPDATE Categories SET categoryParseDate=NOW() WHERE categoryURL=$1;`
	_, err = d.Conn.Exec(context.Background(), sqlStatement, longURL)
	if err != nil {
		return err
	}
	return nil
}
