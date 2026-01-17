# Feature Specification: Realistic ASCII Cards

**Feature Branch**: `006-realistic-cards`
**Created**: 2026-01-17
**Status**: Draft
**Input**: User description: "카드의 디자인을 '실사' 카드처럼 해줘. 실사카드가 어떻게 생긴건지, 뭔지알아? 그걸 Ascii 처럼 그려서 표현하고싶다고" (Design the cards like 'photorealistic' cards. Do you know what a real card looks like? I want to express that by drawing it like ASCII)

## User Scenarios & Testing

### User Story 1 - Artistic Face Cards (Priority: P1)

Users want to see distinct artistic representations for Court cards (King, Queen, Jack) instead of just text labels, making the game feel more authentic and visually interesting.

**Why this priority**: High. Explicit user request for "real card" appearance using ASCII art.

**Independent Test**:
1.  Deal cards until a King, Queen, or Jack appears.
2.  Verify that the center of the card contains an ASCII art representation (e.g., crown, face lines) unique to that rank, rather than just a suit symbol.
3.  Verify the art fits within the card boundaries.

**Acceptance Scenarios**:

1.  **Given** a King card, **When** rendered, **Then** it displays an ASCII "crown" or "beard" motif in the center.
2.  **Given** a Queen card, **When** rendered, **Then** it displays an ASCII "tiara" or feminine profile motif.
3.  **Given** a Jack card, **When** rendered, **Then** it displays a distinct "soldier" or "page" motif.
4.  **Given** different suits, **When** rendering face cards, **Then** the art may vary slightly or share the common motif, but must clearly be the correct rank.

---

### User Story 2 - Realistic Pip Layouts (Priority: P1)

Users want Number cards (Ace-10) to display the correct number of suit symbols (pips) in a layout mimicking real physical cards, rather than a single large symbol.

**Why this priority**: High. "Real" cards are defined by their pip patterns.

**Independent Test**:
1.  Look at a '5' card.
2.  Verify it shows 5 small suit symbols in a pattern (e.g., 2 top, 1 center, 2 bottom) rather than just the number '5' and one big symbol.
3.  Check readability; the card index (Top-Left corner) must still be visible.

**Acceptance Scenarios**:

1.  **Given** a number card (e.g., 7), **When** rendered, **Then** visually count roughly 7 marks/symbols in the center area if space permits, or failing that, a representative pattern that fills the space more "realistically" than a single char.
2.  **Given** limited screen space (7x5 chars), **When** rendering pips, **Then** use dense ASCII patterns or standard pip layouts adapted for the grid (e.g., using `Use 11x7` size if 7x5 is too small for art).
    *   *Self-Correction*: Use updated card dimensions if necessary to fit the art. Current 7x5 (5x3 internal) is extremely small for 10 pips.
    *   **Requirement**: Expand card size to at least 11x7 (9x5 internal) or similar to allow rudimentary pip layouts.

## Requirements

### Functional Requirements

-   **FR-001**: The system MUST support a larger card dimension (e.g., 11 width x 7 height) to accommodate detailed ASCII art.
-   **FR-002**: The system MUST render unique ASCII art patterns for King, Queen, and Jack ranks.
-   **FR-003**: The system MUST render "Pip Patterns" for cards 1-10 where possible, placing multiple suit characters to approximate real card layouts (e.g., corners, center).
-   **FR-004**: The Ace card MUST have a distinct, large center pip (especially Ace of Spades).
-   **FR-005**: All cards MUST retain standard index corners (Rank + Suit) for readability.
-   **FR-006**: The card back pattern MUST be updated to a more intricate ASCII texture (e.g., crosshatch or distinct design).

### Key Entities

-   **CardArtRegistry**: A map or function returning the string/rune grid for the center of a card based on Rank and Suit.
-   **PipLayout**: Definitions for placing suit symbols for numbers 1-10.

### Edge Cases

-   **Small Terminals**: With larger cards (e.g. 11x7), the board width increases. On extremely narrow (e.g. <80 columns) terminals, horizontal scrolling MUST be supported or the layout must behave gracefully.
-   **Non-ASCII Support**: Ensure art uses standard characters or widespread Unicode symbols avoiding rare glyphs that might render as '?' on basic terminals.
-   **Art Clipping**: Ensure ASCII art patterns do not overwrite the corner indices which are crucial for gameplay.

## Success Criteria

### Measurable Outcomes

-   **SC-001**: 100% of Face cards (K, Q, J) display unique internal ASCII art, not just text labels.
-   **SC-002**: Number cards (2-10) display multiple pip symbols (count > 1) mimicking real layouts where space allows.
-   **SC-003**: Card dimensions are increased to at least 11x7 to support this fidelity without breaking the viewport layout (viewport handles scrolling).

## Assumptions
-   The user accepts a larger game board that requires more scrolling in exchange for the requested visual fidelity.
-   ASCII art will be abstract/symbolic given the resolution constraints.
