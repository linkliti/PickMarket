package calc

import (
	"log/slog"
	"math"
	"protos/parser"
)

type calcVault struct {
	prefPointer  *parser.UserPref
	charPointers []*parser.Characteristic
}

type calc struct {
	itemList []*parser.ItemExtended
	userPref []*parser.UserPref
	vaults   map[string]*calcVault
}

var caseFuncs = map[string]func(*parser.Characteristic, *parser.UserPref){
	"color": calc_color,
}

func CalcWeight(itemList []*parser.ItemExtended, userPref []*parser.UserPref) error {
	c := &calc{
		itemList: itemList,
		userPref: userPref,
		vaults:   make(map[string]*calcVault),
	}
	slog.Debug("Filling vaults")
	c.fillPreferences()
	c.fillChars()
	slog.Debug("Calculating weights")
	c.calculateDifferences()
	c.calculateWeights()
	c.sumWeights()
	slog.Debug("Calculation successful")
	return nil
}

func (c *calc) calculateDifferences() {
	for key, v := range c.vaults {
		// Special cases
		if fn, ok := caseFuncs[key]; ok {
			for _, char := range v.charPointers {
				fn(char, v.prefPointer)
			}
		}
		// General cases
		switch v.prefPointer.Value.(type) {
		case *parser.UserPref_NumValue:
			for _, char := range v.charPointers {
				calc_numVal(char, v.prefPointer)
			}
		case *parser.UserPref_ListValue:
			for _, char := range v.charPointers {
				calc_listVal(char, v.prefPointer)
			}
		}
	}
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
	for _, item := range c.itemList {
		for _, char := range item.Chars {
			// Sum weights that are present in the vault
			if _, ok := c.vaults[char.Key]; ok {
				item.TotalWeight += char.CharWeight
			}
		}
	}
}
