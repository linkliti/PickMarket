package items

import (
	"pmutils"
	"protos/parser"
	"sync"
)

func (c *ItemsClient) appendExtraChars(itemsList []*parser.ItemExtended) {
	// Add items info as chars
	var wg sync.WaitGroup
	for _, item := range itemsList {
		wg.Add(1)
		go func(item *parser.ItemExtended) {
			defer wg.Done()
			chars := make([]*parser.Characteristic, 0, 3)
			if item.Item.Comments != nil {
				commentsChar := parser.Characteristic{
					Name:  "Отзывы",
					Key:   "pm_reviews",
					Value: &parser.Characteristic_NumVal{NumVal: float64(*item.Item.Comments)},
				}
				chars = append(chars, &commentsChar)
			}
			if item.Item.Rating != nil {
				ratingChar := parser.Characteristic{
					Name:  "Рейтинг",
					Key:   "pm_rating",
					Value: &parser.Characteristic_NumVal{NumVal: *item.Item.Rating},
				}
				chars = append(chars, &ratingChar)
			}
			if item.Item.OldPrice != nil {
				oldPriceChar := parser.Characteristic{
					Name:  "Старая цена",
					Key:   "pm_oldprice",
					Value: &parser.Characteristic_NumVal{NumVal: *item.Item.OldPrice},
				}
				chars = append(chars, &oldPriceChar)
			}
			originalChar := parser.Characteristic{
				Name:  "Оригинал",
				Key:   "pm_isoriginal",
				Value: &parser.Characteristic_ListVal{ListVal: pmutils.BoolToStringList(item.Item.GetOriginal())},
			}
			chars = append(chars, &originalChar)

			adultChar := parser.Characteristic{
				Name:  "Для взрослых",
				Key:   "pm_isadult",
				Value: &parser.Characteristic_ListVal{ListVal: pmutils.BoolToStringList(item.Item.GetIsAdult())},
			}
			chars = append(chars, &adultChar)

			priceChar := parser.Characteristic{
				Name:  "Цена",
				Key:   "pm_price",
				Value: &parser.Characteristic_NumVal{NumVal: item.Item.Price},
			}
			chars = append(chars, &priceChar)
			item.Chars = append(item.Chars, chars...)
		}(item)
	}
	wg.Wait()
}
