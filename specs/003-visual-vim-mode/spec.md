# Feature Specification: Card Visual Update & Vim Mode Support

**Feature Branch**: `003-visual-vim-mode`  
**Created**: 2026-01-17  
**Status**: Draft  
**Input**: User description: "카드를 더 실제카드처럼 보여줄 수 있게 비주얼 업데이트를 할거야. 또한 vim 모드지원을 할거야."

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Realistic Card Display (Priority: P1)

As a player, I want the cards in the game to look more like real playing cards so that the game feels more authentic and enjoyable to play.

**Why this priority**: Visual appeal is crucial for user engagement and game immersion. Without realistic-looking cards, the game feels like a developer prototype rather than a polished product.

**Independent Test**: Can be fully tested by launching the game and visually confirming that cards display with proper structure (rank in corners, suit symbols, borders) and are clearly distinguishable from each other.

**Acceptance Scenarios**:

1. **Given** a face-up card is displayed, **When** the user views the card, **Then** the card shows the rank in the top-left corner, suit symbol, and has a defined rectangular border resembling a real playing card.
2. **Given** a face-down card is displayed, **When** the user views the card, **Then** the card shows a distinct back pattern that clearly indicates it is face-down.
3. **Given** red and black suit cards, **When** displayed side by side, **Then** the color difference is clearly visible and distinguishable.
4. **Given** any card is selected, **When** the user highlights it, **Then** the card has a visible highlight border that stands out from other cards.

---

### User Story 2 - Vim-Style Keyboard Navigation (Priority: P1)

As a Vim user, I want to navigate and interact with the game using familiar Vim keybindings (h/j/k/l) so that I can play comfortably without learning new controls.

**Why this priority**: Vim users (the target audience for a TUI game) expect Vim-style navigation. This is essential for the core user experience.

**Independent Test**: Can be fully tested by navigating through all game piles and selecting cards using only h/j/k/l keys, then confirming all movements work as expected.

**Acceptance Scenarios**:

1. **Given** a pile is selected, **When** the user presses `h`, **Then** the selection moves to the pile on the left.
2. **Given** a pile is selected, **When** the user presses `l`, **Then** the selection moves to the pile on the right.
3. **Given** a tableau pile with multiple cards is selected, **When** the user presses `k`, **Then** the selection moves up to the previous card in the pile.
4. **Given** a tableau pile with multiple cards is selected, **When** the user presses `j`, **Then** the selection moves down to the next card in the pile.
5. **Given** Vim mode is active, **When** the user presses arrow keys, **Then** the arrow keys also work for navigation (dual support).

---

### User Story 3 - Vim Command Mode Actions (Priority: P2)

As a Vim user, I want to perform game actions using Vim-like commands (e.g., `dd` to draw, `gg` to go to first pile) so that the gameplay feels consistent with Vim's modal editing paradigm.

**Why this priority**: While navigation (P1) is essential, Vim-style commands add an advanced layer of convenience that enhances the experience for power users.

**Independent Test**: Can be tested by drawing a card using `dd`, going to the first pile with `gg`, and moving cards with appropriate commands.

**Acceptance Scenarios**:

1. **Given** the game is running, **When** the user presses `dd`, **Then** a card is drawn from the stock to the waste pile.
2. **Given** the game is running, **When** the user presses `gg`, **Then** the selection jumps to the Stock pile (first pile).
3. **Given** the game is running, **When** the user presses `G`, **Then** the selection jumps to the last tableau pile.
4. **Given** a card is selected, **When** the user presses `u`, **Then** the last move is undone (if possible).
5. **Given** the game is running, **When** the user presses `?`, **Then** a help overlay displays all available keybindings.

---

### User Story 4 - Card Stacking Visual (Priority: P2)

As a player, I want to see tableau cards properly stacked with visible overlap so that I can see the cards underneath and plan my moves.

