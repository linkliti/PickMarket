package items

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"pickmarket/requestHandler/misc"
	"protos/parser"
)

func (c *ItemsClient) GetCategoryFilters(rw http.ResponseWriter, r *http.Request) {
	// Request
	market, err := misc.GetMarketFromVars(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(rw, "url is required", http.StatusBadRequest)
		return
	}
	req := &parser.FiltersRequest{
		Market:      market,
		CategoryUrl: url,
	}
	slog.Debug("GetCategoryFilters", "request", req)
	// gRPC call
	stream, err := c.cl.GetCategoryFilters(context.Background(), req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	var filters []*parser.Filter
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
		if filter := response.GetFilter(); filter != nil {
			filters = append(filters, filter)
		} else if status := response.GetStatus(); status != nil {
			fmt.Printf("Received an error status: %v\n", status)
			http.Error(rw, status.Message, http.StatusInternalServerError)
		}
	}
	// All to JSON
	jsonData, err := json.Marshal(filters)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// Return
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonData)
}
