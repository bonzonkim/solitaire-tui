# Feature Specification: Visual Polish

**Feature Branch**: `005-visual-polish`
**Created**: 2026-01-17
**Status**: Draft
**Input**: User description: "빨간색 카드에 포커스 가있으면 검정색으로 보여. 포커스중이어도 색깔을 유지하게 해주고 배경을 좀 더 고급지게 바꿔줘" (Red cards look black when focused. Keep the color even when focused, and make the background more premium.)

## User Scenarios & Testing

### User Story 1 - Color-Preserving Selection (Priority: P1)

Users need to clearly distinguish red suits (Hearts, Diamonds) from black suits even when a card is currently selected/focused.

**Why this priority**: High. Losing color information on selection is a usability regression and breaks the visual metaphor of playing cards.

**Independent Test**:
1.  Navigate to a red card (Heart/Diamond).
2.  Verify the suit symbol and rank text remain **RED**, not converting to black/white.
3.  Verify the selection is still visible (e.g., via background color change or border highlight) without overriding the foreground text color.

**Acceptance Scenarios**:

1.  **Given** a red card (e.g., 3 of Hearts) on the board, **When** the user selects it with the cursor, **Then** the text "3♥" remains displayed in red color.
2.  **Given** a black card (e.g., King of Spades), **When** selected, **Then** the text remains black/white (contrast appropriate).
3.  **Given** any selected card, **When** rendered, **Then** it is clearly distinguishable from non-selected cards (e.g., via cyan/gold border or distinct background that supports red text).

---

### User Story 2 - Premium Background Visuals (Priority: P2)

Users want a more "luxurious" and high-quality game atmosphere, replacing the basic background with something that feels like a premium casino table.

**Why this priority**: Medium. Improves user satisfaction and aesthetics as requested ("make it prettty").

**Independent Test**:
Launch the game and observe the background. It should not be a flat, default terminal color or a jarring shade. It should resemble a quality felt table.

**Acceptance Scenarios**:

1.  **Given** the game board, **When** rendered, **Then** the background uses a rich, "premium" color (e.g., deep emerald green or textured felt color, distinct from the previous basic green).
2.  **Given** the board, **When** rendered, **Then** the contrast between cards and background is high, ensuring readability.

## Requirements

### Functional Requirements

-   **FR-001**: The system MUST render text for Heart and Diamond suits in RED color used when the card is in the `Selected` state.
-   **FR-002**: The selection indicator MUST NOT force a foreground color that obliterates the suit color (e.g., forcing black text on cyan background).
-   **FR-003**: The application background MUST be updated to a "premium" shade (e.g., matching standard casino felt hex codes like `#35654d` or similar deep rich tones) or use a texture pattern if feasible without noise.

### Key Entities

-   **SelectedCardStyle**: The Lipgloss style definition for the currently active card.
-   **AppBackground**: The color/style definition for the global application background.

### Edge Cases

-   **Face-Down Cards Selection**: When a face-down card is selected, it should retain its pattern/back visual and not try to display suit colors (as they are hidden).
-   **Theme Contrast**: The "premium" background color MUST maintain a contrast ratio of at least 4.5:1 with standard white text and the card colors to ensure accessibility.

## Success Criteria

### Measurable Outcomes

-   **SC-001**: Red cards selected by the user pass visual verification of retaining red text.
-   **SC-002**: Background color is updated to a new, distinct "premium" hex value.
