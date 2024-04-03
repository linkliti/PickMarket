package categories

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"pickmarket/requestHandler/misc"
	"protos/parser"
)

func (c *CategoryClient) GetRootCategories(rw http.ResponseWriter, r *http.Request) {
	// Request
	market, err := misc.GetMarketFromVars(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	req := &parser.RootCategoriesRequest{
		Market: market,
	}
	// gRPC call
	stream, err := c.cl.GetRootCategories(context.Background(), req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	var categories []*parser.Category
	for {
		response, err := stream.Recv()
		// End of stream
		if err == io.EOF {
			break
		}
		// Failed message
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		// Message
		if category := response.GetCategory(); category != nil {
			categories = append(categories, category)
		} else if status := response.GetStatus(); status != nil {
			slog.Warn("Received an error status", "status", status.Message)
			http.Error(rw, status.Message, http.StatusInternalServerError)
		}
	}
	// All to JSON
	jsonData, err := json.Marshal(categories)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// Return
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonData)
}
