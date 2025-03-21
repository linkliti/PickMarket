package items

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"pickmarket/requestHandler/misc"
	"protos/parser"

	"google.golang.org/protobuf/encoding/protojson"
)

func (c *ItemsClient) GetCategoryFilters(rw http.ResponseWriter, r *http.Request) {
	// Request
	market, err := misc.GetMarketFromUrlVar(r)
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
	filterList, err := c.grpcGetFilters(req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// All to JSON
	// parser.Filter to json.RawMessage
	filterRaw := make([]json.RawMessage, len(filterList))
	for i, char := range filterList {
		charJSON, err := protojson.Marshal(char)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		filterRaw[i] = json.RawMessage(charJSON)
	}
	// json.RawMessage to JSONB
	filtersJSON, err := json.Marshal(filterRaw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(filtersJSON)
}
