package relentless

import (
	"encoding/json"
	"strings"
)

var editions []string

// Edition represent the edition of a card
type Edition int8

// Editions
var (
	STANDARD        = newEdition("STANDARD")
	BACKER          = newEdition("BACKER")
	LIMITED         = newEdition("LIMITED")
	BINANCE         = newEdition("BINANCE")
	TRON            = newEdition("TRON")
	UNKNOWN_EDITION = newEdition("UNKNOWN_EDITION")
)

func (e Edition) String() string {
	return editions[int8(e)]
}

func newEdition(s string) Edition {
	editions = append(editions, s)
	return Edition(len(editions) - 1)
}

func toEdition(s string) Edition {
	for i, e := range editions {
		if e == s {
			return Edition(i)
		}
	}
	return UNKNOWN_EDITION
}

// UnmarshalJSON converts a json []byte into a Edition
func (e *Edition) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*e = toEdition(strings.ToUpper(s))
	return nil
}
