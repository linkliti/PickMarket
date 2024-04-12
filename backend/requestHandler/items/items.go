package items

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"pickmarket/requestHandler/misc"
	"protos/parser"
	// "pickmarket/requestHandler/calc"
)

func (c *ItemsClient) GetItems(rw http.ResponseWriter, r *http.Request) {
	// Items request
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
	itemsList, err := c.grpcGetItems(req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// All to JSON
	jsonData, err := json.Marshal(itemsList)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// Return
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonData)
}

// func (c *ItemsClient) getCharsForItems(market parser.Markets) {
// }
