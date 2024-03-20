package handlerservice

import (
	"context"
	"io"
	"protos/parser"
)

func (c *CategoryService) GetCategoryFilters(req *parser.FiltersRequest, srv parser.CategoryParser_GetCategoryFiltersServer) error {
	client := c.connectToParser()
	stream, err := client.GetCategoryFilters(context.Background(), req)
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
