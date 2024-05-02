package categories

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"pickmarket/requestHandler/misc"
	"protos/parser"

	"google.golang.org/protobuf/encoding/protojson"
)

func (c *CategoryClient) GetRootCategories(rw http.ResponseWriter, r *http.Request) {
	// Request
	market, err := misc.GetMarketFromUrlVar(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	req := &parser.RootCategoriesRequest{
		Market: market,
	}
	// gRPC call
	slog.Debug("GetRootCategories", "request", req)
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
