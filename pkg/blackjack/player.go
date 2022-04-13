package blackjack

const (
	// player Type
	Dealer = 0
	Normal = 1

	// player Status
	Exceeded  = 2 // status when player's card value > 21
	NotExceed = 3 // status when player's card value <= 21
	Undefined = 4 // initial state

	// Player display status. Only use for display
	DisplayWin  = "<-- win -->"
	DisplayLose = "<-- lose -->"
	DisplayPush = "<-- push -->" // draw
	Blank       = ""
	DisplayNow  = "<-- now -->"
)

// Player struct with player's properties
type Player struct {
	Name          string       // player's name
	Type          int          // either normal player or dealer (only 1 dealer allowed)
	CardValue     map[int]bool // using map as set to store player's card values
	CardList      []*Card      // list of player's card on hand
	TotalPoints   int          // total player accumulated point
	Status        int          // track player status: Exceeded, NotExceed, Undefined
	DisplayStatus string       // for display only: DisplayWin, DisplayLose, DisplayPush, Blank
}

// call this method at the beginning of each game round
func (p *Player) clear() {

	// clear CardList
	p.CardList = p.CardList[:0]

	// clear CardValue
	for k := range p.CardValue {
		delete(p.CardValue, k)
	}

	// reset status
	p.Status = Undefined
	p.DisplayStatus = Blank
}

// return closest card value to number 21
func (p *Player) getCardValue() int {

	var left int = 0
	var right int = 10000 // aribitary large number, used for comparison

	// return 0 since no card on player's hand
	if len(p.CardValue) == 0 {
		return 0
	}

	// finding the nearest number to 21, both larger and smaller than 21
	for k := range p.CardValue {
		if k <= 21 {
			if k > left {
				left = k
			}
		} else {
			if k < right {
				right = k
			}
		}
	}

	if left != 0 {
		return left
	} else {
		return right
	}
}

// add points to player
func (p *Player) win() {

	// if player's card values equal 21 and only 2 cards on hand
	if len(p.CardList) == 2 && p.CardValue[21] {
		p.TotalPoints += 15
	} else {
		p.TotalPoints += 10
	}
}

// deduct points from player
func (p *Player) lose() {

	p.TotalPoints -= 10
}

// add card to player's hand and calculate card values each time a new card is added
func (p *Player) addCard(card *Card) {

	p.CardList = append(p.CardList, card)

	// only when first card being added
	if len(p.CardValue) == 0 {

		if card.Value != Ace {
			p.CardValue[card.Value] = true
		} else {
			p.CardValue[1] = true
			p.CardValue[11] = true
		}
		p.Status = NotExceed
		return
	}

	// use set "cache" to temporary store card values
	// Prioirty of setting player's status is "NotExceed" > "Exceeded" > "Undefined"
	// Player status is set to "Undefined" before addCard() is called, see game.go
	keys := []int{}
	cache := make(map[int]bool)

	for k := range p.CardValue {
		// not ace
		if card.Value != Ace {
			cache[k+card.Value] = true

			// check if Exceeded 21. carryin out check here to reduce iterating thru CardValue
			if k+card.Value <= 21 {
				p.Status = NotExceed
			} else {
				if p.Status == Undefined {
					p.Status = Exceeded
				}
			}
		} else {
			cache[k+1] = true
			cache[k+11] = true

			// check if Exceeded 21
			if k+1 <= 21 {
				p.Status = NotExceed
			} else {
				if p.Status == Undefined {
					p.Status = Exceeded
				}
			}
		}
		keys = append(keys, k)
	}
	// delete old values in CardValue set
	for _, v := range keys {
		delete(p.CardValue, v)
	}

	// set new card values to set
	for k, v := range cache {
		p.CardValue[k] = v
	}
}
