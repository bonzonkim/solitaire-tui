# Implementation Plan: Card Visual Update & Vim Mode Support

**Branch**: `003-visual-vim-mode` | **Date**: 2026-01-17 | **Spec**: [spec.md](./spec.md)
**Input**: Feature specification from `/specs/003-visual-vim-mode/spec.md`

## Summary

This feature enhances the Solitaire TUI with two main improvements:
1. **Visual Update**: Transform card display to look more like real playing cards with proper borders, overlapping stacks, and clear suit/rank visibility.
2. **Vim Mode Support**: Implement full Vim-style navigation (h/j/k/l) and command shortcuts (dd, gg, G, ?) for power users.

The implementation builds on the existing Bubble Tea + Lip Gloss architecture, modifying primarily the UI layer (view.go, update.go) with minimal changes to the game logic layer.

## Technical Context

**Language/Version**: Go 1.23+
**Primary Dependencies**: Bubble Tea (`bubbletea`), Lip Gloss (`lipgloss`)
**Storage**: N/A (in-memory game state only)
**Testing**: Go standard `testing` package (Table-Driven)
**Target Platform**: Terminal (TUI) - 256-color and true-color support
**Project Type**: Single project
**Performance Goals**: Input response < 16ms for 60fps feel
**Constraints**: Must adhere to The Elm Architecture principles (strict separation of logic/UI)
**Scale/Scope**: TUI Application

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- [X] Does the plan enforce strict separation between `internal/game` (logic) and `internal/ui` (view)?
  - ✅ Card visual changes are purely in `internal/ui/view.go` (styling only)
  - ✅ Vim keybindings are handled in `internal/ui/update.go` (input handling only)
  - ✅ Command logic (dd draw, gg jump) calls existing `internal/game` methods
- [X] Is all game state mutation handled exclusively by the logic layer?
  - ✅ `DrawCard()`, `RecycleWaste()`, `SetSelection()` are already in `internal/game`
  - ✅ New commands will use existing game methods
- [X] Does the plan include tests for all core game logic, including edge cases?
  - ✅ Existing tests cover game logic
  - ✅ New keybinding tests will be integration-style UI tests if needed
- [X] Does the implementation avoid `panic()` for control flow?
  - ✅ All errors handled via return values
- [X] Does the directory structure align with `cmd/`, `internal/game/`, and `internal/ui/`?
  - ✅ No new directories needed

**Result**: All constitution checks PASS. Proceeding to Phase 0.

## Project Structure

### Documentation (this feature)

```text
specs/003-visual-vim-mode/
├── plan.md              # This file
├── research.md          # Phase 0 output
├── data-model.md        # Phase 1 output (card visual model)
├── quickstart.md        # Phase 1 output (usage guide)
└── tasks.md             # Phase 2 output (created by /speckit.tasks)
```

### Source Code (repository root)
```text
cmd/
└── solitaire/
    └── main.go          # Entry point (no changes needed)
internal/
├── game/
│   ├── card.go          # Card Rank/Suit String() methods (minor update for 10→"10")
│   ├── deck.go          # Deck creation (no changes)
│   ├── game.go          # Core game logic (minor: add undo stack if implementing `u`)
│   └── piles.go         # Pile operations (no changes)
└── ui/
    ├── model.go         # UI state (add: vim command buffer, help overlay state)
    ├── update.go        # Input handling (add: vim commands, ? help toggle)
    ├── view.go          # Card rendering (MAJOR: new card visuals, overlapping)
    └── styles/
        └── styles.go    # Lip Gloss styles (MAJOR: new card styles)
tests/
└── game/
    └── game_test.go     # Existing tests (no changes needed)
```

**Structure Decision**: All changes fit within existing architecture. No new packages or directories required beyond what exists.

## Implementation Strategy

### Phase 1: Card Visual Overhaul (P1 - FR-001 to FR-005, FR-008, FR-009)

1. **Update styles.go**: Define new card styles with proper borders
   - Multi-line card template (top border, rank row, middle, suit row, bottom border)
   - Colors: Red for Hearts/Diamonds, White for Spades/Clubs
   - Selection highlight (cyan border)
   - Face-down pattern (░░░ or ╬╬╬)

2. **Update view.go**: Render cards with new visual structure
   - Card template: 5x3 character box minimum
   - Overlapping display for tableau (show 1 line per hidden card)
   - Empty pile placeholders (suit symbol for foundations, K for tableaus)

3. **Update card.go**: Fix "T" to display "10"
   - Special handling for 10 to fit in card width

### Phase 2: Vim Navigation Enhancement (P1 - FR-006, FR-007)

1. Already implemented: h/j/k/l keys work alongside arrow keys
2. Verify and test all directions

### Phase 3: Vim Commands (P2 - FR-011 to FR-014)

1. **Command buffer** in model.go for multi-key sequences (dd, gg)
2. **Update.go enhancements**:
   - `dd` → call DrawCard() or RecycleWaste()
   - `gg` → SetSelection(StockPile, ...)
   - `G` → SetSelection(TableauPile7, ...)
   - `?` → toggle help overlay
3. **Help overlay** in view.go

### Phase 4: Polish (P2)

1. Compact mode for small terminals
2. Clear status messages
3. Proper key debouncing (already handled by Bubble Tea)

## Complexity Tracking

> No violations. All implementation fits within existing architecture.

| Violation | Why Needed | Simpler Alternative Rejected Because |
|-----------|------------|-------------------------------------|
| N/A | N/A | N/A |

## Risk Assessment

| Risk | Mitigation |
|------|------------|
| Card width exceeds terminal | Use compact mode detection, fallback to 1-line cards |
| Multi-key commands feel laggy | Short timeout (200ms) for command sequences |
| Color not supported | Detect color capability, use bold/reverse as fallback |
