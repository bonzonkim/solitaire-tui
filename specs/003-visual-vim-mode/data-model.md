# Data Model: Card Visual Update & Vim Mode

**Feature**: 003-visual-vim-mode
**Date**: 2026-01-17

This document describes the data structures and visual models for the card display enhancement and Vim mode support.

## 1. Card Visual Model

### Card Display Template

A card is rendered as a bordered box with the following structure:

```text
╭───╮
│{R}{S}│
╰───╯
```

Where:
- `{R}` = Rank (A, 2-9, 10, J, Q, K)
- `{S}` = Suit symbol (♠, ♥, ♦, ♣)

### Card States

| State | Visual | Description |
|-------|--------|-------------|
| Face-up (Red) | `│5♥│` | Red foreground color |
| Face-up (Black) | `│K♠│` | White/light foreground color |
| Face-down | `│░░░│` | Patterned background |
| Selected | Cyan border | `BorderForeground("#00FFFF")` |
| Source (moving) | Yellow border | `BorderForeground("#FFFF00")` |

### Card Dimensions

| Property | Value | Notes |
|----------|-------|-------|
| Width | 5 characters | Including borders |
| Height | 3 lines | Full card with borders |
| Content Width | 3 characters | Inside borders |
| Overlap Height | 1 line | For stacked display |

## 2. Pile Visual Models

### Stock Pile

```text
╭───╮
│[?]│   ← Cards remaining
╰───╯

╭───╮
│ ○ │   ← Empty (can recycle)
╰───╯
```

### Waste Pile

```text
╭───╮
│5♥ │   ← Top card visible
╰───╯

(empty: no display or dim placeholder)
```

### Foundation Piles

```text
╭───╮ ╭───╮ ╭───╮ ╭───╮
│ ♠ │ │ ♥ │ │ ♦ │ │ ♣ │   ← Empty (suit placeholders)
╰───╯ ╰───╯ ╰───╯ ╰───╯

╭───╮
│K♠ │   ← With cards (top card shown)
╰───╯
```

### Tableau Piles

**Stacked Display** (vertical overlap):

```text
╭───╮
│░░░│   ← Face-down cards (1 line each)
├───┤
│░░░│
├───┤
│5♥ │   ← First face-up card
╰───╯
│3♠ │   ← Overlapping card
╰───╯
│A♦ │   ← Top card (fully visible)
╰───╯
```

**Simplified Stacked Display** (for compact mode):

```text
 ░░░
 ░░░
 5♥
 3♠
 A♦   ← Top card
```

### Empty Tableau

```text
╭───╮
│ K │   ← Indicates only Kings allowed
╰───╯
```

## 3. UI State Model

### Extended Model Fields

```go
type model struct {
    // Existing fields
    game            *game.Game
    width           int
    height          int
    sourcePileIndex int
    sourceCardIndex int
    showInvalidMove bool
    
    // New fields for Vim mode
    lastKey         string        // For multi-key commands (g, d)
    lastKeyTime     time.Time     // Timeout for command sequences
    showHelp        bool          // Toggle help overlay
}
```

### Vim Command State Machine

```text
                    ┌──────────┐
    any key         │  IDLE    │
    ──────────────► │(default) │
                    └────┬─────┘
                         │
         'g' pressed     │    'd' pressed
         ┌───────────────┼───────────────┐
         ▼               │               ▼
    ┌─────────┐          │         ┌─────────┐
    │ WAIT_G  │          │         │ WAIT_D  │
    │ (300ms) │          │         │ (300ms) │
    └────┬────┘          │         └────┬────┘
         │               │               │
    'g' = gg action      │          'd' = dd action
    other = cancel       │          other = cancel
         │               │               │
         ▼               ▼               ▼
    ┌──────────┐    ┌──────────┐    ┌──────────┐
    │ Jump to  │    │  IDLE    │    │  Draw    │
    │  Stock   │    │          │    │  Card    │
    └──────────┘    └──────────┘    └──────────┘
```

## 4. Keybinding Model

### Navigation Keys (Always Active)

| Key | Action | Vim Equivalent |
|-----|--------|----------------|
| h / ← | Move selection left | Cursor left |
| j / ↓ | Move selection down | Cursor down |
| k / ↑ | Move selection up | Cursor up |
| l / → | Move selection right | Cursor right |

### Command Keys

| Key(s) | Action | Vim Equivalent |
|--------|--------|----------------|
| gg | Jump to Stock pile | Go to first line |
| G | Jump to last Tableau | Go to last line |
| dd | Draw/cycle stock | Delete line (repurposed) |
| u | Undo last move | Undo |
| ? | Toggle help overlay | Help |

### Action Keys (Unchanged)

| Key | Action |
|-----|--------|
| Enter / Space | Select or move card |
| Esc | Cancel selection |
| d | Draw card (single key shortcut) |
| q | Quit game |

## 5. Color Palette

### Card Colors

| Element | Light Theme | Dark Theme |
|---------|-------------|------------|
| Red Suit (♥♦) | #DC143C | #FF6B6B |
| Black Suit (♠♣) | #000000 | #FFFFFF |
| Card Background | #FFFFFF | #1E1E1E |
| Face-down Pattern | #1A365D | #1A365D |
| Border (normal) | #444444 | #444444 |
| Border (selected) | #00FFFF | #00FFFF |
| Border (source) | #FFFF00 | #FFFF00 |

### UI Colors

| Element | Color |
|---------|-------|
| Title Background | #25A065 |
| Error Text | #FF5555 |
| Help Text | #888888 |
| Help Overlay BG | #1E1E1E |

## 6. Help Overlay Model

### Overlay Dimensions

- Width: 50 characters (or terminal width - 10)
- Height: 20 lines (or terminal height - 5)
- Position: Centered

### Content Structure

```text
┌─────────────────────────────────────────────────┐
│              ♠ Solitaire Help ♥                 │
├─────────────────────────────────────────────────┤
│                                                 │
│  NAVIGATION                                     │
│  ───────────                                    │
│  h / ←      Move left                          │
│  j / ↓      Move down (in pile)                │
│  k / ↑      Move up (in pile)                  │
│  l / →      Move right                         │
│  gg         Jump to Stock                      │
│  G          Jump to last Tableau               │
│                                                 │
│  ACTIONS                                        │
│  ───────                                        │
│  Enter      Select / Move card                 │
│  d / dd     Draw from Stock                    │
│  Esc        Cancel selection                   │
│  q          Quit game                          │
│                                                 │
│  Press ? or Esc to close                       │
└─────────────────────────────────────────────────┘
```
