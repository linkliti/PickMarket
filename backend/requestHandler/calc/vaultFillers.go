package calc

import (
	"log/slog"
	"protos/parser"
	"reflect"
	"strings"
)

func (c *calc) getVaultByKey(key string, create bool) *calcVault {
	key = strings.ToLower(key)
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
		switch pref.Value.(type) {
		case *parser.UserPref_NumVal:
			v.prefType = NUM_TYPE
		case *parser.UserPref_ListVal:
			v.prefType = LIST_TYPE
		}
		slog.Debug("Set vault", "key", pref.Key, "type", v.prefType.Name())
	}
}

func (c *calc) fillChars() {
	for _, item := range c.itemList {
		for _, char := range item.Chars {
			v := c.getVaultByKey(char.Key, false)
			if v == nil {
				continue
			}
			// Check type using UserPref types
			var charType reflect.Type
			switch char.Value.(type) {
			case *parser.Characteristic_NumVal:
				charType = NUM_TYPE
			case *parser.Characteristic_ListVal:
				charType = LIST_TYPE
			}
			if charType != v.prefType {
				slog.Warn("fillChars: characteristic type mismatch", "key", char.Key, "expected", v.prefType.Name(), "actual", charType.Name())
				continue
			}
			v.charPointers = append(v.charPointers, char)
		}
	}
}

func (c *calc) removeVaultsWithoutChars() {
	for k, v := range c.vaults {
		if len(v.charPointers) == 0 {
			delete(c.vaults, k)
			slog.Debug("Removed vault", "key", k)
		}
	}
}
