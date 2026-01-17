# Implementation Plan - Realistic ASCII Cards

This plan details the implementation of "photorealistic" ASCII cards, increasing dimensions and adding specific art for face cards and pip patterns for number cards.

## Proposed Changes

### UI Styling & Configuration

#### [MODIFY] [styles.go](file:///Users/b9/bonzonkim/github.com/solitaire-tui/solitaire-tui/internal/ui/styles/styles.go)
-   **Increase Dimensions**: Update `CardWidth` to 11 and `CardHeight` to 7 (to support 7x5 or 9x5 internal drawing area).
-   **Define Art Patterns**:
    -   Add string constants or a map for Face Card art (K, Q, J) for each suit (or shared).
    -   Add logic/maps for Pip placement (e.g. for '5': corners + center).

### UI Rendering

#### [MODIFY] [view.go](file:///Users/b9/bonzonkim/github.com/solitaire-tui/solitaire-tui/internal/ui/view.go)
-   **Update `renderCard`**:
    -   Change content generation logic.
    -   **Face Cards**: If Rank is K, Q, J, lookup and insert the corresponding ASCII art into the center of the card string.
    -   **Number Cards**: Generate a string grid based on the pip count.
    -   **Ace**: Special large center pip.
    -   Ensure borders and padding respect the new 11x7 size.

## Verification Plan

### Manual Verification
1.  **Launch Game**: `go run ./cmd/solitaire`.
2.  **Size Check**: Confirm cards are significantly larger (11x7).
3.  **Face Card Check**:
    -   Find a King. Verify it has a "Crown" or similar ASCII art.
    -   Find a Queen/Jack. Verify distinct art.
4.  **Pip Check**:
    -   Find a '10'. Verify it has 10 pips (or a dense pattern).
5.  **Scroll Check**:
    -   Since the board is wider (~90-100 chars), verify horizontal scrolling or that it fits standard 100+ col terminals.
