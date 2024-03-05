package main

import (
	handlerservice "categoriesWorker/handlerService"
	"log/slog"
	"net"
	"pmutils"
	"protos/parser"

	"google.golang.org/grpc"
)

func main() {
	pmutils.SetupLogging()
	grpcServer := grpc.NewServer()
	categoryService := handlerservice.NewCategoryService()
	parser.RegisterCategoryParserServer(grpcServer, categoryService)
	addr := pmutils.GetEnv("CATEGORIES_WORKER_ADDR", ":1111")
	l, err := net.Listen("tcp", addr)
	slog.Info("Starting categories worker", "addr", addr)
	if err != nil {
		panic(err)
	}

	grpcServer.Serve(l)
}
