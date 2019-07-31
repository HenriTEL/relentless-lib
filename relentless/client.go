package relentless

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

const Nb_editions = 5

type Client struct {
	ApiEndpoint *url.URL
}

func NewClient(endpoint string) (*Client, error) {
	apiEndpoint, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	return &Client{apiEndpoint}, nil
}

type CardsResponse struct {
	Total uint16
	Page  uint16
	Limit uint16
	Cards []Card
}

func (c *Client) ListCards() ([]Card, error) {
	var cards []Card
	var cards_response CardsResponse
	for page := 1; page < 2 || len(cards_response.Cards) > 0; page += 1 {
		cards_url := c.ApiEndpoint.ResolveReference(&url.URL{Path: "cards",
			RawQuery: "page=" + strconv.Itoa(page)})
		resp, err := http.Get(cards_url.String())
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(&cards_response)
		if err != nil {
			panic(err)
		}
		for _, card := range cards_response.Cards {
			for i := 0; i < Nb_editions; i++ {
				card.Edition = Edition(i)
				card.SetSales()
				cards = append(cards, card)
			}
		}
	}
	return cards, nil
}

type SalesResponse struct {
	Sales []Sale `json:"records"`
}

func (c *Client) ListSales(card_id int) ([]Sale, error) {
	var sales []Sale
	var sales_response SalesResponse
	for page := 1; page < 2 || len(sales_response.Sales) > 0; page += 1 {
		sales_url := c.ApiEndpoint.ResolveReference(&url.URL{Path: "trade/available-amount",
			RawQuery: "page=" + strconv.Itoa(page) + "&cardID=" + strconv.Itoa(card_id)})
		resp, err := http.Get(sales_url.String())
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(&sales_response)
		if err != nil {
			panic(err)
		}
		for _, sale := range sales_response.Sales {
			sale.Type = DIRECT
			sale.Currency = ETH
			sales = append(sales, sale)
		}
	}
	return sales, nil
}
