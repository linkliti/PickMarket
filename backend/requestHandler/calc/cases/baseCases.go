package cases

import (
	"math"
	"pmutils"
	"protos/parser"
)

func getNumVals(char *parser.Characteristic, pref *parser.UserPref) (charVal float64, prefVal float64) {
	charVal = char.GetNumVal()
	prefVal = pref.GetNumVal()
	return
}

func getListVals(char *parser.Characteristic, pref *parser.UserPref) (charList []string, prefList []string) {
	charList = char.GetListVal().Values
	prefList = pref.GetListVal().Values
	return
}

// Calculators

func CalcNum_negAbsDifference(char *parser.Characteristic, pref *parser.UserPref) {
	// Weight by the always negative difference
	charVal, prefVal := getNumVals(char, pref)
	char.CharWeight = -math.Abs(prefVal - charVal)
}

func CalcList_numOfMatches(char *parser.Characteristic, pref *parser.UserPref) {
	// Weight by the length of intersection
	charList, prefList := getListVals(char, pref)
	char.CharWeight = float64(len(pmutils.InterSection(charList, prefList)))
}

func CalcNum_difference(char *parser.Characteristic, pref *parser.UserPref) {
	// Weight by the difference
	charVal, prefVal := getNumVals(char, pref)
	char.CharWeight = prefVal - charVal
}

func CalcList_atleastOne(char *parser.Characteristic, pref *parser.UserPref) {
	// Weight by having at least one pref in char
	charList, prefList := getListVals(char, pref)
	if len(pmutils.InterSection(charList, prefList)) > 0 {
		char.CharWeight = 1
	}
}
