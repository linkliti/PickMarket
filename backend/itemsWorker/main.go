package main

import (
	handlerservice "itemsWorker/handlerService"
	"log/slog"
	"net"
	"pmutils"
	"protos/parser"

	"google.golang.org/grpc"
)

func main() {
	pmutils.SetupLogging()
	grpcServer := grpc.NewServer()
	itemsService := handlerservice.NewItemsService()
	parser.RegisterItemParserServer(grpcServer, itemsService)
	addr := pmutils.GetEnv("ITEMS_WORKER_ADDR", ":1111")
	l, err := net.Listen("tcp", addr)
	slog.Info("Starting items worker", "addr", addr)
	if err != nil {
		panic(err)
	}

	grpcServer.Serve(l)
}
