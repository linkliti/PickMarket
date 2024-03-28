package main

import (
	"itemsWorker/db"
	"itemsWorker/service"
	"log/slog"
	"net"
	"pmutils"
	"protos/parser"

	"google.golang.org/grpc"
)

func main() {
	pmutils.SetupLogging("itemsWorker")
	// Connections
	parsClient, err := service.ConnectToParser()
	if err != nil {
		slog.Error("failed to connect to parser", err)
		return
	}
	database, err := db.NewDBConnection()
	if err != nil {
		slog.Error("failed to connect to database", err)
		return
	}

	// Start server
	grpcServer := grpc.NewServer()
	itemsService := service.NewItemsService(parsClient, database)
	parser.RegisterItemParserServer(grpcServer, itemsService)
	addr := pmutils.GetEnv("ITEMS_WORKER_ADDR", ":1111")
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	slog.Info("Starting items worker", "addr", addr)

	grpcServer.Serve(l)
}
