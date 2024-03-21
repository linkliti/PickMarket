package categories

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"protos/parser"

	"log/slog"
)

func (c *CategoryClient) GetRootCategories(rw http.ResponseWriter, r *http.Request) {
	// Get the variables from the request
	market, err := getMarketFromVars(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a request for root categories
	req := &parser.RootCategoriesRequest{
		Market: market,
	}

	// Use the client to send the request
	stream, err := c.cl.GetRootCategories(context.Background(), req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare a slice to hold the categories
	var categories []*parser.Category

	// Receive the categories from the stream
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		// Check if the response contains a Category
		if category := response.GetCategory(); category != nil {
			// Add the category to the slice
			categories = append(categories, category)
		} else if status := response.GetStatus(); status != nil {
			// Handle the error status
			slog.Warn("Received an error status: %v\n", status)
		}
	}

	// Convert the categories to JSON
	jsonData, err := json.Marshal(categories)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON data to the response
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonData)
}
