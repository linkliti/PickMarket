package misc

import (
	"fmt"
	"net/http"
	"protos/parser"
	"strings"

	"github.com/gorilla/mux"
)

func GetMarketFromVars(r *http.Request) (parser.Markets, error) {
	marketStr := mux.Vars(r)["market"]
	if marketStr == "" {
		return 0, fmt.Errorf("market not provided")
	}

	switch strings.ToUpper(marketStr) {
	case "OZON":
		return parser.Markets_OZON, nil
	// Add more cases as you support more markets
	default:
		return 0, fmt.Errorf("unknown market: %s", marketStr)
	}
}

// func GetMarketFromVars(r *http.Request) (parser.Markets, error) {
// 	marketStr := r.URL.Query().Get("market")
// 	if marketStr == "" {
// 		return 0, fmt.Errorf("market not provided")
// 	}

// 	switch strings.ToUpper(marketStr) {
// 	case "OZON":
// 		return parser.Markets_OZON, nil
// 	// Add more cases as you support more markets
// 	default:
// 		return 0, fmt.Errorf("unknown market: %s", marketStr)
// 	}
// }
