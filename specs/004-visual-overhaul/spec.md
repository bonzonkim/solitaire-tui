# Feature Specification: Visual Overhaul

**Feature Branch**: `004-visual-overhaul`
**Created**: 2026-01-17
**Status**: Draft
**Input**: User description: "작은카드로하지마 그리고 더 이쁘게 해야돼 난 실제카드같은 그림을 원한다고" (Don't use small cards, make it prettier, I want pictures like real cards)

## User Scenarios & Testing

### User Story 1 - Immersive Large Card Visuals (Priority: P1)

Users want to play Solitaire with cards that look realistic and are easy to read, avoiding the cramped "mini-card" look.

**Why this priority**: High. The user explicitly rejected small cards and requested realistic visuals. This is the core visual requirement.

**Independent Test**: Can be tested by launching the game and verifying card dimensions are substantial (e.g., 7x5 characters) and resemble playing cards with proper ranks and suits on a green felt background.

**Acceptance Scenarios**:

1. **Given** a new game starts, **When** the board renders, **Then** all cards are displayed with large dimensions (at least 7 characters wide, 5 lines high).
2. **Given** cards are stacked in a tableau, **When** rendered, **Then** they overlap vertically but remain readable, creating a realistic "cascade" effect.
3. **Given** the game board, **When** rendered, **Then** the background is a consistent "table" color (e.g., green) filling the view.

---

### User Story 2 - Dynamic Full-Screen Viewport (Priority: P1)

Users want the game to utilize the entire terminal window without broken scrolling or layout issues, regardless of card playing surface size.

**Why this priority**: High. Large cards require more space than a standard terminal screen might offer. A viewport is essential to prevent the "scrolling off screen" bug while allowing the desired visual fidelity.

**Independent Test**: Can be tested by resizing the terminal window smaller than the game board and verifying that scroll bars or scrolling indicators appear (or keys allow panning), and the UI doesn't break.

**Acceptance Scenarios**:

1. **Given** the game content exceeds the terminal height, **When** the user moves the selection to a card off-screen, **Then** the view automatically scrolls to keep the selection visible.
2. **Given** the visual layout, **When** rendered, **Then** it occupies the full terminal size, centered or appropriately padded, without raw terminal scrolling artifacts.
3. **Given** a large tableau pile, **When** it grows long, **Then** the user can scroll down to see the bottom cards using navigation keys (j/k).

---

### User Story 3 - Persistent Selection State (Priority: P2)

Users expect their selection to remain stable and predictable interactions, especially when drawing cards.

**Why this priority**: Medium. Fixes a reported bug where selection disappeared after drawing. Essential for usability.

**Independent Test**: Can be tested by performing a draw action and verifying the selection cursor moves to the expected location (Waste pile) instead of disappearing.

**Acceptance Scenarios**:

1. **Given** the Stock pile is selected, **When** the user presses 'd' to draw, **Then** the selection cursor automatically moves to the new card on the Waste pile.
2. **Given** the Stock is empty and recycled, **When** the user interacts, **Then** the selection remains valid on the Stock pile.

## Requirements

### Functional Requirements

- **FR-001**: The system MUST render cards with a minimum size of 7 characters wide and 5 lines high.
- **FR-002**: The system MUST use a viewport mechanism to manage game board rendering within the terminal window.
- **FR-003**: The interface MUST allow vertical scrolling (and horizontal if necessary) when the game board size exceeds the terminal window size.
- **FR-004**: The visual theme MUST resemble a classic Solitaire felt table (e.g., green background) with high-contrast card colors (White/Red/Black).
- **FR-005**: Face-down cards MUST have a distinct pattern or color to look like the back of a playing card.
- **FR-006**: Drawing a card MUST automatically update the selection focus to the newly revealed card on the Waste pile.
- **FR-007**: Keyboard navigation (hjkl/arrows) MUST support scrolling the viewport when moving current selection towards off-screen elements.

### Key Entities

- **Viewport**: A UI component that manages a scrollable window into the larger game board content.
- **CardStyle**: A set of visual definitions (colors, borders, dimensions) for rendering different card states (face-up, face-down, selected).

### Edge Cases

- **Terminal Too Small**: If the terminal is smaller than the minimum viewport size (e.g., smaller than a single card), the system should display a "Terminal too small" warning or attempt to render what it can without crashing.
- **Resize Events**: When the terminal is resized, the viewport MUST recalculate its dimensions and scroll position to ensure the active card remains visible if possible.
- **Empty Piles**: Empty foundations or tableau slots must render a placeholder (e.g., empty slot outline) to maintain layout structure.

## Success Criteria

### Measurable Outcomes

- **SC-001**: Cards are rendered at 7x5 resolution or larger.
- **SC-002**: The game UI never triggers native terminal scrollback; all scrolling is handled internally by the application viewport.
- **SC-003**: User can access all cards in a max-height tableau (13+ cards) even on a small terminal window (e.g., 24 lines) via scrolling.
