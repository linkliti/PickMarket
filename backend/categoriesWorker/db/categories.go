package db

import (
	"context"
	"database/sql"
	"protos/parser"
)

func (d *Database) DBGetCategoryChildren(categoryUrl string, market parser.Markets) ([]*parser.Category, error) {
	// Initialize a slice to hold pointers to the child categories
	var categories []*parser.Category
	// Prepare the SQL query to retrieve child categories
	sqlStatement := `SELECT categoryName, categoryURL, Categories_parentURL FROM Categories WHERE Categories_parentURL=$1;`
	// Query the database
	rows, err := d.Conn.Query(context.Background(), sqlStatement, categoryUrl)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Iterate through the result set
	for rows.Next() {
		var category parser.Category
		var parentURL sql.NullString
		// Scan the result into the Category struct
		err := rows.Scan(&category.Title, &category.Url, &parentURL)
		if err != nil {
			return nil, err
		}
		// If parentURL is valid, assign it to the Category struct
		if parentURL.Valid {
			category.ParentUrl = &parentURL.String
		}
		// Append the pointer to the category to the slice
		categories = append(categories, &category)
	}
	// Check for any error that occurred during the iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}
	// Return the slice of pointers to child categories
	return categories, nil
}

func (d *Database) DBGetRootCategoryChildren(market parser.Markets) ([]*parser.Category, error) {
	// Initialize a slice to hold the root category children
	var categories []*parser.Category
	// Prepare the SQL query to retrieve root category children
	sqlStatement := `SELECT categoryName, categoryURL FROM Categories WHERE Marketplaces_marketName=$1 AND Categories_parentURL IS NULL;`
	// Query the database
	rows, err := d.Conn.Query(context.Background(), sqlStatement, market.String())
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
	// Return the slice of root category children
	return categories, nil
}

func (d *Database) DBSaveCategory(category *parser.Category, market parser.Markets) error {
	// Prepare the SQL statement to insert or update the category
	sqlStatement := `
	INSERT INTO Categories (categoryName, categoryURL, Categories_parentURL, Marketplaces_marketName)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (categoryURL) DO UPDATE
	SET categoryName = EXCLUDED.categoryName,
			Categories_parentURL = EXCLUDED.Categories_parentURL
	`
	// Execute the SQL statement
	_, err := d.Conn.Exec(context.Background(), sqlStatement, &category.Title, &category.Url, &category.ParentUrl, market)
	if err != nil {
		return err
	}
	// Return nil if no errors occurred
	return nil
}
