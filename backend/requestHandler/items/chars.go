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

func (c *ItemsClient) GetItemCharacteristics(rw http.ResponseWriter, r *http.Request) {
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
	req := &parser.CharacteristicsRequest{
		Market:  market,
		ItemUrl: url,
	}
	// gRPC call
	stream, err := c.cl.GetItemCharacteristics(context.Background(), req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	var chars []*parser.Characteristic
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
		if char := response.GetCharacteristic(); char != nil {
			chars = append(chars, char)
		} else if status := response.GetStatus(); status != nil {
			slog.Warn("Received an error status", "status", status.Message)
			http.Error(rw, status.Message, http.StatusInternalServerError)
		}
	}
	// All to JSON
	jsonData, err := json.Marshal(chars)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// Return
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonData)
}
