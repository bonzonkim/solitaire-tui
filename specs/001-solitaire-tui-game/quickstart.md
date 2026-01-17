# Quickstart: Solitaire TUI Game

This document provides instructions on how to build and run the Solitaire TUI game.

## Prerequisites

- Go 1.23 or higher

## Build

To build the application, run the following command from the root of the repository:

```bash
go build -o solitaire ./cmd/solitaire
```

## Run

To run the application, execute the following command:

```bash
./solitaire
```

## Controls

### Keyboard (Vim-like)

- **h, j, k, l**: Navigate between piles.
- **Enter**: Select a card or move a selected card.
- **q**: Quit the game.

### Mouse

- **Click**: Select a card or a pile.
- **Click and Drag**: Move a card.
