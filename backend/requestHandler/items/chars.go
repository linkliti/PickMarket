package items

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"pickmarket/requestHandler/misc"
	"protos/parser"
)

func (c *ItemsClient) GetItemCharacteristics(rw http.ResponseWriter, r *http.Request) {
	// Request
	market, err := misc.GetMarketFromUrlVar(r)
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
	slog.Debug("GetItemCharacteristics", "request", req)
	charsList, err := c.grpcGetCharacteristics(req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// All to JSON
	jsonData, err := json.Marshal(charsList)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// Return
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonData)
}
