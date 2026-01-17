package game

import (
	"math/rand"
)

const (
	StockPile = 0
	WastePile = 1
	// Foundation piles
	FoundationPile1 = 2
	FoundationPile2 = 3
	FoundationPile3 = 4
	FoundationPile4 = 5
	// Tableau piles
	TableauPile1 = 6
	TableauPile2 = 7
	TableauPile3 = 8
	TableauPile4 = 9
	TableauPile5 = 10
	TableauPile6 = 11
	TableauPile7 = 12
)

// Game represents the state of the Solitaire game.
type Game struct {
	Stock       Pile
	Waste       Pile
	Foundations [4]Pile
	Tableaus    [7]Pile

	IsWon      bool
	ActivePile int // Using an index for now; could be an enum
	ActiveCard int // Index of the card in the active pile
}

// NewGame creates a new game of Solitaire.
func NewGame() *Game {
	// Create and shuffle a standard 52-card deck.
	deck := NewDeck()
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	g := &Game{
		IsWon:      false,
		ActivePile: -1, // No pile selected initially
		ActiveCard: -1, // No card selected initially
	}

	// Deal cards to the seven tableau piles.
	cardIndex := 0
	for i := 0; i < 7; i++ {
		for j := 0; j <= i; j++ {
			card := deck[cardIndex]
			if j == i {
				card.FaceUp = true
			}
			g.Tableaus[i].Push(card)
			cardIndex++
		}
	}

	// The rest of the cards go to the stock.
	for i := cardIndex; i < len(deck); i++ {
		g.Stock.Push(deck[i])
	}

	return g
}

// RecycleWaste moves all cards from the waste pile back to the stock pile.
func (g *Game) RecycleWaste() {
	if len(g.Stock.Cards) > 0 {
		return // Can only recycle when stock is empty
	}
	// Reverse the waste pile to put it back into the stock
	for i := len(g.Waste.Cards) - 1; i >= 0; i-- {
		card := g.Waste.Cards[i]
		card.FaceUp = false
		g.Stock.Push(card)
	}
	g.Waste.Cards = nil // Empty the waste pile
}

// DrawCard moves a card from the stock to the waste pile.
func (g *Game) DrawCard() {
	if len(g.Stock.Cards) == 0 {
		return // Or handle recycling waste here, will be added later
	}
	card := g.Stock.Pop()
	card.FaceUp = true
	g.Waste.Push(card)
}

// CheckWinCondition verifies if the game has been won and updates the game state.
func (g *Game) CheckWinCondition() {
	if g.HasWon() {
		g.IsWon = true
	}
}

// GetPile returns a pointer to the pile at the given index.
// 0: Stock, 1: Waste, 2-5: Foundations, 6-12: Tableaus
func (g *Game) GetPile(index int) *Pile {
	switch {
	case index == 0:
		return &g.Stock
	case index == 1:
		return &g.Waste
	case index >= 2 && index <= 5:
		return &g.Foundations[index-2]
	case index >= 6 && index <= 12:
		return &g.Tableaus[index-6]
	default:
		return nil
	}
}

// GetActiveCardIndex returns the appropriate card index to select within a pile.
// For tableau piles, it returns the index of the last face-up card.
// For other piles, it returns 0 (top card).
func (g *Game) GetActiveCardIndex(pileIndex int) int {
	pile := g.GetPile(pileIndex)
	if pile == nil || len(pile.Cards) == 0 {
		return -1
	}

	// For tableau piles, select the last face-up card
	if pileIndex >= 6 && pileIndex <= 12 {
		for i := len(pile.Cards) - 1; i >= 0; i-- {
			if pile.Cards[i].FaceUp {
				return i
			}
		}
		return 0 // If no face-up cards, select the first card (for placing a King)
	}

	// For other piles (Stock, Waste, Foundations), select the top card (index 0)
	return len(pile.Cards) - 1 // Select the actual top card
}

// SetSelection sets the active pile and card for UI selection.
func (g *Game) SetSelection(pileIndex, cardIndex int) {
	g.ActivePile = pileIndex
	g.ActiveCard = cardIndex
}

// ClearSelection clears the active pile and card selection.
func (g *Game) ClearSelection() {
	g.ActivePile = -1
	g.ActiveCard = -1
}

