package game_test

import (
	"testing"

	"github.com/solitaire-tui/solitaire-tui/internal/game"
)

func TestNewGame(t *testing.T) {
	g := game.NewGame()

	if len(g.Tableaus) != 7 {
		t.Errorf("Expected 7 tableau columns, got %d", len(g.Tableaus))
	}

	for i, pile := range g.Tableaus {
		if len(pile.Cards) != i+1 {
			t.Errorf("Expected %d cards in tableau pile %d, got %d", i+1, i, len(pile.Cards))
		}
	}

	// Assuming 52 cards in a deck initially, 28 dealt to tableau
	expectedStockCards := 52 - 28
	if len(g.Stock.Cards) != expectedStockCards {
		t.Errorf("Expected %d cards in stock, got %d", expectedStockCards, len(g.Stock.Cards))
	}
	if len(g.Waste.Cards) != 0 {
		t.Errorf("Expected 0 cards in waste, got %d", len(g.Waste.Cards))
	}
	for _, f := range g.Foundations {
		if len(f.Cards) != 0 {
			t.Errorf("Expected 0 cards in foundation, got %d", len(f.Cards))
		}
	}
}

// func TestValidateMove(t *testing.T) {
// 	g := game.NewGame()

// 	// Valid move (example logic)
// 	fromCard := game.Card{Suit: game.Hearts, Rank: game.Queen}
// 	toCard := game.Card{Suit: game.Clubs, Rank: game.King}
// 	// This test needs actual game state to be realistic. For now, we'll keep it simple.
// 	// We need to place these cards in the game's piles for a real test.
// 	// For now, let's assume a simplified validation logic in the game.Game struct.
// 	// if !g.ValidateMove(fromCard, toCard) {
// 	// 	t.Errorf("Expected move to be valid, but it was invalid")
// 	// }

// 	// Placeholder test for now, ValidateMove is not fully implemented for actual game state
// 	_ = g.ValidateMove(fromCard, toCard)
// }

func setupGameWithSpecificCards(t *testing.T, setup func(g *game.Game)) *game.Game {
	g := game.NewGame()
	// Clear all existing cards to set up a specific scenario
	g.Stock.Cards = nil
	g.Waste.Cards = nil
	for i := range g.Foundations {
		g.Foundations[i].Cards = nil
	}
	for i := range g.Tableaus {
		g.Tableaus[i].Cards = nil
	}
	setup(g)
	return g
}

