package service

import (
	"itemsWorker/db"
	"pmutils"
	"protos/parser"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ItemsService struct {
	parsClient                           parser.ItemParserClient
	db                                   *db.Database
	parser.UnimplementedItemParserServer // parser.UnsafeItemsParserServer to require all methods implementation
}

func NewItemsService(parsClient parser.ItemParserClient, db *db.Database) *ItemsService {
	return &ItemsService{parsClient: parsClient, db: db}
}

var serviceConfig = `{
	"loadBalancingPolicy": "round_robin",
	"healthCheckConfig": {
		"serviceName": "parser"
	}
}`

func ConnectToParser() (parser.ItemParserClient, error) {
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(serviceConfig),
	}

	parserAddr := pmutils.GetEnv("PARSER_ADDR", "localhost:1111")
	conn, err := grpc.Dial(parserAddr, options...)
	if err != nil {
		return nil, err
	}
	client := parser.NewItemParserClient(conn)
	return client, nil
}
