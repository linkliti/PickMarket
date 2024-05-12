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

var serviceConfig = `{
	"loadBalancingPolicy": "round_robin",
	"healthCheckConfig": {
		"serviceName": "categories"
	}
}`

func NewCategoryClient() *CategoryClient {
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(serviceConfig),
	}

	addr := pmutils.GetEnv("CATEGORIES_WORKER_ADDR", "localhost:1111")
	conn, err := grpc.Dial(addr, options...)
	if err != nil {
		panic(err)
	}
	client := parser.NewCategoryParserClient(conn)
	cc := CategoryClient{cl: client}
	return &cc
}
