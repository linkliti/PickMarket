package service

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"protos/parser"
)

func (c *ItemsService) GetItemCharacteristics(req *parser.CharacteristicsRequest, srv parser.ItemParser_GetItemCharacteristicsServer) error {
	// Try to get characteristics from the database
	chars, err := c.db.DBGetChars(req.ItemUrl)
	if err != nil {
		// If it fails, get them from the parser
		stream, err := c.parsClient.GetItemCharacteristics(context.Background(), req)
		if err != nil {
			slog.Error("failed to get characteristics from parser", "err", err)
			return err
		}

		// Slice to hold the pointers to characteristics for saving to the database
		var charsToSave []*parser.Characteristic

		// Iterate over the stream
		for {
			charResponse, err := stream.Recv()
			if err == io.EOF {
				// Save the characteristics to the database after receiving all characteristics from the stream
				go func(charsToSave []*parser.Characteristic) {
					if err := c.db.DBSaveChars(charsToSave, req.ItemUrl, req.Market); err != nil {
						slog.Error("failed to save characteristics to database", "err", err)
					}
				}(charsToSave)
				break
			}
			if err != nil {
				slog.Error("failed to receive characteristic from stream", "err", err)
				return err
			}

			// Use a type assertion to get the Characteristic from the CharacteristicResponse
			if char, ok := charResponse.Message.(*parser.CharacteristicResponse_Characteristic); ok {
				// Create a new CharacteristicResponse to send to the caller
				resp := &parser.CharacteristicResponse{
					Message: char,
				}

				// Send the CharacteristicResponse to the caller
				if err := srv.Send(resp); err != nil {
					slog.Error("failed to send characteristic to caller", "err", err)
					return err
				}

				// Add the pointer to the characteristic to the slice
				charsToSave = append(charsToSave, char.Characteristic)
			} else {
				slog.Error("received a non-Characteristic message")
				return errors.New("received a non-Characteristic message")
			}
		}
	} else {
		// If getting characteristics from the database succeeds, send them to the caller
		for _, char := range chars {
			resp := &parser.CharacteristicResponse{
				Message: &parser.CharacteristicResponse_Characteristic{
					Characteristic: char,
				},
			}
			if err := srv.Send(resp); err != nil {
				slog.Error("failed to send characteristic to caller", "err", err)
				return err
			}
		}
	}

	return nil
}
