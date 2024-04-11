package items

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"protos/parser"
)

func (c *ItemsClient) grpcGetCharacteristics(req *parser.CharacteristicsRequest) ([]*parser.Characteristic, error) {
	// gRPC call
	stream, err := c.cl.GetItemCharacteristics(context.Background(), req)
	if err != nil {
		return nil, err
	}
	var chars []*parser.Characteristic
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
		if char := response.GetCharacteristic(); char != nil {
			chars = append(chars, char)
		} else if status := response.GetStatus(); status != nil {
			slog.Warn("Received an error status", "status", status.Message)
			return nil, fmt.Errorf(status.Message)
		}
	}
	return chars, nil
}
