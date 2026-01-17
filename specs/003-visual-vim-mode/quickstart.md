# Quickstart: Card Visual Update & Vim Mode

**Feature**: 003-visual-vim-mode
**Date**: 2026-01-17

This guide explains how to use the enhanced card visuals and Vim-style controls in Solitaire TUI.

## Running the Game

```bash
cd /path/to/solitaire-tui
go run ./cmd/solitaire
```

## Game Layout

```text
╭───╮╭───╮  ╭───╮╭───╮╭───╮╭───╮
│[?]││5♥ │  │ ♠ ││ ♥ ││ ♦ ││ ♣ │   ← Stock, Waste, Foundations
╰───╯╰───╯  ╰───╯╰───╯╰───╯╰───╯

╭───╮╭───╮╭───╮╭───╮╭───╮╭───╮╭───╮
│4♦ ││░░░││░░░││░░░││░░░││░░░││░░░│  ← Tableau piles
╰───╯│9♥ ││░░░││░░░││░░░││░░░││░░░│
     ╰───╯│Q♠ ││░░░││░░░││░░░││░░░│
          ╰───╯│7♣ ││░░░││░░░││░░░│
               ╰───╯│A♦ ││░░░││░░░│
                    ╰───╯│K♠ ││░░░│
                         ╰───╯│3♥ │
                              ╰───╯
```

## Vim Navigation

### Basic Movement (h/j/k/l)

| Key | Action | Description |
|-----|--------|-------------|
| `h` | ← Left | Move to pile on the left |
| `l` | → Right | Move to pile on the right |
| `k` | ↑ Up | Select card above in tableau |
| `j` | ↓ Down | Select card below in tableau |

**Arrow keys also work** - use whichever you prefer!

### Quick Jump Commands

| Key | Action | Description |
|-----|--------|-------------|
| `gg` | First pile | Jump to Stock pile |
| `G` | Last pile | Jump to Tableau 7 (rightmost) |

### Game Actions

| Key | Action | Description |
|-----|--------|-------------|
| `Enter` / `Space` | Select/Move | Select a card, then select destination |
| `d` or `dd` | Draw | Draw card from Stock (or recycle Waste) |
| `Esc` | Cancel | Cancel current card selection |
| `?` | Help | Show/hide keybinding help overlay |
| `q` | Quit | Exit the game |

## How to Play

### 1. Drawing Cards

Press `d` or `dd` to draw a card from the Stock pile to the Waste pile.

When the Stock is empty, pressing `d` will recycle the Waste back to Stock.

### 2. Moving Cards

1. Navigate to a card using `h`/`j`/`k`/`l` (or arrow keys)
2. Press `Enter` or `Space` to select it
3. Navigate to the destination pile
4. Press `Enter` or `Space` to move the card

**Tips**:
- Move cards to Foundations (top right) to win
- Build Tableau piles in descending order, alternating colors
- Only Kings can be placed on empty Tableau spaces

### 3. Using the Help Overlay

Press `?` at any time to see all available keybindings.

Press `?` again or `Esc` to close the help.

## Visual Indicators

| Visual | Meaning |
|--------|---------|
| **Cyan border** | Currently selected pile/card |
| **Yellow border** | Card being moved (source) |
| `░░░` | Face-down card |
| `[?]` | Stock pile has cards |
| `○` | Stock is empty (press `d` to recycle) |
| `♠ ♥ ♦ ♣` | Empty Foundation placeholders |
| `K` | Empty Tableau (only Kings allowed) |

## Troubleshooting

### Cards look wrong / no colors

- Make sure your terminal supports 256 colors or true color
- Try a different terminal emulator (iTerm2, Alacritty, Kitty)

### Keyboard not responding

- Click on the terminal window to focus it
- Make sure you're not in a different application

### Game too wide for terminal

- Resize your terminal to at least 60 columns wide
- The game adapts to smaller terminals but may look compressed

## Test Scenarios

- [ ] **Navigate all piles**: Use `h`/`l` to cycle through all 13 piles
- [ ] **Stack navigation**: In a tableau with multiple cards, use `j`/`k` to move up/down
- [ ] **Quick jumps**: Press `gg` to go to Stock, `G` to go to last tableau
- [ ] **Draw cycle**: Press `d` until Stock is empty, then `d` again to recycle
- [ ] **Move a card**: Select a card with Enter, navigate to destination, press Enter
- [ ] **Help overlay**: Press `?` to show help, `?` again to hide
- [ ] **Color check**: Verify red cards (♥♦) appear red, black cards (♠♣) appear white
