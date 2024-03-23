package service

import (
	"categoriesWorker/db"
	"pmutils"
	"protos/parser"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CategoryService struct {
	parsClient                               parser.CategoryParserClient
	db                                       *db.Database
	parser.UnimplementedCategoryParserServer // parser.UnsafeCategoryParserServer to require all methods implementation
}

func NewCategoryService(parsClient parser.CategoryParserClient, db *db.Database) *CategoryService {
	return &CategoryService{parsClient: parsClient, db: db}
}

func ConnectToParser() (parser.CategoryParserClient, error) {
	parserAddr := pmutils.GetEnv("PARSER_ADDR", "localhost:1111")
	conn, err := grpc.Dial(parserAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := parser.NewCategoryParserClient(conn)
	return client, nil
}
