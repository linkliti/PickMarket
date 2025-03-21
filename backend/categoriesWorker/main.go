package main

import (
	"categoriesWorker/db"
	"categoriesWorker/manager"
	"categoriesWorker/service"
	"log/slog"
	"net"
	"pmutils"
	"protos/parser"
	"runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	pmutils.SetupLogging("categoriesWorker")
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

	slog.Info("Verifying categories...")
	// Verify categories
	var workpoolSize int = runtime.NumCPU() / 3
	manager := manager.NewManager(parsClient, database, workpoolSize)
	if err := manager.UpdateRootCategories(); err != nil {
		slog.Error("failed to update root categories", "err", err)
		return
	}
	slog.Info("All root categories updated")
	if err := manager.UpdateAllSubCategories(); err != nil {
		slog.Error("failed to update sub categories", "err", err)
		return
	}
	slog.Info("All subcategories updated")

	// Create server
	grpcServer := grpc.NewServer()
	categoryService := service.NewCategoryService(parsClient, database)
	healthService := health.NewServer()
	parser.RegisterCategoryParserServer(grpcServer, categoryService)
	// Health
	healthpb.RegisterHealthServer(grpcServer, healthService)
	healthService.SetServingStatus("categories", healthpb.HealthCheckResponse_SERVING)
	// Start server
	addr := pmutils.GetEnv("CATEGORIES_WORKER_ADDR", "localhost:1111")
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	slog.Info("Starting categories worker", "addr", addr)

	grpcServer.Serve(l)
}
