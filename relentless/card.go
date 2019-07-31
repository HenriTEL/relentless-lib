package relentless

type Card struct {
	Ability     string
	BlockHeight uint8 `json:"block_height"`
	Cost        uint8
	Damage      uint8
	Description string
	Edition     Edition
	FlavorText  string `json:"flavor_text"`
	Frame       string
	Health      uint8
	Id          uint16
	ImageUrl    string `json:"image_url"`
	Kind        Kind
	MouldId     string `json:"mould_id"`
	Name        string
	Picture     string
	Rank        Rank
	Rarity      string
	Sales       []Sale
	Set         Set
	Type        Type
	Version     string
}

func (c *Card) SetSales() {
	client, _ := NewClient("https://auth.loom.games/")
	sale_id := int(c.Id)*10 + int(c.Edition)
	sales, _ := client.ListSales(sale_id)
	c.Sales = sales
}
