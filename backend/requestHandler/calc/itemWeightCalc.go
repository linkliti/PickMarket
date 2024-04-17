package calc

import (
	"pmutils"
	"protos/parser"
)

type ItemWeightCalc struct {
	itemList  []*parser.ItemExtended
	userPrefs []*parser.UserPrefs
	prefKeys  []string
	caseFuncs map[string]func(*parser.Characteristic, *parser.UserPrefs)
}

type MinAndMax struct {
	min float64
	max float64
}

func NewItemWeightCalc(itemList []*parser.ItemExtended, userPrefs []*parser.UserPrefs) *ItemWeightCalc {
	o := &ItemWeightCalc{
		itemList:  itemList,
		userPrefs: userPrefs,
		prefKeys:  make([]string, 0, len(userPrefs)),
		caseFuncs: map[string]func(*parser.Characteristic, *parser.UserPrefs){
			"color": calc_color,
		},
	}
	for _, pref := range o.userPrefs {
		o.prefKeys = append(o.prefKeys, pref.Key)
	}
	for _, i := range o.itemList {
		o.calculateDifferences(i.GetChars())
	}
	minAndMax := make(map[string]MinAndMax, len(o.prefKeys))
	for _, k := range o.prefKeys {
		itemWeightList := make([]float64, len(o.prefKeys))
		for i, item := range o.itemList {
			itemWeightList[i] = item.Chars[k].ItemWeight
		}
		minAndMax[o.prefKeys[k]] = o.calcMinAndMax(&itemWeightList)
	}

	return o
}

func (c *ItemWeightCalc) calcMinAndMax(*[]float64) MinAndMax {

}

func (c *ItemWeightCalc) calculateDifferences(char []*parser.Characteristic) {
	for _, k := range c.prefKeys {
		if fn, ok := c.caseFuncs[c.prefKeys[k]]; ok {
			fn(char[k], c.userPrefs[k])
		} else if c.userPrefs[k].Type == parser.PrefType_LIST {
			calc_listVal(char[k], c.userPrefs[k])
		} else if c.userPrefs[k].Type == parser.PrefType_NUM {
			calc_numVal(char[k], c.userPrefs[k])
		} else {
			char[k].ItemWeight = 0
		}
	}
}

func calc_numVal(char *parser.Characteristic, pref *parser.UserPrefs) {
	charVal := char.GetNumVal()
	prefVal := pref.GetNumValue()
	char.ItemWeight = prefVal - charVal
}

func calc_listVal(char *parser.Characteristic, pref *parser.UserPrefs) {
	charList := char.GetListVal().Values
	prefList := pref.GetListValue().Values
	char.ItemWeight = float64(len(pmutils.InterSection(charList, prefList)))
}
