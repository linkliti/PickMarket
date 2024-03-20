package handlerservice

import (
	"context"
	"io"
	"protos/parser"
)

func (c *ItemsService) GetItemCharacteristics(req *parser.CharacteristicsRequest, srv parser.ItemParser_GetItemCharacteristicsServer) error {
	client := c.connectToParser()
	stream, err := client.GetItemCharacteristics(context.Background(), req)
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
