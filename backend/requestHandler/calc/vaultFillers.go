package calc

import "protos/parser"

func (c *calc) getVaultByKey(key string, create bool) *calcVault {
	if v, ok := c.vaults[key]; ok {
		return v
	}
	if create {
		v := &calcVault{
			prefPointer:  nil,
			charPointers: make([]*parser.Characteristic, 0),
		}
		c.vaults[key] = v
		return v
	}
	return nil
}

func (c *calc) fillPreferences() {
	for _, pref := range c.userPref {
		v := c.getVaultByKey(pref.Key, true)
		v.prefPointer = pref
	}
}

func (c *calc) fillChars() {
	for _, item := range c.itemList {
		for _, char := range item.Chars {
			v := c.getVaultByKey(char.Key, false)
			if v == nil {
				continue
			}
			v.charPointers = append(v.charPointers, char)
		}
	}
}
