package relentless

import (
	"encoding/json"
	"strings"
)

var types []string

// Type represent the type of a card
type Type int8

// Types
var (
	FERAL        = newType("FERAL")
	HEAVY        = newType("HEAVY")
	UNKNOWN_TYPE = newType("UNKNOWN_TYPE")
	WALKER       = newType("WALKER")
)

func (t Type) String() string {
	return types[int8(t)]
}

func newType(s string) Type {
	types = append(types, s)
	return Type(len(types) - 1)
}

func toType(s string) Type {
	for i, e := range types {
		if e == s {
			return Type(i)
		}
	}
	return UNKNOWN_TYPE
}

// UnmarshalJSON converts a json []byte into a Type
func (t *Type) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*t = toType(strings.Title(s))
	return nil
}
