package db

import (
	"context"
	"encoding/json"
	"protos/parser"

	"google.golang.org/protobuf/encoding/protojson"
)

func (d *Database) DBGetChars(itemUrl string) ([]*parser.Characteristic, error) {
	// SQL
	sqlStatement := `SELECT itemChars FROM Items WHERE itemURL=$1;`
	row := d.conn.QueryRow(context.Background(), sqlStatement, itemUrl)
	var charsJSONB []byte
	err := row.Scan(&charsJSONB)
	if err != nil {
		return nil, err
	}
	if len(charsJSONB) == 0 {
		return []*parser.Characteristic{}, nil
	}
	// JSONB to json.RawMessage
	var charsRaw []json.RawMessage
	err = json.Unmarshal(charsJSONB, &charsRaw)
	if err != nil {
		return nil, err
	}
	// json.RawMessage to parser.Characteristic
	chars := make([]*parser.Characteristic, len(charsRaw))
	for i, v := range charsRaw {
		char := &parser.Characteristic{}
		err = protojson.Unmarshal(v, char)
		if err != nil {
			return nil, err
		}
		chars[i] = char
	}
	return chars, nil
}

func (d *Database) DBSaveChars(chars []*parser.Characteristic, itemUrl string, market parser.Markets) error {
	// parser.Characteristic to json.RawMessage
	charsRaw := make([]json.RawMessage, len(chars))
	for i, char := range chars {
		charJSON, err := protojson.Marshal(char)
		if err != nil {
			return err
		}
		charsRaw[i] = json.RawMessage(charJSON)
	}
	// json.RawMessage to JSONB
	charsJSON, err := json.Marshal(charsRaw)
	if err != nil {
		return err
	}
	// SQL
	sqlStatement := `
	INSERT INTO Items (itemURL, Marketplaces_marketName, itemChars, itemParseDate)
	VALUES ($1, $2, $3, NOW())
	ON CONFLICT (itemURL) DO UPDATE
	SET itemChars = EXCLUDED.itemChars,
			itemParseDate = NOW();
	`
	_, err = d.conn.Exec(context.Background(), sqlStatement, itemUrl, market, charsJSON)
	if err != nil {
		return err
	}
	return nil
}
