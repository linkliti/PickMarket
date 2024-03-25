package main

import (
	"categoriesWorker/db"
	"categoriesWorker/manager"
	"categoriesWorker/service"
	"log/slog"
	"net"
	"pmutils"
	"protos/parser"

	"google.golang.org/grpc"
)

func main() {
	pmutils.SetupLogging()
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
	slog.Info("Verifying categories...")
	// Verify categories
	manager := manager.NewManager(parsClient, database)
	if err := manager.UpdateRootCategories(); err != nil {
		slog.Error("failed to update root categories", err)
		return
	}
	slog.Info("all root categories updated")
	if err := manager.UpdateAllSubCategories(); err != nil {
		slog.Error("failed to update sub categories", err)
		return
	}
	slog.Info("all subcategories updated")

	// Start server
	grpcServer := grpc.NewServer()
	categoryService := service.NewCategoryService(parsClient, database)
	parser.RegisterCategoryParserServer(grpcServer, categoryService)
	addr := pmutils.GetEnv("CATEGORIES_WORKER_ADDR", "localhost:1111")
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	slog.Info("Starting categories worker", "addr", addr)

	grpcServer.Serve(l)
}
