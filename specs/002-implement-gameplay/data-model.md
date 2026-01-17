# Data Model: Solitaire Gameplay

- **Author**: AI Assistant
- **Created**: 2026-01-17

This document outlines the core data structures used to represent the state of the Klondike Solitaire game. These structs will primarily reside in the `internal/game` package.

## 1. Core Enumerations

These types provide strict definitions for card properties.

### Suit

The suit of a card.

- `Spades`
- `Hearts`
- `Clubs`
- `Diamonds`

### Rank

The rank of a card.

- `Ace`
- `Two`
- `Three`
- `Four`
- `Five`
- `Six`
- `Seven`
- `Eight`
- `Nine`
- `Ten`
- `Jack`
- `Queen`
- `King`

## 2. Main Entities

### Card

Represents a single playing card.

| Field | Type | Description |
|---|---|---|
| `Suit` | `Suit` | The suit of the card (e.g., Hearts). |
| `Rank` | `Rank` | The rank of the card (e.g., King). |
| `FaceUp`| `bool` | `true` if the card's face is visible, `false` otherwise. |

### Pile

A slice of `*Card` pointers, representing any collection of cards on the board.

| Field | Type | Description |
|---|---|---|
| `Cards`| `[]*Card` | An ordered slice representing the stack of cards. |

## 3. Game State Models

### Game

This is the top-level struct that holds the entire state of the game board and the UI selection state.

| Field | Type | Description |
|---|---|---|
| `Stock` | `Pile` | The pile of face-down cards to be drawn. |
| `Waste` | `Pile` | The pile of face-up cards drawn from the stock. |
| `Foundations`| `[4]Pile`| The four piles where suits are built up from Ace to King. |
| `Tableaus` | `[7]Pile`| The seven main playing piles. |
| `IsWon` | `bool` | Flag that is set to `true` when the win condition is met. |
| `ActivePile` | `int` | Index or enum representing the currently selected pile group (e.g., Stock, Waste, Foundation 0-3, Tableau 0-6). |
| `ActiveCard` | `int` | Index of the currently selected card within the `ActivePile`. A value of `-1` may indicate no card is selected. |
