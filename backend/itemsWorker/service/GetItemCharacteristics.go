package service

import (
	"context"
	"io"
	"log/slog"
	"protos/parser"

	"google.golang.org/grpc/codes"

	statuspb "google.golang.org/genproto/googleapis/rpc/status"
)

func (c *ItemsService) GetItemCharacteristics(req *parser.CharacteristicsRequest, srv parser.ItemParser_GetItemCharacteristicsServer) error {
	// Get Chars from DB
	slog.Debug("Incoming GetItemCharacteristics request", "request", req)
	chars, err := c.db.DBGetChars(req.ItemUrl, req.Market)
	if err != nil {
		errText := "failed to get characteristics from database"
		slog.Error(errText, "err", err)
		return sendErrorStatus_GetItemCharacteristics(srv, errText)
	}
	// Get from parser
	slog.Debug("Sending chars from parser", "request", req)
	if len(chars) == 0 {
		stream, err := c.parsClient.GetItemCharacteristics(context.Background(), req)
		if err != nil {
			errText := "failed to get characteristics from parser"
			slog.Error(errText, "err", err)
			return sendErrorStatus_GetItemCharacteristics(srv, errText)
		}
		var charsToSave []*parser.Characteristic
		for {
			charResponse, err := stream.Recv()
			// Save final array to DB
			if err == io.EOF {
				go func() {
					if err := c.db.DBSaveChars(charsToSave, req.ItemUrl, req.Market); err != nil {
						slog.Error("failed to save characteristics to database", "err", err)
						return
					}
				}()
				break
			}
			// Failed message
			if err != nil {
				errText := "failed to receive characteristic from stream"
				slog.Error(errText, "err", err)
				return sendErrorStatus_GetItemCharacteristics(srv, errText)
			}
			// Message
			if char, ok := charResponse.Message.(*parser.CharacteristicResponse_Char); ok {
				resp := &parser.CharacteristicResponse{
					Message: char,
				}
				if err := srv.Send(resp); err != nil {
					slog.Error("failed to send characteristic to caller", "err", err)
					return err
				}
				charsToSave = append(charsToSave, char.Char)
			} else {
				errText := "received a non-characteristic message"
				slog.Error(errText, "message", charResponse)
				return sendErrorStatus_GetItemCharacteristics(srv, errText)
			}
		}
	} else {
		// Send chars from DB
		slog.Debug("Sending chars from database", "request", req)
		for _, char := range chars {
			resp := &parser.CharacteristicResponse{
				Message: &parser.CharacteristicResponse_Char{
					Char: char,
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

func sendErrorStatus_GetItemCharacteristics(srv parser.ItemParser_GetItemCharacteristicsServer, errText string) error {
	resp := &parser.CharacteristicResponse{
		Message: &parser.CharacteristicResponse_Status{
			Status: &statuspb.Status{
				Code:    int32(codes.Internal),
				Message: errText,
			},
		},
	}
	return srv.Send(resp)
}
