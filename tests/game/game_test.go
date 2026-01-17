package game

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	game := NewGame()

	if len(game.Tableau.Columns) != 7 {
		t.Errorf("Expected 7 tableau columns, got %d", len(game.Tableau.Columns))
	}

	for i, col := range game.Tableau.Columns {
		if len(col) != i+1 {
			t.Errorf("Expected %d cards in column %d, got %d", i+1, i, len(col))
		}
	}

	if len(game.Stock.Cards) != 52-28 {
		t.Errorf("Expected %d cards in stock, got %d", 52-28, len(game.Stock.Cards))
	}
}

func TestValidateMove(t *testing.T) {
	game := NewGame()

	// Valid move
	fromCard := Card{Suit: Hearts, Rank: Queen}
	toCard := Card{Suit: Clubs, Rank: King}
	if !game.ValidateMove(fromCard, toCard) {
		t.Errorf("Expected move to be valid, but it was invalid")
	}

	// Invalid move (same color)
	fromCard = Card{Suit: Hearts, Rank: Queen}
	toCard = Card{Suit: Diamonds, Rank: King}
	if game.ValidateMove(fromCard, toCard) {
		t.Errorf("Expected move to be invalid, but it was valid")
	}

	// Invalid move (wrong rank)
	fromCard = Card{Suit: Hearts, Rank: Jack}
	toCard = Card{Suit: Clubs, Rank: King}
	if game.ValidateMove(fromCard, toCard) {
		t.Errorf("Expected move to be invalid, but it was valid")
	}
}