**Why this priority**: Proper card stacking is essential for gameplay visibility. Without seeing underlying cards, strategic planning becomes impossible.

**Independent Test**: Can be tested by viewing a tableau pile with multiple cards and confirming that each card's rank and suit is visible in the overlapping section.

**Acceptance Scenarios**:

1. **Given** a tableau pile has multiple face-up cards, **When** displayed, **Then** each card overlaps the previous card, showing only the top portion (rank/suit) of cards beneath.
2. **Given** a tableau pile has face-down cards, **When** displayed, **Then** face-down cards are shown as thin horizontal lines representing the card backs.
3. **Given** any pile is empty, **When** displayed, **Then** an empty placeholder is shown indicating where cards can be placed.

---

### Edge Cases

- What happens when the terminal is very small (< 40 columns)?
  - Cards should display in a compact mode, showing minimal information but remaining functional.
- How does the system handle rapid key presses?
  - The system should debounce rapid inputs and process them sequentially without lag or dropped inputs.
- What happens when Vim commands conflict with standard controls?
  - Vim keybindings take precedence, but arrow keys continue to work as alternatives.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Cards MUST display with a border resembling a real playing card (rectangular with rounded corners).
- **FR-002**: Face-up cards MUST show the rank (A, 2-10, J, Q, K) and suit symbol (♠, ♥, ♦, ♣) clearly.
- **FR-003**: Face-down cards MUST display a visually distinct card back pattern.
- **FR-004**: Red suits (Hearts, Diamonds) MUST be displayed in a clearly visible red color.
- **FR-005**: Black suits (Spades, Clubs) MUST be displayed in a clearly visible light/white color.
- **FR-006**: The game MUST support `h`, `j`, `k`, `l` keys for left, down, up, right navigation respectively.
- **FR-007**: The game MUST continue to support arrow keys alongside Vim keybindings.
- **FR-008**: Selected cards/piles MUST have a visible highlight (distinct border color).
- **FR-009**: Tableau cards MUST be displayed with vertical overlap showing rank/suit of underlying cards.
- **FR-010**: The game MUST display a help text showing available keybindings.
- **FR-011**: The `dd` command MUST draw a card from the stock.
- **FR-012**: The `gg` command MUST jump to the first pile (Stock).
- **FR-013**: The `G` command MUST jump to the last tableau pile.
- **FR-014**: The `?` key MUST display a help overlay with all available commands.

### Key Entities

- **Card**: A playing card with rank, suit, and face-up/face-down state. Visual representation includes border, color, and symbols.
- **Pile**: A collection of cards (Stock, Waste, Foundation, Tableau) with specific visual stacking rules.
- **Selection State**: The currently highlighted pile and card, displayed with a distinct visual indicator.
- **Vim Mode**: The modal state that determines how keystrokes are interpreted (navigation vs. command).

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: 100% of card ranks and suits are immediately identifiable by users without guessing.
- **SC-002**: New Vim users can navigate all piles using h/j/k/l within 30 seconds of starting the game.
- **SC-003**: Card display renders correctly on terminals with 60+ columns width.
- **SC-004**: All Vim keybindings (h/j/k/l, gg, G, dd, ?) work consistently with expected Vim behavior.
- **SC-005**: Face-up and face-down cards are visually distinguishable at a glance.
- **SC-006**: Users can see at least the rank and suit of all face-up cards in a stacked tableau pile.

## Assumptions

- Vim keybindings (h/j/k/l) are already partially supported (based on existing codebase).
- The TUI framework (Bubble Tea) supports the required styling for card visuals.
- Terminal color support (256 colors or true color) is available.
- Users are familiar with basic Vim navigation concepts.
- The game currently displays cards in a minimal format that needs enhancement.

## Out of Scope

- Custom keybinding configuration (use preset Vim bindings)
- Full Vim modal editing (only navigation and game-action commands)
- Card animation effects
- Sound effects
- Mouse-free mode (mouse support remains alongside Vim controls)
