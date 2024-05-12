package service

import (
	"context"
	"io"
	"log/slog"
	"protos/parser"

	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
)

func (c *ItemsService) GetItems(req *parser.ItemsRequest, srv parser.ItemParser_GetItemsServer) error {
	// Get from parser
	slog.Debug("Incoming GetItems request", "request", req)
	stream, err := c.parsClient.GetItems(context.Background(), req)
	if err != nil {
		errText := "failed to get items from parser"
		slog.Error(errText, "err", err)
		return sendErrorStatus_GetItems(srv, errText)
	}
	slog.Debug("Sending items from parser", "request", req)
	for {
		itemResponse, err := stream.Recv()
		// Break on final item
		if err == io.EOF {
			break
		}
		// Failed message
		if err != nil {
			errText := "failed to receive items from stream"
			slog.Error(errText, "err", err)
			return sendErrorStatus_GetItems(srv, errText)
		}
		// Message
		if item, ok := itemResponse.Message.(*parser.ItemResponse_Item); ok {
			resp := &parser.ItemResponse{
				Message: item,
			}
			if err := srv.Send(resp); err != nil {
				slog.Error("failed to send item to caller", "err", err)
				return err
			}
		} else {
			errText := "received a non-item message"
			slog.Error(errText, "message", itemResponse)
			return sendErrorStatus_GetItems(srv, errText)
		}
	}
	return nil
}

func sendErrorStatus_GetItems(srv parser.ItemParser_GetItemsServer, errText string) error {
	resp := &parser.ItemResponse{
		Message: &parser.ItemResponse_Status{
			Status: &statuspb.Status{
				Code:    int32(codes.Internal),
				Message: errText,
			},
		},
	}
	return srv.Send(resp)
}
