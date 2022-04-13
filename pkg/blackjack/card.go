package blackjack

import (
	"fmt"
	"math/rand"
	"time"
)

// Card struct
type Card struct {
	Name  string
	Value int
}

// use hashtable (Cards) as cards' container for quick access, only one Deck hastable is being created
type Deck struct {
	Cards  map[int]*Card // container for deck of cards
	Top    int           // point to top of deck
	Bottom int           // point to bottom of deck
	Lock   int           // point to start of cards in the deck that are in players' hand.
	// Lock is set to Top at the start of each game round
}

const (
	FirstCard                        = 1
	LastCard                         = 52
	MaxCardAllowed                   = 5
	CardRemainingToTriggerReschuffle = 15
	Ace                              = 1
)

// swap between 2 cards in the hastable, used for card shuffling and moving Lock cards to bottom of deck
func (d *Deck) swap(i int, j int) {
	temp := d.Cards[i]
	d.Cards[i] = d.Cards[j]
	d.Cards[j] = temp
}

// simple Shuffling method
func (d *Deck) Shuffle(top int, bottom int) {

	// shuffle 3 times
	for count := 0; count < 2; count++ {
		rand.Seed(time.Now().UnixNano())

		for i := top; i <= bottom; i++ {
			j := rand.Intn(bottom-top) + top
			d.swap(i, j)
		}
	}
}

// cards on players' hand (dealed) are move to bottom of Deck stack
func (d *Deck) moveDealedCards(start int, end int) {
	count := end - start + 1
	i := LastCard
	j := end

	for {
		d.swap(i, j)
		i--
		j--
		count--
		if count <= 0 {
			break
		}
	}
}

// pop the top of the Deck. Top points to the key of the hashtable to allow quick access to
// pointer of Card
func (d *Deck) Pop() *Card {

	card := d.Cards[d.Top]
	d.Top++

	// when top reaches bottom perform reschuffling
	if d.Top >= d.Bottom {

		// move dealed player's card to bottom of Deck stack
		d.moveDealedCards(d.Lock, d.Top-1)
		dealedCardCount := d.Top - d.Lock

		d.Top = 1
		d.Bottom = LastCard - dealedCardCount - CardRemainingToTriggerReschuffle
		d.Lock = LastCard - dealedCardCount + 1

		d.Shuffle(d.Top, d.Bottom)
	}

	return card
}

// view Deck content, used for debugging
func (d *Deck) PrintDeck() {

	fmt.Printf("\nTop: %v, Bottom: %v, Lock: %v\n", d.Top, d.Bottom, d.Lock)

	for i := 1; i < 11; i++ {
		fmt.Printf(" [%2d]", i)
	}
	fmt.Println()

	for i := FirstCard; i <= LastCard; i++ {
		fmt.Printf(" %4s", d.Cards[i].Name)
		if i%10 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}
