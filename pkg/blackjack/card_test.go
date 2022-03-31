package blackjack_test

import (
	"testing"

	"github.com/huichuno/blackjack/pkg/blackjack"
)

func TestPop(t *testing.T) {

	deck := blackjack.Deck{
		Cards: map[int]*blackjack.Card{
			1: {Name: "A♠", Value: 1}, 2: {Name: "2♠", Value: 2}, 3: {Name: "3♠", Value: 3},
			4: {Name: "4♠", Value: 4}, 5: {Name: "5♠", Value: 5}, 6: {Name: "6♠", Value: 6},
			7: {Name: "7♠", Value: 7}, 8: {Name: "8♠", Value: 8}, 9: {Name: "9♠", Value: 9},
			10: {Name: "10♠", Value: 10}, 11: {Name: "J♠", Value: 10}, 12: {Name: "Q♠", Value: 10},
			13: {Name: "K♠", Value: 10},
			14: {Name: "A♡", Value: 1}, 15: {Name: "2♡", Value: 2}, 16: {Name: "3♡", Value: 3},
			17: {Name: "4♡", Value: 4}, 18: {Name: "5♡", Value: 5}, 19: {Name: "6♡", Value: 6},
			20: {Name: "7♡", Value: 7}, 21: {Name: "8♡", Value: 8}, 22: {Name: "9♡", Value: 9},
			23: {Name: "10♡", Value: 10}, 24: {Name: "J♡", Value: 10}, 25: {Name: "Q♡", Value: 10},
			26: {Name: "K♡", Value: 10},
			27: {Name: "A♣", Value: 1}, 28: {Name: "2♣", Value: 2}, 29: {Name: "3♣", Value: 3},
			30: {Name: "4♣", Value: 4}, 31: {Name: "5♣", Value: 5}, 32: {Name: "6♣", Value: 6},
			33: {Name: "7♣", Value: 7}, 34: {Name: "8♣", Value: 8}, 35: {Name: "9♣", Value: 9},
			36: {Name: "10♣", Value: 10}, 37: {Name: "J♣", Value: 10}, 38: {Name: "Q♣", Value: 10},
			39: {Name: "K♣", Value: 10},
			40: {Name: "A♦", Value: 1}, 41: {Name: "2♦", Value: 2}, 42: {Name: "3♦", Value: 3},
			43: {Name: "4♦", Value: 4}, 44: {Name: "5♦", Value: 5}, 45: {Name: "6♦", Value: 6},
			46: {Name: "7♦", Value: 7}, 47: {Name: "8♦", Value: 8}, 48: {Name: "9♦", Value: 9},
			49: {Name: "10♦", Value: 10}, 50: {Name: "J♦", Value: 10}, 51: {Name: "Q♦", Value: 10},
			52: {Name: "K♦", Value: 10},
		},
		Top:    1,
		Bottom: 11,
		Lock:   1,
	}

	for i := 0; i < 10; i++ {
		deck.Pop()
	}

	if deck.Top != 1 {
		t.Errorf("Top value error. Expected: 1, Actual: %v", deck.Top)
	}
	if deck.Bottom != 27 {
		t.Errorf("Bottom value error. Expected: 27, Actual: %v", deck.Bottom)
	}
	if deck.Lock != 43 {
		t.Errorf("Lock value error. Expected: 43, Actual: %v", deck.Lock)
	}
	if deck.Cards[43].Name != "A♠" {
		t.Errorf("Cards[43].Name value error. Expected: A♠, Actual: %v", deck.Cards[43])
	}
	if deck.Cards[52].Name != "10♠" {
		t.Errorf("Cards[52].Name value error. Expected: A♠, Actual: %v", deck.Cards[52])
	}
}

// todo: add more unit tests
