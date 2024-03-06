package handlerservice

import (
	"context"
	"io"
	"pmutils"
	"protos/parser"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CategoryService struct {
	parser.UnimplementedCategoryParserServer // parser.UnsafeCategoryParserServer to require all methods implementation
}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func (c *CategoryService) connectToParser() parser.CategoryParserClient {
	parserAddr := pmutils.GetEnv("PARSER_ADDR", "localhost:1111")
	conn, err := grpc.Dial(parserAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := parser.NewCategoryParserClient(conn)
	return client
}

func (c *CategoryService) GetRootCategories(req *parser.RootCategoriesRequest, srv parser.CategoryParser_GetRootCategoriesServer) error {
	client := c.connectToParser()
	stream, err := client.GetRootCategories(context.Background(), req)
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

func (c *CategoryService) GetSubCategories(req *parser.SubCategoriesRequest, srv parser.CategoryParser_GetSubCategoriesServer) error {
	client := c.connectToParser()
	stream, err := client.GetSubCategories(context.Background(), req)
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
