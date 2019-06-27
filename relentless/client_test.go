package relentless

import (
  "log"
  "net/url"
	"net/http"
  "testing"
)

func TestListCards(t *testing.T) {
  apiEndpoint, err := url.Parse("https://api.loom.games/zb/v1/")
  if err != nil {
		log.Fatal(err)
	}
  client := Client{apiEndpoint, http.DefaultClient}
  cards, err := client.ListCards()
  if err != nil {
		log.Fatal(err)
	}
  if len(cards) < 183 {
       t.Errorf("There should be at least 183 cars, got %d", len(cards))
   }
}
