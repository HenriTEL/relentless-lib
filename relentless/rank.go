package relentless

import (
	"encoding/json"
	"strings"
)

var ranks []string

// Rank represent the how many instances of a card you can have in your deck
type Rank int8

// ranks
var (
	COMMANDER    = newRank("COMMANDER")
	GENERAL      = newRank("GENERAL")
	MINION       = newRank("MINION")
	OFFICER      = newRank("OFFICER")
	UNKNOWN_RANK = newRank("UNKNOWN_RANK")
)

func (r Rank) String() string {
	return ranks[int8(r)]
}

func newRank(s string) Rank {
	ranks = append(ranks, s)
	return Rank(len(ranks) - 1)
}

func toRank(s string) Rank {
	for i, e := range ranks {
		if e == s {
			return Rank(i)
		}
	}
	return UNKNOWN_RANK
}

// UnmarshalJSON converts a json []byte into a Type
func (r *Rank) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*r = toRank(strings.ToUpper(s))
	return nil
}
