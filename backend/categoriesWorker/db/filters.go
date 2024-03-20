package db

import (
	"context"
	"encoding/json"
	"protos/parser"
)

func (d *Database) DBGetFilters(categoryUrl string) ([]parser.Filter, error) {
	// Initialize a slice to hold the filters
	var filters []parser.Filter

	// Prepare the SQL query to retrieve filters
	sqlStatement := `SELECT categoryFilters FROM Categories WHERE categoryURL=$1;`

	// Query the database
	row := d.conn.QueryRow(context.Background(), sqlStatement, categoryUrl)

	// Variable to hold the filters JSONB data
	var filtersJSONB []byte

	// Scan the result into the filtersJSONB variable
	err := row.Scan(&filtersJSONB)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSONB data into a slice of parser.Filter
	err = json.Unmarshal(filtersJSONB, &filters)
	if err != nil {
		return nil, err
	}

	// Return the slice of filters
	return filters, nil
}

func (d *Database) DBSaveFilters(filters []parser.Filter, categoryUrl string) error {
	// Marshal the filters into JSON format
	filtersJSON, err := json.Marshal(filters)
	if err != nil {
		return err
	}

	// Prepare the SQL statement to update the filters
	sqlStatement := `UPDATE Categories SET categoryFilters=$1 WHERE categoryURL=$2;`

	// Execute the SQL statement
	_, err = d.conn.Exec(context.Background(), sqlStatement, filtersJSON, categoryUrl)
	if err != nil {
		return err
	}

	// Return nil if no errors occurred
	return nil
}
