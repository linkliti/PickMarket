package main

import (
	"itemsWorker/db"
	"itemsWorker/service"
	"log/slog"
	"net"
	"pmutils"
	"protos/parser"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	pmutils.SetupLogging("itemsWorker")
	// Connections
	parsClient, err := service.ConnectToParser()
	if err != nil {
		slog.Error("failed to connect to parser", "err", err)
		return
	}
	database, err := db.NewDBConnection()
	if err != nil {
		slog.Error("failed to connect to database", "err", err)
		return
	}
	defer database.Conn.Close()

	// Create server
	grpcServer := grpc.NewServer()
	itemsService := service.NewItemsService(parsClient, database)
	healthService := health.NewServer()
	parser.RegisterItemParserServer(grpcServer, itemsService)
	// Health
	healthpb.RegisterHealthServer(grpcServer, healthService)
	healthService.SetServingStatus("items", healthpb.HealthCheckResponse_SERVING)
	// Start server
	addr := pmutils.GetEnv("ITEM_WORKER_ADDR", ":1111")
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	slog.Info("Starting items worker", "addr", addr)

	grpcServer.Serve(l)
}
