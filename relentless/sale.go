package relentless

type Sale struct {
	Currency  Currency
	Amount    uint8   `json:"amount_give,string"`
	UnitPrice float64 `json:"price_per_unit,string"`
	Type      SaleType
}
