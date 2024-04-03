package items

import (
	"pmutils"
	"protos/parser"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ItemsClient struct {
	cl parser.ItemParserClient
}

func NewItemsClient() *ItemsClient {
	addr := pmutils.GetEnv("ITEM_WORKER_ADDR", "localhost:1111")
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := parser.NewItemParserClient(conn)
	cc := ItemsClient{cl: client}
	return &cc
}
