package calc

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
