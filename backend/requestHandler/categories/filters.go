package categories

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"protos/parser"
)

func (c *CategoryClient) GetCategoryFilters(rw http.ResponseWriter, r *http.Request) {
	// Market
	market, err := getMarketFromVars(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// Category URL
	categoryUrl := r.URL.Query().Get("categoryUrl")
	if categoryUrl == "" {
		http.Error(rw, "categoryUrl not provided", http.StatusInternalServerError)
		return
	}
	req := &parser.FiltersRequest{
		Market:      market,
		CategoryUrl: categoryUrl,
	}
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
