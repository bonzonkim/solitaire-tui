# Implementation Plan - Visual Polish

This plan addresses visual feedback improvements, specifically ensuring red cards retain their color when selected, and updating the game background to a more premium aesthetics.

## Proposed Changes

### UI Styling

#### [MODIFY] [styles.go](file:///Users/b9/bonzonkim/github.com/solitaire-tui/solitaire-tui/internal/ui/styles/styles.go)
-   **Update `AppBackground`**: Change to a richer, premium felt green (e.g., `#0a4d0a` -> `#052d05` or similar deep tone).
-   **Add Selection Styles**:
    -   `SelectedRedCard`: Based on `BaseCard`, with `SelectedBorder` (Cyan) and `RedSuitColor` foreground.
    -   `SelectedBlackCard`: Based on `BaseCard`, with `SelectedBorder` (Cyan) and `BlackSuitColor` foreground.
    -   `SourceRedCard` / `SourceBlackCard`: distinct borders for moving cards while preserving color.

### UI Rendering

#### [MODIFY] [view.go](file:///Users/b9/bonzonkim/github.com/solitaire-tui/solitaire-tui/internal/ui/view.go)
-   **Update `renderCard` logic**:
    -   Instead of applying a generic `styles.SelectedCard`, check the suit color when `isActive` is true.
    -   If Red: apply `styles.SelectedRedCard`.
    -   If Black: apply `styles.SelectedBlackCard`.
    -   Similar logic for `isSource` (moving) cards.

## Verification Plan

### Manual Verification
1.  **Launch Game**: Run `go run ./cmd/solitaire`.
2.  **Background Check**: Confirm the background looks deeper/richer.
3.  **Selection Check**:
    -   Navigate to a Heart/Diamond. Verify text is **RED**.
    -   Navigate to a Spade/Club. Verify text is **BLACK**.
    -   Verify the Cyan double border still appears for selection.
