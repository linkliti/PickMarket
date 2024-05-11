package itemMerge

import (
	"log/slog"
	"protos/parser"
	"regexp"
	"strings"
)

func OzonMergeItems(items []*parser.ItemExtended) []*parser.ItemExtended {
	uniqueItems := make(map[string]*parser.ItemExtended)
	re := regexp.MustCompile(`\d+`)

	for _, item := range items {
		urlParts := strings.Split(item.Item.Url, "/")
		productDetails := strings.Split(urlParts[len(urlParts)-2], "-")

		// Exclude the article at the end and max 4 words before it
		lastIndex := len(productDetails) - 2
		endIndex := lastIndex
		for i := endIndex; lastIndex-i < 5 && i > 0; i-- {
			endIndex -= 1
			if re.MatchString(productDetails[i]) {
				break
			}
		}
		endIndex += 2
		newUrl := strings.Join(productDetails[:endIndex], "-")

		if existingItem, exists := uniqueItems[newUrl]; exists {
			slog.Debug("Merging item", "itemUrl", item.Item.Url, "existingUrl", existingItem.Item.Url)
			MergeItemData(existingItem, item)
		} else {
			uniqueItems[newUrl] = item
		}
	}
	// Convert the map to a slice
	result := make([]*parser.ItemExtended, 0, len(uniqueItems))
	for _, item := range uniqueItems {
		result = append(result, item)
	}

	return result
}

func MergeItemData(existingItem *parser.ItemExtended, item *parser.ItemExtended) {
	// Add similar items
	for _, similarItem := range existingItem.Similar {
		if similarItem.Url == item.Item.Url {
			return
		}
	}
	existingItem.Similar = append(existingItem.Similar, item.Item)
	// Chars
	for _, char := range item.Chars {
		switch v := char.Value.(type) {
		case *parser.Characteristic_NumVal:
			for _, existingChar := range existingItem.Chars {
				if existingChar.Key == char.Key {
					if existingNumVal, ok := existingChar.Value.(*parser.Characteristic_NumVal); ok {
						if v.NumVal > existingNumVal.NumVal {
							slog.Debug("Updated value", "value", v.NumVal, "oldValue", existingNumVal.NumVal)
							existingNumVal.NumVal = v.NumVal
						}
						break
					}
				}
			}
		case *parser.Characteristic_ListVal:
			for _, existingChar := range existingItem.Chars {
				if existingChar.Key == char.Key {
					if existingListVal, ok := existingChar.Value.(*parser.Characteristic_ListVal); ok {
						existingValues := make(map[string]bool)
						for _, value := range existingListVal.ListVal.Values {
							existingValues[value] = true
						}
						for _, value := range v.ListVal.Values {
							if !existingValues[value] {
								slog.Debug("Added value", "value", value, "oldListVal", existingListVal.ListVal.Values)
								existingListVal.ListVal.Values = append(existingListVal.ListVal.Values, value)
							}
						}
						break
					}
				}
			}
		}
	}
}
