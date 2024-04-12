package db

import (
	"context"
	"encoding/json"
	"protos/parser"

	"google.golang.org/protobuf/encoding/protojson"
)

func (d *Database) DBGetFilters(categoryUrl string, market parser.Markets) ([]*parser.Filter, error) {
	// SQL
	longURL, err := d.MakeFullURL(categoryUrl, market)
	if err != nil {
		return nil, err
	}
	sqlStatement := `SELECT categoryFilters FROM Categories WHERE categoryURL=$1;`
	row := d.Conn.QueryRow(context.Background(), sqlStatement, longURL)
	var filtersJSONB []byte
	err = row.Scan(&filtersJSONB)
	if err != nil {
		return nil, err
	}
	if len(filtersJSONB) == 0 {
		return []*parser.Filter{}, nil
	}
	// JSONB to json.RawMessage
	var filtersRaw []json.RawMessage
	err = json.Unmarshal(filtersJSONB, &filtersRaw)
	if err != nil {
		return nil, err
	}
	// json.RawMessage to parser.Filter
	filters := make([]*parser.Filter, len(filtersRaw))
	for i, v := range filtersRaw {
		filter := &parser.Filter{}
		err = protojson.Unmarshal(v, filter)
		if err != nil {
			return nil, err
		}
		filters[i] = filter
	}
	return filters, nil
}

func (d *Database) DBSaveFilters(filters []*parser.Filter, categoryUrl string, market parser.Markets) error {
	// parser.Filter to json.RawMessage
	filtersRaw := make([]json.RawMessage, len(filters))
	for i, filter := range filters {
		filterJSON, err := protojson.Marshal(filter)
		if err != nil {
			return err
		}
		filtersRaw[i] = json.RawMessage(filterJSON)
	}
	// json.RawMessage to JSONB
	filtersJSON, err := json.Marshal(filtersRaw)
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
