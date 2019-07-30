package relentless

import (
	"encoding/json"
	"strings"
)

var sale_types []string

// SaleType represent the type of a card
type SaleType int8

// SaleTypes
var (
	DIRECT            = newSaleType("DIRECT")
	UNKNOWN_SALE_TYPE = newSaleType("UNKNOWN_SALE_TYPE")
)

func (st SaleType) String() string {
	return sale_types[int8(st)]
}

func newSaleType(s string) SaleType {
	sale_types = append(sale_types, s)
	return SaleType(len(sale_types) - 1)
}

func toSaleType(s string) SaleType {
	for i, st := range sale_types {
		if st == s {
			return SaleType(i)
		}
	}
	return UNKNOWN_SALE_TYPE
}

// UnmarshalJSON converts a json []byte into a SaleType
func (st *SaleType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*st = toSaleType(strings.ToUpper(s))
	return nil
}
