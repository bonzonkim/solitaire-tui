# Research: Card Rendering Implementation

**Date**: 2026-01-18  
**Feature**: 007-card-rendering

## Key Decisions

### Decision 1: Border Rendering Approach
**Decision**: Use manual string construction with Unicode Box Drawing characters instead of Lipgloss's border system.

**Rationale**: 
- The user's specification requires exact Unicode characters (`┌`, `┐`, `└`, `┘`, `─`, `│`)
- Lipgloss's `RoundedBorder()` uses different characters that don't match the spec
- Manual string construction provides exact control over character placement

**Alternatives Considered**:
- Custom Lipgloss border style → Rejected because Lipgloss borders wrap content, making exact 11x7 sizing harder to guarantee
- Continue using Lipgloss RoundedBorder → Rejected because it doesn't match user specification

### Decision 2: Rank Padding Strategy
**Decision**: Use `fmt.Sprintf` with dynamic width specifiers to handle 1-character and 2-character ranks.

**Rationale**:
- Only "10" has 2 characters; all other ranks are 1 character
- Left-aligned top: `"%-2s"` (A becomes "A ", 10 becomes "10")
- Right-aligned bottom: `"%2s"` (A becomes " A", 10 becomes "10")
- This ensures border alignment is never broken

**Alternatives Considered**:
- Fixed-width rank constants → More code, same outcome
- String manipulation with len() → Less readable than format strings

### Decision 3: Card Content Structure
**Decision**: Build 7 complete lines as a string array, then join with newlines.

**Rationale**:
- Matches the 7-line specification exactly
- Easy to verify each line's character count
- Separates border logic from content logic

## Codebase Analysis

### Existing Components (Reusable)
- `internal/game/card.go` - `Card`, `Suit`, `Rank` types already exist with `String()` methods
- `internal/ui/styles/styles.go` - Already has `CardWidth = 11`, `CardHeight = 7`
- Suit symbols (♠, ♥, ♦, ♣) already returned by `Suit.String()`

### Components to Modify
- `internal/ui/view.go` - Replace `renderCard()` function's string building logic
- No changes needed to game logic layer (adheres to constitution)

## Unicode Characters Reference

| Character | Name | Unicode |
|-----------|------|---------|
| ┌ | Box Drawings Light Down and Right | U+250C |
| ┐ | Box Drawings Light Down and Left | U+2510 |
| └ | Box Drawings Light Up and Right | U+2514 |
| ┘ | Box Drawings Light Up and Left | U+2518 |
| ─ | Box Drawings Light Horizontal | U+2500 |
| │ | Box Drawings Light Vertical | U+2502 |
| ░ | Light Shade | U+2591 |
