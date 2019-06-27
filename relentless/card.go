package relentless

type Card struct {
    Id   int
    MouldId string `json:"mould_id"`
    Version string
    Kind string
    Set string
    Name string
    Description string
    FlavorText string `json:"flavor_text"`
    Picture string
    Rank string
    Type string
    Rarity string
    Frame string
    Damage int
    Health int
    Cost int
    BlockHeight int `json:"block_height"`
    Ability string
    ImageUrl string `json:"image_url"`
}

type CardsResponse struct {
    Total   int
    Page   int
    Limit   int
    Cards []Card
}
