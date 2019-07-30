package relentless

type Sale struct {
	Currency Currency
	Amount   uint8
	Price    float64
	Type     SaleType
}
