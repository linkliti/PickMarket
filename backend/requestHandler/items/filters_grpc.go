package items

import (
	"context"
	"fmt"
	"io"
	"protos/parser"
)

func (c *ItemsClient) grpcGetFilters(req *parser.FiltersRequest) ([]*parser.Filter, error) {
	// gRPC call
	stream, err := c.cl.GetCategoryFilters(context.Background(), req)
	if err != nil {
		return nil, err
	}
	var filters []*parser.Filter
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
		if filter := response.GetFilter(); filter != nil {
			filters = append(filters, filter)
		} else if status := response.GetStatus(); status != nil {
			fmt.Printf("Received an error status: %v\n", status)
			return nil, fmt.Errorf(status.Message)
		}
	}
	return filters, nil
}