// Move attempts to move a card (or stack of cards) from a source to a destination pile.
// Returns true if the move was successful, false otherwise.
func (g *Game) Move(sourcePileIndex, sourceCardIndex, destPileIndex int) bool {
	sourcePile := g.GetPile(sourcePileIndex)
	destPile := g.GetPile(destPileIndex)

	if sourcePile == nil || destPile == nil {
		return false // Invalid pile indices
	}
	if sourceCardIndex < 0 || sourceCardIndex >= len(sourcePile.Cards) {
		return false // Invalid source card index
	}

	cardsToMove := sourcePile.Cards[sourceCardIndex:]
	if len(cardsToMove) == 0 {
		return false // No cards to move
	}

	// Rule: Cards moved from waste or tableau must be face up.
	for _, card := range cardsToMove {
		if !card.FaceUp {
			return false // Cannot move face-down cards
		}
	}

	// Determine move type based on source and destination pile indices
	// Indices: 0:Stock, 1:Waste, 2-5:Foundations, 6-12:Tableaus
	switch {
	case sourcePileIndex >= TableauPile1 && sourcePileIndex <= TableauPile7 && destPileIndex >= TableauPile1 && destPileIndex <= TableauPile7:
		// Tableau to Tableau
		if len(destPile.Cards) == 0 && cardsToMove[0].Rank == King {
			// Valid move: King to empty tableau
			destPile.Cards = append(destPile.Cards, cardsToMove...)
			sourcePile.Cards = sourcePile.Cards[:sourceCardIndex]
			// Flip the new top card of the source tableau if it's face down
			if len(sourcePile.Cards) > 0 && !sourcePile.Peek().FaceUp {
				sourcePile.Peek().FaceUp = true
			}
			return true
		} else if len(destPile.Cards) > 0 && g.isValidTableauMove(cardsToMove[0], destPile.Peek()) {
			// Valid move: card to non-empty tableau
			destPile.Cards = append(destPile.Cards, cardsToMove...)
			sourcePile.Cards = sourcePile.Cards[:sourceCardIndex]
			// Flip the new top card of the source tableau if it's face down
			if len(sourcePile.Cards) > 0 && !sourcePile.Peek().FaceUp {
				sourcePile.Peek().FaceUp = true
			}
			return true
		}
	case sourcePileIndex == WastePile && destPileIndex >= FoundationPile1 && destPileIndex <= FoundationPile4:
		// Waste to Foundation
		if g.isValidFoundationMove(cardsToMove[0], destPile, destPileIndex-FoundationPile1) {
			destPile.Cards = append(destPile.Cards, cardsToMove...)
			sourcePile.Cards = sourcePile.Cards[:sourceCardIndex]
			return true
		}
	case sourcePileIndex == WastePile && destPileIndex >= TableauPile1 && destPileIndex <= TableauPile7:
		// Waste to Tableau
		if len(destPile.Cards) == 0 && cardsToMove[0].Rank == King {
			// Only Kings can be placed on empty tableaus
			destPile.Cards = append(destPile.Cards, cardsToMove...)
			sourcePile.Cards = sourcePile.Cards[:sourceCardIndex]
			return true
		} else if len(destPile.Cards) > 0 && g.isValidTableauMove(cardsToMove[0], destPile.Peek()) {
			destPile.Cards = append(destPile.Cards, cardsToMove...)
			sourcePile.Cards = sourcePile.Cards[:sourceCardIndex]
			return true
		}
	case sourcePileIndex >= TableauPile1 && sourcePileIndex <= TableauPile7 && destPileIndex >= FoundationPile1 && destPileIndex <= FoundationPile4:
		// Tableau to Foundation
		if len(cardsToMove) == 1 && g.isValidFoundationMove(cardsToMove[0], destPile, destPileIndex-FoundationPile1) {
			destPile.Cards = append(destPile.Cards, cardsToMove...)
			sourcePile.Cards = sourcePile.Cards[:sourceCardIndex]
			// Flip the new top card of the source tableau if it's face down
			if len(sourcePile.Cards) > 0 && !sourcePile.Peek().FaceUp {
				sourcePile.Peek().FaceUp = true
			}
			return true
		}
	case sourcePileIndex >= FoundationPile1 && sourcePileIndex <= FoundationPile4 && destPileIndex >= TableauPile1 && destPileIndex <= TableauPile7:
		// Foundation to Tableau
		if len(cardsToMove) == 1 && (len(destPile.Cards) == 0 || g.isValidTableauMove(cardsToMove[0], destPile.Peek())) {
			destPile.Cards = append(destPile.Cards, cardsToMove...)
			sourcePile.Cards = sourcePile.Cards[:sourceCardIndex]
			return true
		}
	case sourcePileIndex >= FoundationPile1 && sourcePileIndex <= FoundationPile4 && destPileIndex >= FoundationPile1 && destPileIndex <= FoundationPile4:
		// Foundation to Foundation is invalid directly, but cards may move from one to another if empty
		// This is generally not allowed in Klondike, only to build up.
		return false
	}

	return false
}

func (g *Game) isValidTableauMove(movingCard *Card, topDestCard *Card) bool {
	// Must be opposite color and one rank lower
	return movingCard.Suit.Color() != topDestCard.Suit.Color() && movingCard.Rank == topDestCard.Rank-1
}

func (g *Game) isValidFoundationMove(movingCard *Card, destPile *Pile, foundationIndex int) bool {
	// If foundation is empty, only an Ace can be placed
	if len(destPile.Cards) == 0 {
		return movingCard.Rank == Ace // && movingCard.Suit == g.Foundations[foundationIndex].expectedSuit (need to track expected suit)
	}

	// Otherwise, must be same suit and one rank higher
	topDestCard := destPile.Peek()
	return movingCard.Suit == topDestCard.Suit && movingCard.Rank == topDestCard.Rank+1
}

// IsWon checks if the game has been won.
func (g *Game) HasWon() bool {
	for _, pile := range g.Foundations {
		if len(pile.Cards) != 13 {
			return false
		}
	}
	return true
}
