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

func ConnectToParser() (parser.ItemParserClient, error) {
	parserAddr := pmutils.GetEnv("PARSER_ADDR", "localhost:1111")
	conn, err := grpc.Dial(parserAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := parser.NewItemParserClient(conn)
	return client, nil
}
