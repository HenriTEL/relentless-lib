package relentless

import (
	"log"
	"testing"
)

func TestListCards(t *testing.T) {
	client, _ := NewClient("https://api.loom.games/zb/v1/")
	cards, err := client.ListCards()
	if err != nil {
		log.Fatal(err)
	}
	if len(cards) < 183 {
		t.Errorf("There should be at least 183 cars, got %d", len(cards))
	}
}
