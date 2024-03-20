package db

import (
	"context"
	"database/sql"
	"protos/parser"
)

func (d *Database) DBGetCategoryChildren(categoryUrl string) ([]*parser.Category, error) {
	// Initialize a slice to hold pointers to the child categories
	var categories []*parser.Category

	// Prepare the SQL query to retrieve child categories
	sqlStatement := `SELECT categoryName, categoryURL, Categories_parentURL FROM Categories WHERE Categories_parentURL=$1;`

	// Query the database
	rows, err := d.conn.Query(context.Background(), sqlStatement, categoryUrl)
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

func (d *Database) DBSaveCategory(category *parser.Category) error {
	// Prepare the SQL statement to insert or update the category
	sqlStatement := `
	INSERT INTO Categories (categoryName, categoryURL, Categories_parentURL, Marketplaces_marketName)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (categoryURL) DO UPDATE
	SET categoryName = EXCLUDED.categoryName,
			Categories_parentURL = EXCLUDED.Categories_parentURL
	`

	// Execute the SQL statement
	_, err := d.conn.Exec(context.Background(), sqlStatement, &category.Title, &category.Url, &category.ParentUrl, d.market)
	if err != nil {
		return err
	}

	// Return nil if no errors occurred
	return nil
}

func (d *DatabaseManager) DBSaveCategory(category *parser.Category, market parser.Markets) error {
	// Prepare the SQL statement to insert or update the category
	sqlStatement := `
	INSERT INTO Categories (categoryName, categoryURL, Categories_parentURL, Marketplaces_marketName)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (categoryURL) DO UPDATE
	SET categoryName = EXCLUDED.categoryName,
			Categories_parentURL = EXCLUDED.Categories_parentURL
	`

	// Execute the SQL statement
	_, err := d.conn.Exec(context.Background(), sqlStatement, &category.Title, &category.Url, &category.ParentUrl, market)
	if err != nil {
		return err
	}

	// Return nil if no errors occurred
	return nil
}

func (d *DatabaseManager) DBSetCategoryUpdateTime(categoryUrl string) error {
	// Prepare the SQL statement to update the category's update time
	sqlStatement := `UPDATE Categories SET categoryParseDate=NOW() WHERE categoryURL=$1;`

	// Execute the SQL statement
	_, err := d.conn.Exec(context.Background(), sqlStatement, categoryUrl)
	if err != nil {
		return err
	}

	return nil
}

func (d *DatabaseManager) DBGetCategoriesWithEmptyParseDate() ([]*parser.SubCategoriesRequest, error) {
	// Initialize a slice to hold the categories
	var categories []*parser.SubCategoriesRequest

	// Prepare the SQL query to retrieve categories with an empty ParseDate
	sqlStatement := `SELECT categoryURL, Marketplaces_marketName FROM Categories WHERE categoryParseDate IS NULL;`

	// Query the database
	rows, err := d.conn.Query(context.Background(), sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var category parser.SubCategoriesRequest

		// Scan the result into the Category struct
		err := rows.Scan(&category.CategoryUrl, &category.Market)
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

	// Return the slice of categories
	return categories, nil
}
