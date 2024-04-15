package service

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"protos/parser"
)

func (c *ItemsService) grpcGetItems(req *parser.ItemsRequest) ([]*parser.Item, error) {
	// gRPC call
	stream, err := c.parsClient.GetItems(context.Background(), req)
	if err != nil {
		return nil, err
	}
	var items []*parser.Item
	for {
		response, err := stream.Recv()
		// End of stream
		if err == io.EOF {
			break
		}
		// Failed message
		if err != nil {
			return nil, err
		}
		// Message
		if item := response.GetItem(); item != nil {
			items = append(items, item)
		} else if status := response.GetStatus(); status != nil {
			slog.Warn("Received an error status", "status", status.Message)
			return nil, fmt.Errorf(status.Message)
		}
	}
	return items, nil
}
