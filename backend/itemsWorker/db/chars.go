package db

import (
	"context"
	"encoding/json"
	"protos/parser"
)

func (d *Database) DBGetChars(itemUrl string) ([]parser.Characteristic, error) {
	// Initialize a slice to hold the characteristics
	var characteristics []parser.Characteristic

	// Prepare the SQL query to retrieve characteristics
	sqlStatement := `SELECT itemChars FROM Items WHERE itemURL=$1;`

	// Query the database
	row := d.conn.QueryRow(context.Background(), sqlStatement, itemUrl)

	// Variable to hold the itemChars JSONB data
	var charsJSONB []byte

	// Scan the result into the charsJSONB variable
	err := row.Scan(&charsJSONB)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSONB data into a slice of parser.Characteristic
	err = json.Unmarshal(charsJSONB, &characteristics)
	if err != nil {
		return nil, err
	}

	// Return the slice of characteristics
	return characteristics, nil
}

func (d *Database) DBSaveChars(chars []parser.Characteristic, item *parser.Item, market string) error {
	// Marshal the characteristics into JSON format
	charsJSON, err := json.Marshal(chars)
	if err != nil {
		return err
	}

	// Prepare the SQL statement to insert or update the item
	sqlStatement := `
	INSERT INTO Items (itemURL, Marketplaces_marketName, itemName, itemChars, itemParseDate)
	VALUES ($1, $2, $3, $4, NOW())
	ON CONFLICT (itemURL) DO UPDATE
	SET Marketplaces_marketName = EXCLUDED.Marketplaces_marketName,
			itemName = EXCLUDED.itemName,
			itemChars = EXCLUDED.itemChars,
			itemParseDate = NOW();
	`

	// Execute the SQL statement
	_, err = d.conn.Exec(context.Background(), sqlStatement, &item.Url, market, &item.Name, charsJSON)
	if err != nil {
		return err
	}

	// Return nil if no errors occurred
	return nil
}
