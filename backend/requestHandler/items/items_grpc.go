package items

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"protos/parser"
	"sync"
)

func (c *ItemsClient) grpcGetItems(req *parser.ItemsRequest) ([]*parser.ItemExtended, error) {
	// gRPC call
	stream, err := c.cl.GetItems(context.Background(), req)
	if err != nil {
		return nil, err
	}
	var items []*parser.ItemExtended
	var mutex sync.Mutex
	var wg sync.WaitGroup
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
			wg.Add(1)
			go func(item *parser.Item) {
				// Create ItemExtended with chars and without weight
				defer wg.Done()
				charsReq := &parser.CharacteristicsRequest{
					Market:  req.Market,
					ItemUrl: item.Url,
				}
				chars, err := c.grpcGetCharacteristics(charsReq)
				if err != nil {
					slog.Warn("Failed to get characteristics", "error", err, "item", item)
					return
				}
				itemExt := &parser.ItemExtended{
					Item:  item,
					Chars: chars,
				}
				mutex.Lock()
				items = append(items, itemExt)
				mutex.Unlock()
			}(item)
		} else if status := response.GetStatus(); status != nil {
			slog.Warn("Received an error status", "status", status.Message)
			return nil, fmt.Errorf(status.Message)
		}
	}
	wg.Wait()
	return items, nil
}
