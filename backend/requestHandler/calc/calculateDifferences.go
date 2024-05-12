package calc

import (
	"log/slog"
	"pickmarket/requestHandler/calc/cases"
	"pmutils"
)

func (c *calc) calculateDifferences() {
	for key, v := range c.vaults {
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

	}
}
