package relentless

type Card struct {
	Id          uint16
	MouldId     string `json:"mould_id"`
	Version     string
	Kind        Kind
	Set         string
	Name        string
	Description string
	FlavorText  string `json:"flavor_text"`
	Picture     string
	Rank        string
	Type        string
	Rarity      string
	Frame       string
	Damage      uint8
	Health      uint8
	Cost        uint8
	BlockHeight uint8 `json:"block_height"`
	Ability     string
	ImageUrl    string `json:"image_url"`
}
