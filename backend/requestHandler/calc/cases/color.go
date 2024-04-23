package cases

import (
	"log/slog"
	"pmutils"
	"protos/parser"
	"strings"
)

func CalcList_Color(char *parser.Characteristic, pref *parser.UserPref) {
	listVal := char.GetListVal()
	for _, color := range listVal.Values {
		extraColors := MatchColors(color)
		listVal.Values = append(listVal.Values, extraColors...)
	}
	CalcList_atleastOne(char, pref)
}

func MatchColors(rawColor string) (matchedColors []string) {
	rawColor = strings.ToLower(rawColor)
	if len(rawColor) > 0 {
		rawColorCap := pmutils.CapitalFirstLet(rawColor)
		if _, ok := colors[rawColorCap]; ok {
			slog.Debug("Color is equal", "rawColor", rawColor)
			return
		}
	}
	for fullColor, colorVariants := range colors {
		for _, colorVariant := range colorVariants {
			if strings.Contains(rawColor, colorVariant) {
				matchedColors = append(matchedColors, fullColor)
				break
			}
		}
	}
	slog.Debug("Found color matches", "matchedColors", matchedColors, "rawColor", rawColor)
	return
}

var colors = map[string][2]string{
	"Бежевый":      {"бежев", "beige"},
	"Белый":        {"бел", "white"},
	"Бордовый":     {"бордов", "maroon"},
	"Бронзовый":    {"бронзов", "bronze"},
	"Голубой":      {"голуб", "lightblue"},
	"Желтый":       {"желт", "yellow"},
	"Зеленый":      {"зелен", "green"},
	"Золотой":      {"золот", "gold"},
	"Коричневый":   {"коричнев", "brown"},
	"Красный":      {"красн", "red"},
	"Оранжевый":    {"оранжев", "orange"},
	"Прозрачный":   {"прозрачн", "transparent"},
	"Разноцветный": {"разноцветн", "multicolor"},
	"Розовый":      {"розов", "pink"},
	"Серебряный":   {"серебр", "silver"},
	"Серый":        {"сер", "gray"},
	"Синий":        {"син", "blue"},
	"Фиолетовый":   {"фиолетов", "purple"},
	"Хаки":         {"хак", "khaki"},
	"Черный":       {"черн", "black"},
}
