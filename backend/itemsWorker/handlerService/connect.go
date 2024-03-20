package handlerservice

import (
	"pmutils"
	"protos/parser"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ItemsService struct {
	parser.UnimplementedItemParserServer // parser.UnsafeItemsParserServer to require all methods implementation
}

func NewItemsService() *ItemsService {
	return &ItemsService{}
}

func (c *ItemsService) connectToParser() parser.ItemParserClient {
	parserAddr := pmutils.GetEnv("PARSER_ADDR", "localhost:1111")
	conn, err := grpc.Dial(parserAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := parser.NewItemParserClient(conn)
	return client
}
