# Data Model: Card Rendering

**Date**: 2026-01-18  
**Feature**: 007-card-rendering

## Overview

No new data structures are required. This feature purely modifies the **presentation layer** using existing game models.

## Existing Entities (No Changes)

### Card (`internal/game/card.go`)
```go
type Card struct {
    Suit   Suit  // Spades, Hearts, Diamonds, Clubs
    Rank   Rank  // Ace, 2-10, Jack, Queen, King
    FaceUp bool  // True = face up, False = face down
}
```

### Suit (`internal/game/card.go`)
```go
type Suit int  // Spades=0, Hearts=1, Diamonds=2, Clubs=3
func (s Suit) String() string  // Returns ♠, ♥, ♦, ♣
func (s Suit) Color() string   // Returns "Red" or "Black"
```

### Rank (`internal/game/card.go`)
```go
type Rank int  // Ace=1, Two=2, ..., King=13
func (r Rank) String() string  // Returns A, 2-10, J, Q, K
```

## New Template Structures (View Layer Only)

These are string templates used in `internal/ui/view.go`, not data models.

### Face-Up Card Template
```
Line 1: ┌─────────┐
Line 2: │ {RANK_L}      │
Line 3: │         │
Line 4: │    {SUIT}    │
Line 5: │         │
Line 6: │      {RANK_R} │
Line 7: └─────────┘
```

### Face-Down Card Template
```
Line 1: ┌─────────┐
Line 2: │░░░░░░░░░│
Line 3: │░░░░░░░░░│
Line 4: │░░░░░░░░░│
Line 5: │░░░░░░░░░│
Line 6: │░░░░░░░░░│
Line 7: └─────────┘
```

## Constraints

- No changes to `internal/game/` package (constitution requirement)
- All rendering logic stays in `internal/ui/` package
