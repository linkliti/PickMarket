package calc

import (
	"fmt"
	"log/slog"
	"math"
	"pickmarket/requestHandler/calc/cases"
	"protos/parser"
)

// Keys must be lowered
var caseFuncs = map[string]caseStruct{
	"color": {fn: cases.CalcList_Color, calcType: LIST_TYPE},
}

func CalcWeight(itemList []*parser.ItemExtended, userPref map[string]*parser.UserPref, req *parser.ItemsRequest) error {
	if req == nil || itemList == nil || userPref == nil {
		return fmt.Errorf("calc received nil pointer")
	}
	c := &calc{
		itemList: itemList,
		userPref: userPref,
		vaults:   make(map[string]*calcVault),
		req:      req,
	}
	slog.Debug("Filling vaults", "pageUrl", c.req.PageUrl, "prefs", len(c.userPref), "items", len(c.itemList))
	c.fillPreferences()
	c.fillChars()
	c.removeVaultsWithoutChars()
	slog.Debug("Calculating weights", "pageUrl", c.req.PageUrl, "vaults", len(c.vaults))
	c.calculateDifferences()
	c.calculateWeights()
	c.sumWeights()
	slog.Debug("Calculation successful", "pageUrl", c.req.PageUrl)
	return nil
}

func (c *calc) calculateWeights() {
	for key, v := range c.vaults {
		// Min and max values
		min, max := math.MaxFloat64, -math.MaxFloat64
		for _, char := range v.charPointers {
			if char.CharWeight < min {
				min = char.CharWeight
			}
			if char.CharWeight > max {
				max = char.CharWeight
			}
		}
		// Multiplying same values by priority
		if min == max {
			slog.Debug("calcWeight: Removing vault with equal max and min", "key", key)
			for _, char := range v.charPointers {
				char.MaxWeight = float64(v.prefPointer.Priority)
				char.CharWeight = max * float64(v.prefPointer.Priority)
			}
			return
		}
		// Normalizing and multiplying by priority
		for _, char := range v.charPointers {
			char.MaxWeight = float64(v.prefPointer.Priority)
			char.CharWeight = (char.CharWeight - min) / (max - min) * float64(v.prefPointer.Priority)
			slog.Debug("calcWeight: Calculated weight", "key", key, "char", char.Key, "weight", char.CharWeight)
		}
	}
}
