package relentless

import (
	"encoding/json"
	"strings"
)

var kinds []string

// Kind represent the type of a card
type Kind int8

// Kinds
var (
	CREATURE     = newKind("CREATURE")
	SPELL        = newKind("SPELL")
	UNKNOWN_KIND = newKind("UNKNOWN_KIND")
)

func (k Kind) String() string {
	return kinds[int8(k)]
}

func newKind(s string) Kind {
	kinds = append(kinds, s)
	return Kind(len(kinds) - 1)
}

func toKind(s string) Kind {
	for i, e := range kinds {
		if e == s {
			return Kind(i)
		}
	}
	return UNKNOWN_KIND
}

// UnmarshalJSON converts a json []byte into a Type
func (k *Kind) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*k = toKind(strings.ToUpper(s))
	return nil
}
