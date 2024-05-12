package items

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"pickmarket/requestHandler/calc"
	"pickmarket/requestHandler/itemMerge"
	"pickmarket/requestHandler/misc"
	"protos/parser"

	"google.golang.org/protobuf/encoding/protojson"
)

func (c *ItemsClient) PostItems(rw http.ResponseWriter, r *http.Request) {
	// Items request
	market, err := misc.GetMarketFromUrlVar(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	var reqBody json.RawMessage
	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	reqBodyProto := &parser.ItemsRequestWithPrefs{}
	err = protojson.Unmarshal(reqBody, reqBodyProto)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	errString, ok := ValidateItemsRequestWithPrefs(reqBodyProto)
	if !ok {
		http.Error(rw, errString, http.StatusBadRequest)
		return
	}
	itemR := reqBodyProto.GetRequest()

	req := &parser.ItemsRequest{
		Market:    market,
		PageUrl:   itemR.PageUrl,
		UserQuery: itemR.UserQuery,
		// Params:     itemR.Params,
		NumOfPages: itemR.NumOfPages,
	}
	slog.Debug("GetItems", "request", req)

	// Extended items
	itemsList, err := c.grpcGetItems(req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusServiceUnavailable)
		return
	}
	userPrefs := reqBodyProto.GetPrefs()
	// Append extra chars
	c.appendExtraChars(itemsList)
	// Merge duplicates
	switch req.Market {
	case parser.Markets_OZON:
		itemsList = itemMerge.OzonMergeItems(itemsList)
	}
	// Calc Weight
	err = calc.CalcWeight(itemsList, userPrefs, req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// All to JSON
	// parser.ItemExtended to json.RawMessage
	itemsRaw := make([]json.RawMessage, len(itemsList))
	for i, item := range itemsList {
		itemJSON, err := protojson.Marshal(item)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		itemsRaw[i] = json.RawMessage(itemJSON)
	}
	// json.RawMessage to JSONB
	itemsJSON, err := json.Marshal(itemsRaw)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(itemsJSON)
}

func ValidateItemsRequestWithPrefs(req *parser.ItemsRequestWithPrefs) (string, bool) {
	if req == nil {
		return "Request is nil", false
	}

	if req.Request.PageUrl == "" {
		return "PageUrl is empty", false
	}

	if *req.Request.NumOfPages <= 0 || *req.Request.NumOfPages > 5 {
		return "NumOfPages is not within the range of 1 to 5", false
	}

	if len(req.Prefs) == 0 {
		return "Prefs is empty", false
	}

	return "", true
}
