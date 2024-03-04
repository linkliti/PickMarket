package parserGetters

import (
	"context"
	"fmt"
	"io"
	"log"

	"protos/parser"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetRootCategories() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	cc := parser.NewCategoryParserClient(conn)
	req := &parser.RootCategoriesRequest{Market: parser.Markets_OZON}
	stream, err := cc.GetRootCategories(context.Background(), req)
	if err != nil {
		panic(err)
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break // end of stream
		}
		if err != nil {
			log.Fatalf("stream error: %v", err)
		}
		// check the oneof message field
		switch m := resp.Message.(type) {
		case *parser.CategoryResponse_Category:
			// print the category title and url
			fmt.Printf("Category: %s, URL: %s\n", m.Category.Title, m.Category.Url)
		case *parser.CategoryResponse_Status:
			// print the status code and message
			fmt.Printf("Status: %d, Message: %s\n", m.Status.Code, m.Status.Message)
		default:
			// unknown message type
			log.Fatalf("unknown message type: %T", m)
		}
	}
}
