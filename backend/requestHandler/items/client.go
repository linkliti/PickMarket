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

var serviceConfig = `{
	"loadBalancingPolicy": "round_robin",
	"healthCheckConfig": {
		"serviceName": "items"
	}
}`

func NewItemsClient() *ItemsClient {
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(serviceConfig),
	}

	addr := pmutils.GetEnv("ITEM_WORKER_ADDR", "localhost:1111")
	conn, err := grpc.Dial(addr, options...)
	if err != nil {
		panic(err)
	}
	client := parser.NewItemParserClient(conn)
	cc := ItemsClient{cl: client}
	return &cc
}
