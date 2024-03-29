package db

import (
	"context"
	"protos/parser"
)

func (d *Database) DBGetCategoryChildren(parentUrl string, market parser.Markets) ([]*parser.Category, error) {
	// SQL
	var categories []*parser.Category
	longParentURL, err := d.MakeFullURL(parentUrl, market)
	if err != nil {
		return nil, err
	}
	sqlStatement := `SELECT categoryName, categoryURL FROM Categories WHERE Categories_parentURL=$1;`
	rows, err := d.Conn.Query(context.Background(), sqlStatement, longParentURL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Saving to parser.Category
	for rows.Next() {
		var category parser.Category
		err := rows.Scan(&category.Title, &category.Url)
		if err != nil {
			return nil, err
		}
		category.ParentUrl = &parentUrl
		shortURL, err := d.MakeShortURL(category.Url, market)
		if err != nil {
			return nil, err
		}
		category.Url = shortURL
		categories = append(categories, &category)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

func (d *Database) DBGetRootCategoryChildren(market parser.Markets) ([]*parser.Category, error) {
	// SQL
	var categories []*parser.Category
	sqlStatement := `SELECT categoryName, categoryURL FROM Categories WHERE Marketplaces_marketName=$1 AND Categories_parentURL IS NULL;`
	rows, err := d.Conn.Query(context.Background(), sqlStatement, market.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Saving to parser.Category
	for rows.Next() {
		var category parser.Category
		err := rows.Scan(&category.Title, &category.Url)
		if err != nil {
			return nil, err
		}
		shortURL, err := d.MakeShortURL(category.Url, market)
		if err != nil {
			return nil, err
		}
		category.Url = shortURL
		categories = append(categories, &category)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

func (d *Database) DBSaveCategory(category *parser.Category, market parser.Markets) error {
	// Long URLs
	longURL, err := d.MakeFullURL(category.Url, market)
	if err != nil {
		return err
	}
	longURLParent, err := d.MakeFullURL(*category.ParentUrl, market)
	if err != nil {
		return err
	}
	// SQL
	sqlStatement := `
	INSERT INTO Categories (categoryName, categoryURL, Categories_parentURL, Marketplaces_marketName)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (categoryURL) DO UPDATE
	SET categoryName = EXCLUDED.categoryName,
			Categories_parentURL = EXCLUDED.Categories_parentURL
	`
	_, err = d.Conn.Exec(context.Background(), sqlStatement, &category.Title, longURL, longURLParent, market)
	if err != nil {
		return err
	}
	return nil
}
