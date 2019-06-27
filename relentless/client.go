package relentless

import (
  "encoding/json"
  "net/http"
  "net/url"
  "strconv"
)

type Client struct {
    apiEndpoint *url.URL
    httpClient *http.Client
}

func (c *Client) ListCards() ([]Card, error) {
    var cards []Card
    var cards_response CardsResponse
    for page := 1; page < 2 || len(cards_response.Cards) > 0 ; page += 1 {
      cards_url := c.apiEndpoint.ResolveReference(&url.URL{Path: "cards",
                                                           RawQuery: "page=" + strconv.Itoa(page)})
      resp, err := http.Get(cards_url.String())
      if err != nil {
          return nil, err
      }
      err = json.NewDecoder(resp.Body).Decode(&cards_response)
      if err != nil {
        panic(err)
      }
      cards = append(cards, cards_response.Cards...)
    }
    return cards, nil
}