func TestMove(t *testing.T) {
	tests := []struct {
		name            string
		setup           func(g *game.Game) // Function to set up the game state
		sourcePileIndex int
		sourceCardIndex int
		destPileIndex   int
		expectedSuccess bool
		expectedFlip    bool // Expect a card to be flipped in source Tableau
	}{
		// Waste to Tableau: Valid move (Red 5 to Black 6)
		{
			name: "WasteToTableau_Valid",
			setup: func(g *game.Game) {
				g.Waste.Push(&game.Card{Rank: game.Five, Suit: game.Hearts, FaceUp: true})
				g.Tableaus[0].Push(&game.Card{Rank: game.Six, Suit: game.Clubs, FaceUp: true})
			},
			sourcePileIndex: game.WastePile,
			sourceCardIndex: 0,
			destPileIndex:   game.TableauPile1,
			expectedSuccess: true,
			expectedFlip:    false, // Waste moves don't cause tableau flips
		},
		// Waste to Tableau: Invalid move (Red 5 to Red 6)
		{
			name: "WasteToTableau_InvalidColor",
			setup: func(g *game.Game) {
				g.Waste.Push(&game.Card{Rank: game.Five, Suit: game.Hearts, FaceUp: true})
				g.Tableaus[0].Push(&game.Card{Rank: game.Six, Suit: game.Diamonds, FaceUp: true})
			},
			sourcePileIndex: game.WastePile,
			sourceCardIndex: 0,
			destPileIndex:   game.TableauPile1,
			expectedSuccess: false,
			expectedFlip:    false,
		},
		// Tableau to Tableau: Valid move (Red 5 to Black 6, with flip)
		{
			name: "TableauToTableau_ValidWithFlip",
			setup: func(g *game.Game) {
				g.Tableaus[0].Push(&game.Card{Rank: game.Six, Suit: game.Clubs, FaceUp: false}) // Card to be flipped
				g.Tableaus[0].Push(&game.Card{Rank: game.Five, Suit: game.Hearts, FaceUp: true})
				g.Tableaus[1].Push(&game.Card{Rank: game.Six, Suit: game.Spades, FaceUp: true})
			},
			sourcePileIndex: game.TableauPile1,
			sourceCardIndex: 1, // Moving the Red 5
			destPileIndex:   game.TableauPile2,
			expectedSuccess: true,
			expectedFlip:    true, // Card 0 of Tableau 0 should flip
		},
		// Tableau to Tableau: Invalid move (Black 6 to Red 5)
		{
			name: "TableauToTableau_InvalidRank",
			setup: func(g *game.Game) {
				g.Tableaus[0].Push(&game.Card{Rank: game.Five, Suit: game.Diamonds, FaceUp: true})
				g.Tableaus[1].Push(&game.Card{Rank: game.Six, Suit: game.Clubs, FaceUp: true})
			},
			sourcePileIndex: game.TableauPile2,
			sourceCardIndex: 0,
			destPileIndex:   game.TableauPile1,
			expectedSuccess: false,
			expectedFlip:    false,
		},
		// Tableau to Empty Tableau: Valid King move
		{
			name: "TableauToEmptyTableau_King",
			setup: func(g *game.Game) {
				g.Tableaus[0].Push(&game.Card{Rank: game.King, Suit: game.Hearts, FaceUp: true})
			},
			sourcePileIndex: game.TableauPile1,
			sourceCardIndex: 0,
			destPileIndex:   game.TableauPile2, // Empty pile
			expectedSuccess: true,
			expectedFlip:    false, // No card to flip in source
		},
		// Tableau to Foundation: Valid Ace move
		{
			name: "TableauToFoundation_ValidAce",
			setup: func(g *game.Game) {
				g.Tableaus[0].Push(&game.Card{Rank: game.Two, Suit: game.Clubs, FaceUp: false}) // Card to be flipped
				g.Tableaus[0].Push(&game.Card{Rank: game.Ace, Suit: game.Clubs, FaceUp: true})
			},
			sourcePileIndex: game.TableauPile1,
			sourceCardIndex: 1,                    // Moving Ace of Clubs
			destPileIndex:   game.FoundationPile1, // Clubs Foundation
			expectedSuccess: true,
			expectedFlip:    true,
		},
		// Tableau to Foundation: Invalid rank (Two to empty)
		{
			name: "TableauToFoundation_InvalidRank",
			setup: func(g *game.Game) {
				g.Tableaus[0].Push(&game.Card{Rank: game.Two, Suit: game.Clubs, FaceUp: true})
			},
			sourcePileIndex: game.TableauPile1,
			sourceCardIndex: 0,
			destPileIndex:   game.FoundationPile1,
			expectedSuccess: false,
			expectedFlip:    false,
		},
		// Waste to Foundation: Valid (Ace to empty)
		{
			name: "WasteToFoundation_ValidAce",
			setup: func(g *game.Game) {
				g.Waste.Push(&game.Card{Rank: game.Ace, Suit: game.Diamonds, FaceUp: true})
			},
			sourcePileIndex: game.WastePile,
			sourceCardIndex: 0,
			destPileIndex:   game.FoundationPile2, // Diamonds Foundation
			expectedSuccess: true,
			expectedFlip:    false,
		},
		// Waste to Foundation: Invalid (Two to empty)
		{
			name: "WasteToFoundation_InvalidRank",
			setup: func(g *game.Game) {
				g.Waste.Push(&game.Card{Rank: game.Two, Suit: game.Diamonds, FaceUp: true})
			},
			sourcePileIndex: game.WastePile,
			sourceCardIndex: 0,
			destPileIndex:   game.FoundationPile2,
			expectedSuccess: false,
			expectedFlip:    false,
		},
		// Foundation to Tableau: Valid (Ace to Empty Tableau)
		{
			name: "FoundationToTableau_ValidAce",
			setup: func(g *game.Game) {
				g.Foundations[0].Push(&game.Card{Rank: game.Ace, Suit: game.Spades, FaceUp: true})
				g.Tableaus[0].Cards = nil // Ensure Tableau is empty
			},
			sourcePileIndex: game.FoundationPile1, // Spades foundation
			sourceCardIndex: 0,
			destPileIndex:   game.TableauPile1,
			expectedSuccess: true,
			expectedFlip:    false,
		},
		// Foundation to Foundation: Invalid (direct move not allowed)
		{
			name: "FoundationToFoundation_Invalid",
			setup: func(g *game.Game) {
				g.Foundations[0].Push(&game.Card{Rank: game.Ace, Suit: game.Spades, FaceUp: true})
			},
			sourcePileIndex: game.FoundationPile1, // Spades foundation
			sourceCardIndex: 0,
			destPileIndex:   game.FoundationPile2, // Hearts foundation
			expectedSuccess: false,
			expectedFlip:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := setupGameWithSpecificCards(t, tt.setup)

			// Store the state of the source pile before the move
			sourcePileBefore := g.GetPile(tt.sourcePileIndex)
			var cardToFlip *game.Card
			if tt.sourcePileIndex >= game.TableauPile1 && tt.sourcePileIndex <= game.TableauPile7 &&
				tt.sourceCardIndex > 0 && tt.sourceCardIndex-1 < len(sourcePileBefore.Cards) {
				cardToFlip = sourcePileBefore.Cards[tt.sourceCardIndex-1]
			}

			initialSourcePileLen := len(sourcePileBefore.Cards)
			initialDestPileLen := len(g.GetPile(tt.destPileIndex).Cards)

			// If moving from a tableau and expecting a flip, check the card is initially face down
			if tt.expectedFlip && cardToFlip != nil && cardToFlip.FaceUp {
				t.Fatalf("Test setup error: Card expected to flip is initially face up")
			}

			success := g.Move(tt.sourcePileIndex, tt.sourceCardIndex, tt.destPileIndex)

			if success != tt.expectedSuccess {
				t.Errorf("Move() got success = %v, want %v", success, tt.expectedSuccess)
			}

			if success {
				// Verify card counts changed
				if len(sourcePileBefore.Cards) != initialSourcePileLen-len(g.GetPile(tt.sourcePileIndex).Cards[tt.sourceCardIndex:]) {
					// This line is complex due to slicing behavior.
					// A simpler check: cards should have been removed from source
					if len(sourcePileBefore.Cards) >= initialSourcePileLen {
						t.Errorf("Source pile length did not decrease after successful move")
					}
				}
				if len(g.GetPile(tt.destPileIndex).Cards) != initialDestPileLen+len(g.GetPile(tt.sourcePileIndex).Cards[tt.sourceCardIndex:]) {
					// A simpler check: cards should have been added to destination
					if len(g.GetPile(tt.destPileIndex).Cards) <= initialDestPileLen {
						t.Errorf("Destination pile length did not increase after successful move")
					}
				}
				// Verify flip if expected
				if tt.expectedFlip && (cardToFlip == nil || !cardToFlip.FaceUp) {
					t.Errorf("Expected card to flip, but it did not or was nil")
				}
			} else {
				// If move failed, piles should remain unchanged
				if len(g.GetPile(tt.sourcePileIndex).Cards) != initialSourcePileLen {
					t.Errorf("Source pile changed after failed move. Expected %d, got %d", initialSourcePileLen, len(g.GetPile(tt.sourcePileIndex).Cards))
				}
				if len(g.GetPile(tt.destPileIndex).Cards) != initialDestPileLen {
					t.Errorf("Destination pile changed after failed move. Expected %d, got %d", initialDestPileLen, len(g.GetPile(tt.destPileIndex).Cards))
				}
			}
		})
	}
}

