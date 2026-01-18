# Implementation Plan: Card Visual Rendering

**Branch**: `007-card-rendering` | **Date**: 2026-01-18 | **Spec**: [spec.md](file:///Users/bumgukang/bonzonkim/github.com/solitaire-tui/specs/007-card-rendering/spec.md)  
**Input**: Feature specification for 11×7 ASCII card design with Unicode Box Drawing borders

## Summary

Modify the card rendering in `internal/ui/view.go` to use manually constructed Unicode Box Drawing character borders (┌┐└┘─│) instead of Lipgloss's built-in border system. This ensures exact 11×7 character dimensions and consistent display across terminals.

The current implementation uses Lipgloss's `RoundedBorder()` which wraps content; the new implementation will build each card line explicitly.

## Technical Context

**Language/Version**: Go 1.23+  
**Primary Dependencies**: Bubble Tea, Lip Gloss  
**Storage**: N/A  
**Testing**: Go standard `testing` package (Table-Driven)  
**Target Platform**: Terminal (TUI)  
**Project Type**: Single project  
**Constraints**: Must adhere to The Elm Architecture principles

## Constitution Check

*GATE: Must pass before Phase 0 research. Re-check after Phase 1 design.*

- [x] Does the plan enforce strict separation between `internal/game` (logic) and `internal/ui` (view)?
- [x] Is all game state mutation handled exclusively by the logic layer?
- [x] Does the plan include tests for all core game logic, including edge cases?
- [x] Does the implementation avoid `panic()` for control flow?
- [x] Does the directory structure align with `cmd/`, `internal/game/`, and `internal/ui/`?

> **Note**: This feature modifies only the UI layer (`internal/ui`). No changes to game logic (`internal/game`).

---

## Proposed Changes

### UI Layer

#### [MODIFY] [view.go](file:///Users/bumgukang/bonzonkim/github.com/solitaire-tui/internal/ui/view.go)

**Replace the `renderCard()` function** (lines 222-373) to use manual Unicode box drawing:

1. **Face-Up Card Rendering**:
   - Line 1: `┌─────────┐` (11 chars)
   - Line 2: `│ {RANK_L}      │` where `{RANK_L}` is left-aligned (1-2 chars)
   - Line 3: `│         │` (empty middle)
   - Line 4: `│    {SUIT}    │` with centered suit symbol
   - Line 5: `│         │` (empty middle)
   - Line 6: `│      {RANK_R} │` where `{RANK_R}` is right-aligned (1-2 chars)
   - Line 7: `└─────────┘` (11 chars)

2. **Face-Down Card Rendering**:
   - Line 1: `┌─────────┐`
   - Lines 2-6: `│░░░░░░░░░│` (Light Shade U+2591)
   - Line 7: `└─────────┘`

3. **Overlap Mode** (for stacked tableau cards):
   - Show only top 2 lines of card
   - Face-up overlap: `┌─────────┐\n│ {RANK_L}      │`
   - Face-down overlap: `┌─────────┐\n│░░░░░░░░░│`

4. **Styling**:
   - Remove `BorderStyle()` from Lipgloss styles (borders now embedded in content)
   - Keep color styling for red/black suits, selected states

#### [MODIFY] [styles.go](file:///Users/bumgukang/bonzonkim/github.com/solitaire-tui/internal/ui/styles/styles.go)

**Add new card rendering helper constants**:
```go
const (
    BorderTop    = "┌─────────┐"
    BorderBottom = "└─────────┘"
    BorderVert   = "│"
    FaceDownFill = "░░░░░░░░░"
    InnerWidth   = 9  // Characters between vertical borders
)
```

**Update card styles** to remove built-in borders (they're now in the content string).

---

## Verification Plan

### Automated Tests

**Existing tests** (no modifications needed):
```bash
go test ./tests/game/...
```
These test game logic only. Card rendering is visual, not covered by existing unit tests.

**Build verification**:
```bash
go build -o solitaire ./cmd/solitaire
```

### Manual Verification

The user should run the game and verify visually:

1. **Launch the game**:
   ```bash
   ./solitaire
   ```

2. **Check Face-Up Cards**:
   - Pick any visible face-up card
   - Verify it has Unicode box borders (┌┐└┘─│)
   - Verify rank appears in top-left and bottom-right corners
   - Verify suit symbol is centered on line 4
   - Count characters: should be exactly 11 wide × 7 tall

3. **Check "10" Rank**:
   - Draw cards until you see a 10
   - Verify "10" doesn't break border alignment
   - Top-left should show "10" followed by appropriate spacing

4. **Check Face-Down Cards**:
   - Look at any face-down card in tableaus
   - Verify it shows ░ pattern fill (5 lines of ░░░░░░░░░)
   - Verify same box border style as face-up cards

5. **Check Overlapping Cards**:
   - Look at stacked tableaus
   - Verify overlapped cards show only top 2 lines
   - Verify full card is visible at bottom of each stack
