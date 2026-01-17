package game

import "math/rand"

// Deck represents a collection of 52 cards.
type Deck struct {
	Cards []Card
}

// NewDeck creates a new standard 52-card deck.
func NewDeck() *Deck {
	suits := []Suit{Clubs, Diamonds, Hearts, Spades}
	ranks := []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
	deck := &Deck{}
	for _, suit := range suits {
		for _, rank := range ranks {
			deck.Cards = append(deck.Cards, Card{Suit: suit, Rank: rank})
		}
	}
	return deck
}

// Shuffle shuffles the deck.
func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

// Deal removes and returns the top n cards from the deck.
func (d *Deck) Deal(n int) []Card {
	if n > len(d.Cards) {
		n = len(d.Cards)
	}
	cards := d.Cards[:n]
	d.Cards = d.Cards[n:]
	return cards
}