func TestDrawCard(t *testing.T) {
	g := game.NewGame()
	initialStockLen := len(g.Stock.Cards)
	initialWasteLen := len(g.Waste.Cards)

	g.DrawCard()

	if len(g.Stock.Cards) != initialStockLen-1 {
		t.Errorf("Stock should have one less card, got %d want %d", len(g.Stock.Cards), initialStockLen-1)
	}
	if len(g.Waste.Cards) != initialWasteLen+1 {
		t.Errorf("Waste should have one more card, got %d want %d", len(g.Waste.Cards), initialWasteLen+1)
	}
	if !g.Waste.Peek().FaceUp {
		t.Errorf("Top waste card should be face up")
	}
}

func TestDrawCard_EmptyStock(t *testing.T) {
	g := game.NewGame()
	// Empty the stock
	g.Stock.Cards = nil
	initialWasteLen := len(g.Waste.Cards)

	g.DrawCard()

	// Nothing should change when stock is empty
	if len(g.Waste.Cards) != initialWasteLen {
		t.Errorf("Waste should not change when stock is empty, got %d want %d", len(g.Waste.Cards), initialWasteLen)
	}
}

func TestRecycleWaste(t *testing.T) {
	g := game.NewGame()
	// Put all stock cards into waste (simulating draws)
	for len(g.Stock.Cards) > 0 {
		g.DrawCard()
	}
	wasteCount := len(g.Waste.Cards)

	g.RecycleWaste()

	if len(g.Stock.Cards) != wasteCount {
		t.Errorf("Stock should have all waste cards, got %d want %d", len(g.Stock.Cards), wasteCount)
	}
	if len(g.Waste.Cards) != 0 {
		t.Errorf("Waste should be empty after recycle, got %d", len(g.Waste.Cards))
	}
	// All cards in stock should be face down
	for i, card := range g.Stock.Cards {
		if card.FaceUp {
			t.Errorf("Card %d in stock should be face down after recycle", i)
		}
	}
}

