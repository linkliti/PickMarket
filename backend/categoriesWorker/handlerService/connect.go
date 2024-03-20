package handlerservice

import (
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
