# Data Model: Solitaire TUI Game

This document defines the data structures for the Solitaire TUI game.

## Card

Represents a single playing card.

- **Suit**: `string` (Clubs, Diamonds, Hearts, Spades)
- **Rank**: `string` (A, 2, 3, 4, 5, 6, 7, 8, 9, 10, J, Q, K)
- **FaceUp**: `boolean`

## Deck

Represents a collection of 52 cards.

- **Cards**: `[]Card`

**Validation Rules**:

- A deck must contain exactly 52 unique cards.

## Tableau

Represents the seven columns of cards in the main playing area.

- **Columns**: `[][]Card` (7 columns)

**Validation Rules**:

- Cards can only be placed on a card of the opposite color and one rank higher.
- An empty column can only be filled with a King.

## Stock

Represents the pile of cards that have not yet been dealt to the waste pile.

- **Cards**: `[]Card`

## Waste

Represents the pile of cards from the stock that can be played.

- **Cards**: `[]Card`

## Foundation

Represents the four piles where cards are moved to win the game.

- **Piles**: `[][]Card` (4 piles)

**Validation Rules**:

- Each pile must start with an Ace.
- Cards must be placed on a card of the same suit and one rank higher.
