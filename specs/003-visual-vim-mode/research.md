# Research: Card Visual Update & Vim Mode Support

**Feature**: 003-visual-vim-mode
**Date**: 2026-01-17
**Status**: Complete

## 1. Card Visual Design Patterns

### Decision: ASCII Art Card Template

**Chosen Approach**: 5-width x 3-height character cards with Unicode borders

```text
╭───╮
│A♠ │   ← Face-up card (rank + suit)
╰───╯

╭───╮
│░░░│   ← Face-down card
╰───╯

╭───╮
│ ♠ │   ← Empty foundation (suit placeholder)
╰───╯
```

**Rationale**:
- Compact enough for 80-column terminals (7 piles × 6 chars = 42 chars + spacing)
- Lipgloss `RoundedBorder()` provides clean visual
- Single-line content allows for overlapping stacks

**Alternatives Considered**:
1. **Larger cards (7x5)**: Rejected - too wide for 80-column terminals with 7 tableau piles
2. **No borders**: Rejected - cards blend together, poor visual separation
3. **Double-line borders**: Rejected - too heavy visually

### Decision: Overlapping Tableau Display

**Chosen Approach**: Face-down cards show 1 line, face-up show full card

```text
Tableau Column:
╭───╮
│░░░│   ← face-down (compressed to 1 line when stacked)
│░░░│   ← face-down
│5♥ │   ← face-up (visible)
╰───╯
│4♠ │   ← face-up (only middle line shown)
╰───╯
│3♥ │   ← face-up (top card - full display)
╰───╯
```

**Simplified Approach for TUI**:
```text
  ───      ← face-down indicator line
  ───      ← face-down indicator line
╭───╮
│5♥ │      ← first face-up card
╰┬──╯
 │3♠│      ← overlapping card (abbreviated)
 ╰──╯
```

**Final Simplified**:
```text
░░░        ← face-down
░░░        ← face-down
5♥         ← face-up
3♠         ← face-up (bottom shows top portion)
```

**Rationale**: Keep it simple. Show rank+suit on each line. Use borders only on selection.

## 2. Vim Keybinding Patterns

### Decision: Partial Vim Mode (Not Modal)

**Chosen Approach**: Always-active Vim keys, no mode switching (INSERT/NORMAL)

| Key | Action | Notes |
|-----|--------|-------|
| h | Move left | Same as ← |
| j | Move down | Same as ↓ |
| k | Move up | Same as ↑ |
| l | Move right | Same as → |
| gg | Jump to Stock | First pile |
| G | Jump to Tableau 7 | Last pile |
| dd | Draw card | Same as 'd' key |
| u | Undo | Future: requires undo stack |
| ? | Toggle help | Show/hide overlay |

**Rationale**:
- Modal editing (pressing `i` to enter input mode) makes no sense for a card game
- Vim navigation keys are intuitive for the target audience
- Multi-key commands (gg, dd) add power-user convenience

**Alternatives Considered**:
1. **Full modal**: Rejected - card game has no "insert" mode concept
2. **Single-key only**: Rejected - gg/G jumps are very useful
3. **`:` command mode**: Rejected - overkill for a card game

### Decision: Command Buffer Implementation

**Chosen Approach**: Track last keystroke with timeout

```go
type model struct {
    // ...existing fields...
    lastKey      string    // Previous keystroke
    lastKeyTime  time.Time // When it was pressed
}

const commandTimeout = 300 * time.Millisecond

func (m model) handleVimCommand(key string) (model, tea.Cmd) {
    now := time.Now()
    
    // Check for multi-key command
    if now.Sub(m.lastKeyTime) < commandTimeout {
        switch m.lastKey + key {
        case "gg":
            // Jump to Stock
        case "dd":
            // Draw card
        }
    }
    
    // Handle single-key commands
    switch key {
    case "G":
        // Jump to last tableau
    case "?":
        // Toggle help
    }
    
    m.lastKey = key
    m.lastKeyTime = now
    return m, nil
}
```

**Rationale**: Simple state machine. No complex parsing needed.

## 3. Color and Terminal Compatibility

### Decision: True Color with Fallback

**Chosen Approach**: Use Lipgloss adaptive colors

```go
var (
    RedColor   = lipgloss.AdaptiveColor{Light: "#DC143C", Dark: "#FF6B6B"}
    WhiteColor = lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FFFFFF"}
)
```

**Rationale**:
- Lipgloss handles terminal capability detection
- AdaptiveColor provides light/dark theme support
- No manual TERM checking needed

## 4. Help Overlay Design

### Decision: Simple Text Overlay

**Chosen Approach**: Full-screen overlay activated by `?`

```text
╭─────────────── Solitaire Help ───────────────╮
│                                               │
│  Navigation:                                  │
│    h/←     Move left                         │
│    l/→     Move right                        │
│    k/↑     Move up (in tableau)              │
│    j/↓     Move down (in tableau)            │
│    gg      Jump to Stock                     │
│    G       Jump to last Tableau              │
│                                               │
│  Actions:                                     │
│    Enter   Select/Move card                  │
│    Space   Select/Move card                  │
│    d/dd    Draw from Stock                   │
│    Esc     Cancel selection                  │
│    u       Undo (coming soon)                │
│    q       Quit game                         │
│                                               │
│  Press ? or Esc to close this help           │
╰───────────────────────────────────────────────╯
```

**Rationale**: Clear, comprehensive, follows Vim's `:help` philosophy.

## 5. Existing Code Analysis

### Current Implementation Status

| Feature | Status | Location |
|---------|--------|----------|
| h/j/k/l navigation | ✅ Implemented | update.go:44-69 |
| Arrow key navigation | ✅ Implemented | update.go:44-69 |
| Card rendering | ⚠️ Basic | view.go (needs visual upgrade) |
| Selection highlight | ✅ Implemented | view.go:renderCard() |
| Multi-key commands | ❌ Not implemented | Needs addition |
| Help overlay | ❌ Not implemented | Needs addition |
| Undo functionality | ❌ Not implemented | Future scope |

### Files to Modify

1. **internal/ui/styles/styles.go** - Card visual styles (MAJOR)
2. **internal/ui/view.go** - Card rendering (MAJOR)
3. **internal/ui/model.go** - Add lastKey, showHelp fields (MINOR)
4. **internal/ui/update.go** - Add vim commands (MODERATE)
5. **internal/game/card.go** - Optional: "10" instead of "T" (MINOR)

## Summary

All technical decisions resolved. No NEEDS CLARIFICATION items remain. Ready for Phase 1 design artifacts.
