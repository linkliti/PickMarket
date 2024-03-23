package service

import (
	"context"
	"io"
	"protos/parser"
)

func (c *ItemsService) GetItems(req *parser.ItemsRequest, srv parser.ItemParser_GetItemsServer) error {
	stream, err := c.parsClient.GetItems(context.Background(), req)
	if err != nil {
		return err
	}

	// Loop over the stream and forward the responses to the original caller
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if err := srv.Send(res); err != nil {
			return err
		}
	}
	return nil
}
