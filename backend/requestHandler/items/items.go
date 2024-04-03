package items

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"pickmarket/requestHandler/misc"
	"protos/parser"
)

func (c *ItemsClient) GetItems(rw http.ResponseWriter, r *http.Request) {
	// Request
	market, err := misc.GetMarketFromVars(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(rw, "url is required", http.StatusBadRequest)
	}
	userQuery := r.URL.Query().Get("userQuery")
	params := r.URL.Query().Get("params")
	var numOfPages int32 = 1
	req := &parser.ItemsRequest{
		Market:     market,
		PageUrl:    url,
		UserQuery:  &userQuery,
		Params:     &params,
		NumOfPages: &numOfPages,
	}
	slog.Debug("GetItems", "request", req)
	// gRPC call
	stream, err := c.cl.GetItems(context.Background(), req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	var items []*parser.Item
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
		if item := response.GetItem(); item != nil {
			items = append(items, item)
		} else if status := response.GetStatus(); status != nil {
			slog.Warn("Received an error status", "status", status.Message)
			http.Error(rw, status.Message, http.StatusInternalServerError)
		}
	}
	// All to JSON
	jsonData, err := json.Marshal(items)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// Return
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonData)
}
