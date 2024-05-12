package calc

import (
	"protos/parser"
	"reflect"
)

type calcVault struct {
	prefType     reflect.Type
	prefPointer  *parser.UserPref
	charPointers []*parser.Characteristic
}

var (
	LIST_TYPE = reflect.TypeOf(parser.UserPref_ListVal{})
	NUM_TYPE  = reflect.TypeOf(parser.UserPref_NumVal{})
)

type calc struct {
	itemList []*parser.ItemExtended
	userPref map[string]*parser.UserPref
	vaults   map[string]*calcVault
	req      *parser.ItemsRequest
}

type caseStruct struct {
	fn       func(*parser.Characteristic, *parser.UserPref)
	calcType reflect.Type
}
