package calc

import (
	"pmutils"
	"protos/parser"
)

func calc_numVal(char *parser.Characteristic, pref *parser.UserPref) {
	charVal := char.GetNumVal()
	prefVal := pref.GetNumValue()
	char.CharWeight = prefVal - charVal
}

func calc_listVal(char *parser.Characteristic, pref *parser.UserPref) {
	charList := char.GetListVal().Values
	prefList := pref.GetListValue().Values
	char.CharWeight = float64(len(pmutils.InterSection(charList, prefList)))
}

func calc_color(char *parser.Characteristic, pref *parser.UserPref) {
	calc_listVal(char, pref)
}
