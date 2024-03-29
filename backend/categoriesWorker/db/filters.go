package db

import (
	"context"
	"encoding/json"
	"protos/parser"
)

func (d *Database) DBGetFilters(categoryUrl string, market parser.Markets) ([]*parser.Filter, error) {
	// SQL
	var filters []*parser.Filter
	longURL, err := d.MakeFullURL(categoryUrl, market)
	if err != nil {
		return nil, err
	}
	sqlStatement := `SELECT categoryFilters FROM Categories WHERE categoryURL=$1;`
	row := d.Conn.QueryRow(context.Background(), sqlStatement, longURL)
	// Unmarshalling JSON
	var filtersJSONB []byte
	err = row.Scan(&filtersJSONB)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(filtersJSONB, &filters)
	if err != nil {
		return nil, err
	}
	return filters, nil
}

func (d *Database) DBSaveFilters(filters []*parser.Filter, categoryUrl string, market parser.Markets) error {
	// Marhalling JSON
	filtersJSON, err := json.Marshal(&filters)
	if err != nil {
		return err
	}
	// SQL
	longURL, err := d.MakeFullURL(categoryUrl, market)
	if err != nil {
		return err
	}
	sqlStatement := `UPDATE Categories SET categoryFilters=$1 WHERE categoryURL=$2;`
	_, err = d.Conn.Exec(context.Background(), sqlStatement, filtersJSON, longURL)
	if err != nil {
		return err
	}
	return nil
}
