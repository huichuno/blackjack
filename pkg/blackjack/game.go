package blackjack

import (
	"bufio"
	"fmt"
	"os"
)

// Game struct with Players and Deck fields
type Game struct {
	Players []*Player // list of pointer players in the game
	Deck    *Deck     // pointer to card deck
}

// pop the top of deck
func (g *Game) pop() *Card {
	return g.Deck.Pop()
}

// clear screen for linux terminal
func (g *Game) clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// display all players' card, status, and points
func (g *Game) displayPlayersInfo(dealerCardHidden bool) {
	for _, player := range g.Players {
		// normal player
		if player.Type == Normal {
			fmt.Printf(" [%v]\t: ", player.Name)
			for _, card := range player.CardList {
				fmt.Printf(" %-3v", card.Name)
			}

			fmt.Printf("%-3v[", " ")
			for cardValue := range player.CardValue {
				fmt.Printf(" %-2v ", cardValue)
			}
			fmt.Printf("]\n")
			fmt.Printf(" [   Status]\t:  %v\n", player.DisplayStatus)
			fmt.Printf(" [   Points]\t:  %v\n\n", player.TotalPoints)

		} else {
			// dealer
			fmt.Printf(" [%v]\t: ", player.Name)
			for i, card := range player.CardList {
				if i == 0 && dealerCardHidden {
					fmt.Printf(" (*)")
				} else {
					fmt.Printf(" %-3v", card.Name)
				}
			}

			if dealerCardHidden {
				fmt.Printf("%-3v[ ? ]\n", " ")
				fmt.Printf(" [   Status]\t:  %v\n", player.DisplayStatus)
			} else {
				fmt.Printf("%-3v[", " ")
				for cardValue := range player.CardValue {
					fmt.Printf(" %-2v ", cardValue)
				}
				fmt.Printf("]\n")
				fmt.Printf(" [   Status]\t:  %v\n", player.DisplayStatus)
			}
		}

	}
	fmt.Println()
}

// find winner and loser against the dealer
func (g *Game) findWinnerLoser(dealerCardValue int) {
	for _, player := range g.Players {

		// skip dealer, we only care about normal player in this func
		if player.Type == Dealer {
			continue
		}
		// players with status == Exceeded, already have their points deducted
		if player.Status == Exceeded {
			player.DisplayStatus = DisplayLose
			continue
		}

		// for player with card value <= 21
		if dealerCardValue > 21 {
			player.win()
			player.DisplayStatus = DisplayWin

		} else {
			// player with card value > 21
			playerCardValue := player.getCardValue()

			if playerCardValue > dealerCardValue {
				player.win()
				player.DisplayStatus = DisplayWin
			} else if playerCardValue < dealerCardValue {
				player.lose()
				player.DisplayStatus = DisplayLose
			} else {
				player.DisplayStatus = DisplayPush
			}
		}
	}
}

func (g *Game) displayGameInfo(num int, dealerCardHidden bool) {
	g.clearScreen()
	fmt.Printf("Blacjack Game #%v \n\n", num)
	g.displayPlayersInfo(dealerCardHidden)
}

func (g *Game) Start(num int) error {

	g.Deck.Top = FirstCard
	g.Deck.Bottom = LastCard - CardRemainingToTriggerReschuffle

	// interate thru number of gaming round
	for i := 1; i < num+1; i++ {

		// set lock to top of deck at begginng of each round
		g.Deck.Lock = g.Deck.Top

		// dealing 1st card to all players
		for _, player := range g.Players {
			// since we are iterating thru all players, let reset player's card,
			// cardvalue and status at beginning of each round,
			player.clear()

			card := g.pop()
			player.addCard(card)
		}

		// dealing 2nd card..
		for _, player := range g.Players {
			card := g.pop()
			player.addCard(card)
		}

		var dealerCardValue int
		var loseCount int // count number of losing player

		for _, player := range g.Players {
			player.DisplayStatus = DisplayNow

			// normal player's draw loop
			if player.Type == Normal {
				for {
					g.displayGameInfo(i, true)
					fmt.Printf("Press [h] to Hit, [s] to Stay, [q] to Quit. [%v] : ", player.Name)

					// read keyboard input
					reader := bufio.NewReader(os.Stdin)
					c, _, err := reader.ReadRune()
					if err != nil {
						return err
					}

					if c == 'h' {
						player.Status = Undefined
						card := g.pop()
						player.addCard(card)
					} else if c == 's' {
						break
					} else if c == 'q' {
						os.Exit(0)
					} else {
						continue
					}

					// break loop if player's card value exceeded 21
					if player.Status == Exceeded {
						player.lose()
						player.DisplayStatus = DisplayLose
						loseCount++
						break
					}

					// break loop when max number of card reached
					cardCount := len(player.CardList)
					if cardCount >= MaxCardAllowed {
						break
					}
				}

			} else {
				// dealer's draw loop
				// all players have lost, nothing to do for dealer
				if loseCount == len(g.Players)-1 {
					continue
				}

				for {
					g.displayGameInfo(i, false)
					fmt.Printf("Press [Enter] to continue... [%v] : ", player.Name)
					fmt.Scanln()

					// break if dealer's card value is larger than 16
					dealerCardValue = player.getCardValue()
					if dealerCardValue > 16 {
						break
					}

					player.Status = Undefined
					card := g.pop()
					player.addCard(card)
				}
			}
			if player.DisplayStatus == DisplayNow {
				player.DisplayStatus = Blank
			}
		}
		// find winners and losers
		g.findWinnerLoser(dealerCardValue)

		g.displayGameInfo(i, false)
		fmt.Printf("Press [Enter] to continue...")
		fmt.Scanln()
	}

	fmt.Printf("\n** Thanks for playing Blackjack **\n\n")
	return nil
}
