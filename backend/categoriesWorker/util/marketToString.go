package util

import (
	"fmt"
	"protos/parser"
)

// MarketToString converts a parser.Markets value to a string.
func MarketToString(market parser.Markets) string {
	return market.String()
}

// StringToMarket converts a string to a parser.Markets value.
func StringToMarket(market string) (parser.Markets, error) {
	if val, ok := parser.Markets_value[market]; ok {
		return parser.Markets(val), nil
	}
	return -1, fmt.Errorf("invalid market: %s", market)
}
