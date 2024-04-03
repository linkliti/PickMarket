package db

import (
	"context"
	"encoding/json"
	"protos/parser"

	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/encoding/protojson"
)

func (d *Database) DBGetChars(longItemURL string, market parser.Markets) ([]*parser.Characteristic, error) {
	// SQL
	longURL, err := d.MakeFullURL(longItemURL, market)
	if err != nil {
		return nil, err
	}
	sqlStatement := `SELECT itemChars FROM Items WHERE itemURL=$1;`
	row := d.Conn.QueryRow(context.Background(), sqlStatement, longURL)
	var charsJSONB []byte
	err = row.Scan(&charsJSONB)
	if err != nil {
		if err == pgx.ErrNoRows {
			return []*parser.Characteristic{}, nil
		} else {
			return nil, err
		}
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
	longItemURL, err := d.MakeFullURL(itemUrl, market)
	if err != nil {
		return err
	}
	sqlStatement := `
	INSERT INTO Items (itemURL, Marketplaces_marketName, itemChars, itemParseDate)
	VALUES ($1, $2, $3, NOW())
	ON CONFLICT (itemURL) DO UPDATE
	SET itemChars = EXCLUDED.itemChars,
			itemParseDate = NOW();
	`
	_, err = d.Conn.Exec(context.Background(), sqlStatement, longItemURL, market, charsJSON)
	if err != nil {
		return err
	}
	return nil
}
