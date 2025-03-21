package categories

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"pickmarket/requestHandler/misc"
	"protos/parser"

	"log/slog"

	"google.golang.org/protobuf/encoding/protojson"
)

func (c *CategoryClient) GetSubCategories(rw http.ResponseWriter, r *http.Request) {
	// Market
	market, err := misc.GetMarketFromUrlVar(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// Parent category URL
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(rw, "url is required", http.StatusInternalServerError)
		return
	}
	req := &parser.SubCategoriesRequest{
		Market:      market,
		CategoryUrl: url,
	}
	slog.Debug("GetSubCategories", "req", req)
	// gRPC call
	stream, err := c.cl.GetSubCategories(context.Background(), req)
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
	// parser.Category to json.RawMessage
	categsRaw := make([]json.RawMessage, len(categories))
	for i, item := range categories {
		categJSON, err := protojson.Marshal(item)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		categsRaw[i] = json.RawMessage(categJSON)
	}
	// json.RawMessage to JSONB
	categsJSON, err := json.Marshal(categsRaw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// Return
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(categsJSON)
}
