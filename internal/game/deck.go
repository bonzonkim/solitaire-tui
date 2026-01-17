package game

// NewDeck creates a new standard 52-card deck.
func NewDeck() []*Card {
	suits := []Suit{Clubs, Diamonds, Hearts, Spades}
	ranks := []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
	deck := make([]*Card, 0, 52)
	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, &Card{Suit: suit, Rank: rank})
		}
	}
	return deck
}
