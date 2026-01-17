package game

// Game represents the state of the Solitaire game.
type Game struct {
	Deck       *Deck
	Tableau    Tableau
	Stock      Stock
	Waste      Waste
	Foundation Foundation
}

// NewGame creates a new game of Solitaire.
func NewGame() *Game {
	deck := NewDeck()
	deck.Shuffle()

	tableau := Tableau{Columns: make([][]Card, 7)}
	for i := 0; i < 7; i++ {
		tableau.Columns[i] = deck.Deal(i + 1)
		for j := 0; j < i; j++ {
			tableau.Columns[i][j].FaceUp = false
		}
		tableau.Columns[i][i].FaceUp = true
	}

	stock := Stock{Cards: deck.Cards}
	waste := Waste{Cards: []Card{}}
	foundation := Foundation{Piles: make([][]Card, 4)}

	return &Game{
		Deck:       deck,
		Tableau:    tableau,
		Stock:      stock,
		Waste:      waste,
		Foundation: foundation,
	}
}

// ValidateMove checks if a card can be placed on another card in the tableau.
func (g *Game) ValidateMove(fromCard, toCard Card) bool {
	// Must be opposite colors
	if (fromCard.Suit == Clubs || fromCard.Suit == Spades) == (toCard.Suit == Clubs || toCard.Suit == Spades) {
		return false
	}
	// Rank must be one less
	return rankToInt(fromCard.Rank) == rankToInt(toCard.Rank)-1
}

func rankToInt(r Rank) int {
	switch r {
	case Ace:
		return 1
	case Two:
		return 2
	case Three:
		return 3
	case Four:
		return 4
	case Five:
		return 5
	case Six:
		return 6
	case Seven:
		return 7
	case Eight:
		return 8
	case Nine:
		return 9
	case Ten:
		return 10
	case Jack:
		return 11
	case Queen:
		return 12
	case King:
		return 13
	}
	return 0
}

// IsUnwinnable checks if the game is in an unwinnable state.
func (g *Game) IsUnwinnable() bool {
	// Basic implementation for now
	// A more sophisticated check will be implemented later
	return false
}

// IsWon checks if the game has been won.
func (g *Game) IsWon() bool {
	for _, pile := range g.Foundation.Piles {
		if len(pile) != 13 {
			return false
		}
	}
	return true
}
