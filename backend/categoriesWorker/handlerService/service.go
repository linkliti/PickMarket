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

func (c *CategoryService) connectToParser() *grpc.ClientConn {
	parserAddr := pmutils.GetEnv("PARSER_ADDR", "localhost:1111")
	conn, err := grpc.Dial(parserAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return conn
}

func (c *CategoryService) GetRootCategories(req *parser.RootCategoriesRequest, srv parser.CategoryParser_GetRootCategoriesServer) error {
	conn := c.connectToParser()
	client := parser.NewCategoryParserClient(conn)
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

// func (c *CategoryService) GetSubCategories(*parser.SubCategoriesRequest, parser.CategoryParser_GetSubCategoriesServer) error {

// }

// func (c *CategoryService) GetCategoryFilters(*parser.FiltersRequest, parser.CategoryParser_GetCategoryFiltersServer) error {

// }
