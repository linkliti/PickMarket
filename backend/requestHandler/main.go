package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"pickmarket/requestHandler/categories"
	"pickmarket/requestHandler/items"
	"pickmarket/requestHandler/misc"
	"pmutils"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	pmutils.SetupLogging("requestHandler")
	bindAddress := pmutils.GetEnv("HANDLER_ADDR", "localhost:1111")
	sm := mux.NewRouter()

	categClient := categories.NewCategoryClient()
	itemClient := items.NewItemsClient()

	mainR := sm.Methods(http.MethodGet).Subrouter()
	mainR.HandleFunc("/health", misc.HealthCheck)

	catR := sm.Methods(http.MethodGet).PathPrefix("/categories").Subrouter()
	catR.HandleFunc("/{market}/root", categClient.GetRootCategories)
	catR.HandleFunc("/{market}/sub", categClient.GetSubCategories)
	catR.HandleFunc("/{market}/filter", itemClient.GetCategoryFilters)

	// itR := sm.Methods(http.MethodGet).PathPrefix("/items").Subrouter()
	// itR.HandleFunc("/{market}/chars", itemClient.GetItemCharacteristics)

	itCalc := sm.Methods(http.MethodPost).PathPrefix("/calc").Subrouter()
	itCalc.HandleFunc("/{market}/list", itemClient.PostItems)

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	// create a new server
	s := http.Server{
		Addr:         bindAddress,
		Handler:      ch(sm),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	go func() {
		slog.Info("Starting server", "bindAddress", bindAddress)

		err := s.ListenAndServe()
		if err != nil {
			slog.Error("Error starting server", "err", err)
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
