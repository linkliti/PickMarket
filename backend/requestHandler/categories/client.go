package categories

import (
	"pmutils"
	"protos/parser"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CategoryClient struct {
	cl parser.CategoryParserClient
}

func NewCategoryClient() *CategoryClient {
	parserAddr := pmutils.GetEnv("CATEGORIES_WORKER_ADDR", "localhost:1111")
	conn, err := grpc.Dial(parserAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := parser.NewCategoryParserClient(conn)
	cc := CategoryClient{cl: client}
	return &cc
}
