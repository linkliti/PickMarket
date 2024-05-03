package calc

import (
	"fmt"
	"log/slog"
	"math"
	"pickmarket/requestHandler/calc/cases"
	"pmutils"
	"protos/parser"
	"sync"
)

// Keys must be lowered
var caseFuncs = map[string]caseStruct{
	"color":       {fn: cases.CalcList_Color, calcType: LIST_TYPE},
	"pm_price":    {fn: cases.CalcNum_difference, calcType: NUM_TYPE},
	"pm_oldprice": {fn: cases.CalcNum_difference, calcType: NUM_TYPE},
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

func (c *calc) calculateDifferences() {
	// var wg sync.WaitGroup
	for key, v := range c.vaults {
		// wg.Add(1)
		func(key string, v *calcVault) {
			// defer wg.Done()
			// Special cases
			if fnStruct, ok := caseFuncs[key]; ok {
				if v.prefType != fnStruct.calcType {
					slog.Error("calcDif: Special case key has incorrect type", "key", key, "expected", fnStruct.calcType.Name(), "got", v.prefType.Name())
					return
				}
				funcName := pmutils.GetFunctionName(fnStruct.fn)
				slog.Debug("calcDif: Calculating key using special func", "key", key, "func", funcName, "isList", fnStruct.calcType.Name())
				for _, char := range v.charPointers {
					fnStruct.fn(char, v.prefPointer)
				}
				return
			}
			// General cases
			switch v.prefType {
			case NUM_TYPE:
				slog.Debug("calcDif: Calculating key using numVal func", "key", key)
				for _, char := range v.charPointers {
					cases.CalcNum_negAbsDifference(char, v.prefPointer)
				}
			case LIST_TYPE:
				slog.Debug("calcDif: Calculating key using listVal func", "key", key)
				for _, char := range v.charPointers {
					cases.CalcList_numOfMatches(char, v.prefPointer)
				}
			}
		}(key, v)
	}
	// wg.Wait()
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
		// Remove vault from calculations if all values are equal
		if min == max {
			slog.Debug("calcWeight: Removing vault with equal max and min", "key", key)
			delete(c.vaults, key)
			continue
		}
		// Normalizing and multiplying by priority
		for _, char := range v.charPointers {
			char.CharWeight = (char.CharWeight - min) / (max - min) * float64(v.prefPointer.Priority)
		}
	}
}

func (c *calc) sumWeights() {
	var wg sync.WaitGroup
	for _, item := range c.itemList {
		wg.Add(1)
		go func(item *parser.ItemExtended) {
			wg.Done()
			for _, char := range item.Chars {
				// Sum weights that are present in the vault
				if _, ok := c.vaults[char.Key]; ok {
					item.TotalWeight += char.CharWeight
				}
			}
		}(item)
	}
	wg.Wait()
}
