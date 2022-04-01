package main

import (
	"fmt"
	"os"

	"github.com/huichuno/blackjack/pkg/blackjack"
)

// factory function to create blackjack deck
func createCardDeck() *blackjack.Deck {

	// construct deck
	return &blackjack.Deck{
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
		Top:    blackjack.FirstCard,
		Bottom: blackjack.LastCard,
		Lock:   blackjack.FirstCard,
	}
}

// factory function to register blackjack player
func registerPlayers() []*blackjack.Player {

	return []*blackjack.Player{
		{
			Name:          "John Wick",
			Type:          blackjack.Normal,
			CardValue:     make(map[int]bool),
			CardList:      []*blackjack.Card{},
			TotalPoints:   0,
			Status:        blackjack.Undefined,
			DisplayStatus: blackjack.Blank,
		},
		{
			Name:          "  GI Jane",
			Type:          blackjack.Normal,
			CardValue:     make(map[int]bool),
			CardList:      []*blackjack.Card{},
			TotalPoints:   0,
			Status:        blackjack.Undefined,
			DisplayStatus: blackjack.Blank,
		},
		{
			Name:          "Shang Chi",
			Type:          blackjack.Normal,
			CardValue:     make(map[int]bool),
			CardList:      []*blackjack.Card{},
			TotalPoints:   0,
			Status:        blackjack.Undefined,
			DisplayStatus: blackjack.Blank,
		},
		{
			Name:          "    Alice",
			Type:          blackjack.Normal,
			CardValue:     make(map[int]bool),
			CardList:      []*blackjack.Card{},
			TotalPoints:   0,
			Status:        blackjack.Undefined,
			DisplayStatus: blackjack.Blank,
		},
		{
			Name:          "   Dealer",
			Type:          blackjack.Dealer, // dealer
			CardValue:     make(map[int]bool),
			CardList:      []*blackjack.Card{},
			TotalPoints:   0,
			Status:        blackjack.Undefined,
			DisplayStatus: blackjack.Blank,
		},
	}
}

// factory function to create blackjack game
func createGame(players []*blackjack.Player, deck *blackjack.Deck) *blackjack.Game {

	return &blackjack.Game{
		Players: players,
		Deck:    deck,
	}
}

func main() {
	// create blackjack deck
	deck := createCardDeck()

	// shuffle deck
	deck.Shuffle(deck.Top, deck.Bottom)

	// register players. add or remove player in registerPlayers() func
	// make sure  only 1 dealer among players
	players := registerPlayers()

	// create game engine
	blackjackGame := createGame(players, deck)

	// start game. change numberOfRound value to reflect number of round you want to play
	numberOfRound := 5
	err := blackjackGame.Start(numberOfRound)

	if err != nil {
		fmt.Printf("Games start error: %v", err)
		os.Exit(1)
	}
}