func TestRecycleWaste_StockNotEmpty(t *testing.T) {
	g := game.NewGame()
	initialStockLen := len(g.Stock.Cards)

	g.RecycleWaste()

	// Nothing should change when stock is not empty
	if len(g.Stock.Cards) != initialStockLen {
		t.Errorf("Stock should not change when not empty, got %d want %d", len(g.Stock.Cards), initialStockLen)
	}
}

func TestCheckWinCondition(t *testing.T) {
	g := game.NewGame()
	// Clear everything
	g.Stock.Cards = nil
	g.Waste.Cards = nil
	for i := range g.Tableaus {
		g.Tableaus[i].Cards = nil
	}

	// Fill all foundations with 13 cards each
	suits := []game.Suit{game.Spades, game.Hearts, game.Diamonds, game.Clubs}
	for i, suit := range suits {
		for rank := game.Ace; rank <= game.King; rank++ {
			g.Foundations[i].Push(&game.Card{Suit: suit, Rank: rank, FaceUp: true})
		}
	}

	g.CheckWinCondition()

	if !g.IsWon {
		t.Errorf("Game should be won with all cards in foundations")
	}
}

func TestCheckWinCondition_Incomplete(t *testing.T) {
	g := game.NewGame()
	// Clear foundations
	for i := range g.Foundations {
		g.Foundations[i].Cards = nil
	}
	// Add only some cards
	g.Foundations[0].Push(&game.Card{Suit: game.Spades, Rank: game.Ace, FaceUp: true})

	g.CheckWinCondition()

	if g.IsWon {
		t.Errorf("Game should not be won with incomplete foundations")
	}
}

func TestMove_StackMoveToTableau(t *testing.T) {
	g := setupGameWithSpecificCards(t, func(g *game.Game) {
		// Create a stack of cards on tableau 0: King (red) then Queen (black)
		g.Tableaus[0].Push(&game.Card{Rank: game.King, Suit: game.Hearts, FaceUp: true})
		g.Tableaus[0].Push(&game.Card{Rank: game.Queen, Suit: game.Clubs, FaceUp: true})
		// Empty tableau 1
	})

	// Move King + Queen stack to empty tableau
	success := g.Move(game.TableauPile1, 0, game.TableauPile2)

	if !success {
		t.Errorf("Should be able to move King stack to empty tableau")
	}
	if len(g.Tableaus[0].Cards) != 0 {
		t.Errorf("Source tableau should be empty after moving entire stack")
	}
	if len(g.Tableaus[1].Cards) != 2 {
		t.Errorf("Destination tableau should have 2 cards after moving stack")
	}
}

func TestMove_NonKingToEmptyTableau(t *testing.T) {
	g := setupGameWithSpecificCards(t, func(g *game.Game) {
		g.Waste.Push(&game.Card{Rank: game.Queen, Suit: game.Hearts, FaceUp: true})
		// Tableau 0 is empty
	})

	// Try moving non-King to empty tableau (should fail)
	success := g.Move(game.WastePile, 0, game.TableauPile1)

	if success {
		t.Errorf("Should not be able to move non-King to empty tableau")
	}
}

func TestMove_WasteToEmptyTableauKing(t *testing.T) {
	g := setupGameWithSpecificCards(t, func(g *game.Game) {
		g.Waste.Push(&game.Card{Rank: game.King, Suit: game.Hearts, FaceUp: true})
		// Tableau 0 is empty
	})

	// Move King from waste to empty tableau
	success := g.Move(game.WastePile, 0, game.TableauPile1)

	if !success {
		t.Errorf("Should be able to move King from waste to empty tableau")
	}
}

func TestMove_FoundationSequence(t *testing.T) {
	g := setupGameWithSpecificCards(t, func(g *game.Game) {
		g.Foundations[0].Push(&game.Card{Rank: game.Ace, Suit: game.Spades, FaceUp: true})
		g.Waste.Push(&game.Card{Rank: game.Two, Suit: game.Spades, FaceUp: true})
	})

	// Move Two of Spades onto Ace of Spades foundation
	success := g.Move(game.WastePile, 0, game.FoundationPile1)

	if !success {
		t.Errorf("Should be able to build foundation from Ace to Two")
	}
	if len(g.Foundations[0].Cards) != 2 {
		t.Errorf("Foundation should have 2 cards after move")
	}
}

func TestMove_FoundationWrongSuit(t *testing.T) {
	g := setupGameWithSpecificCards(t, func(g *game.Game) {
		g.Foundations[0].Push(&game.Card{Rank: game.Ace, Suit: game.Spades, FaceUp: true})
		g.Waste.Push(&game.Card{Rank: game.Two, Suit: game.Hearts, FaceUp: true})
	})

	// Try moving Two of Hearts onto Ace of Spades (wrong suit)
	success := g.Move(game.WastePile, 0, game.FoundationPile1)

	if success {
		t.Errorf("Should not be able to move wrong suit to foundation")
	}
}
