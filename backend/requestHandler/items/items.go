package items

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"pickmarket/requestHandler/calc"
	"pickmarket/requestHandler/misc"
	"pmutils"
	"protos/parser"
	"sync"

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
	itemR := reqBodyProto.GetRequest()
	req := &parser.ItemsRequest{
		Market:     market,
		PageUrl:    itemR.PageUrl,
		UserQuery:  itemR.UserQuery,
		Params:     itemR.Params,
		NumOfPages: itemR.NumOfPages,
	}
	slog.Debug("GetItems", "request", req)

	// Extended items
	itemsList, err := c.grpcGetItems(req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusServiceUnavailable)
		return
	}
	userPrefs := reqBodyProto.GetUserPrefs()
	c.appendExtraChars(itemsList)
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

func (c *ItemsClient) appendExtraChars(itemsList []*parser.ItemExtended) {
	// Add items info as chars
	var wg sync.WaitGroup
	for _, item := range itemsList {
		wg.Add(1)
		go func(item *parser.ItemExtended) {
			defer wg.Done()
			chars := make([]*parser.Characteristic, 0, 3)
			if item.Item.Comments != nil {
				commentsChar := parser.Characteristic{
					Name:  "Отзывы",
					Key:   "pm_reviews",
					Value: &parser.Characteristic_NumVal{NumVal: float64(*item.Item.Comments)},
				}
				chars = append(chars, &commentsChar)
			}
			if item.Item.Rating != nil {
				ratingChar := parser.Characteristic{
					Name:  "Рейтинг",
					Key:   "pm_rating",
					Value: &parser.Characteristic_NumVal{NumVal: *item.Item.Rating},
				}
				chars = append(chars, &ratingChar)
			}
			if item.Item.OldPrice != nil {
				oldPriceChar := parser.Characteristic{
					Name:  "Старая цена",
					Key:   "pm_oldPrice",
					Value: &parser.Characteristic_NumVal{NumVal: *item.Item.OldPrice},
				}
				chars = append(chars, &oldPriceChar)
			}
			originalChar := parser.Characteristic{
				Name:  "Оригинал",
				Key:   "pm_isOriginal",
				Value: &parser.Characteristic_ListVal{ListVal: pmutils.BoolToStringList(item.Item.GetOriginal())},
			}
			chars = append(chars, &originalChar)

			adultChar := parser.Characteristic{
				Name:  "Для взрослых",
				Key:   "pm_isAdult",
				Value: &parser.Characteristic_ListVal{ListVal: pmutils.BoolToStringList(item.Item.GetIsAdult())},
			}
			chars = append(chars, &adultChar)

			priceChar := parser.Characteristic{
				Name:  "Цена",
				Key:   "pm_price",
				Value: &parser.Characteristic_NumVal{NumVal: item.Item.Price},
			}
			chars = append(chars, &priceChar)
			item.Chars = append(item.Chars, chars...)
		}(item)
	}
	wg.Wait()
}
