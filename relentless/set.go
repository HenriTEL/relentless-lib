package relentless

import (
	"encoding/json"
	"strings"
)

var sets []string

// Set represent the set the card comes from
type Set int8

// sets
var (
	AIR         = newSet("AIR")
	EARTH       = newSet("EARTH")
	FIRE        = newSet("FIRE")
	ITEM        = newSet("ITEM")
	LIFE        = newSet("LIFE")
	OTHERS      = newSet("OTHERS")
	TOXIC       = newSet("TOXIC")
	UNKNOWN_SET = newSet("UNKNOWN_SET")
	WATER       = newSet("WATER")
)

func (s Set) String() string {
	return sets[int8(s)]
}

func newSet(s string) Set {
	sets = append(sets, s)
	return Set(len(sets) - 1)
}

func toSet(s string) Set {
	for i, e := range sets {
		if e == s {
			return Set(i)
		}
	}
	return UNKNOWN_SET
}

// UnmarshalJSON converts a json []byte into a Set
func (s *Set) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	*s = toSet(strings.ToUpper(str))
	return nil
}
