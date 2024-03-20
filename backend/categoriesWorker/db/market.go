package db

import (
	"context"
	"protos/parser"
)

func (d *DatabaseManager) DBGetMarketsWithEmptyUpdateTime() ([]parser.Markets, error) {
	var markets []parser.Markets
	sqlStatement := `SELECT marketName FROM Marketplaces WHERE marketParseDate IS NULL;`
	rows, err := d.conn.Query(context.Background(), sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var marketName parser.Markets
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

func (d *DatabaseManager) DBUpdateMarketUpdateTime(marketName parser.Markets) error {
	sqlStatement := `UPDATE Marketplaces SET marketParseDate = NOW() WHERE marketName = $1;`
	_, err := d.conn.Exec(context.Background(), sqlStatement, marketName)
	if err != nil {
		return err
	}
	return nil
}

func (d *DatabaseManager) DBSaveRootCategory(marketName parser.Markets, category *parser.Category) error {
	sqlStatement := `INSERT INTO Categories (Marketplaces_marketName, categoryName, categoryURL) VALUES ($1, $2, $3);`
	_, err := d.conn.Exec(context.Background(), sqlStatement, marketName, category.Title, category.Url)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) DBGetRootCategories() ([]*parser.Category, error) {
	// Initialize a slice to hold the root categories
	var categories []*parser.Category

	// Prepare the SQL query to retrieve root categories
	sqlStatement := `SELECT categoryName, categoryURL FROM Categories WHERE Marketplaces_marketName=$1 AND Categories_parentURL IS NULL;`

	// Query the database
	rows, err := d.conn.Query(context.Background(), sqlStatement, d.market)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var category parser.Category

		// Scan the result into the Category struct
		err := rows.Scan(&category.Title, &category.Url)
		if err != nil {
			return nil, err
		}

		// Append the pointer to the category to the slice
		categories = append(categories, &category)
	}

	// Check for any error that occurred during the iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Return the slice of root categories
	return categories, nil
}
