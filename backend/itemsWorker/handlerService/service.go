package handlerservice

import (
	"context"
	"io"
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

func (c *ItemsService) GetItems(req *parser.ItemsRequest, srv parser.ItemParser_GetItemsServer) error {
	client := c.connectToParser()
	stream, err := client.GetItems(context.Background(), req)
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

func (c *ItemsService) GetItemCharacteristics(req *parser.CharacteristicsRequest, srv parser.ItemParser_GetItemCharacteristicsServer) error {
	client := c.connectToParser()
	stream, err := client.GetItemCharacteristics(context.Background(), req)
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
