package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"pickmarket/requestHandler/categories"
	"pmutils"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	pmutils.SetupLogging()
	bindAddress := pmutils.GetEnv("HANDLER_ADDR", "localhost:1111")
	sm := mux.NewRouter()

	categClient := categories.NewCategoryClient()

	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/categories/root", categClient.GetRootCategories)
	getR.HandleFunc("/categories/sub", categClient.GetSubCategories)
	getR.HandleFunc("/categories/filter", categClient.GetFilterCategories)
	getR.HandleFunc("/health", categClient.HealthCheck)

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	// create a new server
	s := http.Server{
		Addr:         bindAddress,
		Handler:      ch(sm),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	go func() {
		slog.Info("Starting server on port " + bindAddress)

		err := s.ListenAndServe()
		if err != nil {
			slog.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
	cancel()
}
