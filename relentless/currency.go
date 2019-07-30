package relentless

import (
	"encoding/json"
	"strings"
)

var currencies []string

// Currency represent a cryptocurrency
type Currency int8

// Currencies
var (
	ETH              = newCurrency("ETH")
	UNKNOWN_CURRENCY = newCurrency("UNKNOWN_CURRENCY")
)

func (c Currency) String() string {
	return currencies[int8(c)]
}

func newCurrency(s string) Currency {
	currencies = append(currencies, s)
	return Currency(len(currencies) - 1)
}

func toCurrency(s string) Currency {
	for i, c := range currencies {
		if c == s {
			return Currency(i)
		}
	}
	return UNKNOWN_CURRENCY
}

// UnmarshalJSON converts a json []byte into a Currency
func (c *Currency) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*c = toCurrency(strings.ToUpper(s))
	return nil
}
