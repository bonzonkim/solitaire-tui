# Implementation Plan: Implement Solitaire Gameplay

- **Feature**: [Implement Solitaire Gameplay and Controls](./spec.md)
- **Author**: AI Assistant
- **Created**: 2026-01-17
- **Status**: DRAFT

## 1. Technical Context

- **Language/Framework**: Go 1.23+ with the Charm Bracelet libraries: Bubble Tea for the application framework and Lip Gloss for styling. This aligns with the established project stack.
- **Dependencies**: No new major external dependencies are required. The existing `bubbletea` and `lipgloss` libraries are sufficient.
- **Data Storage**: All game state will be managed in-memory within Go structs. No database or persistent storage is needed for this feature.
- **APIs/Integrations**: None. This is a self-contained TUI application.
- **Hosting/Deployment**: The application is a local binary, executed via `go run` or a compiled executable.
- **Testing Strategy**: Go's standard `testing` package will be used for table-driven unit tests. The focus will be on testing the game logic in the `internal/game` package to ensure rules are enforced correctly.
- **Observability**: Standard logging via Go's `log` package may be used for debugging during development if necessary. No production-level observability is required.
- **Security**: Not applicable for this offline, local TUI game.
- **Assumptions/Unknowns**: The implementation will need to define a concrete, user-friendly navigation flow for the keyboard (e.g., how arrow keys move between all 13 piles). This is considered a detailed design decision for the implementation phase.

## 2. Constitution Check

The proposed plan was checked against the [Project Constitution](../../.specify/memory/constitution.md).

- **[PASS] Role & Persona**: The plan adopts the mindset of a Senior Go Engineer building a TUI.
- **[PASS] Technology Stack Constraints**: The plan exclusively uses the approved stack: Go 1.23+, Bubble Tea, and Lip Gloss.
- **[PASS] Architecture Principles**: The implementation will strictly enforce the separation of the logic layer (`internal/game`) and the UI layer (`internal/ui`). Game state will be managed by `internal/game` and updated via messages from the `internal/ui` `Update` function.
- **[PASS] Domain Logic & Edge Case Handling**: Core game logic for moves, rules, and win conditions will be implemented and unit-tested within the `internal/game` package, adhering to the constitution's requirements.
- **[PASS] Coding Standards**: The implementation will follow `gofmt` style, standard error handling, and include table-driven unit tests for core logic.
- **[PASS] Directory Structure**: The plan will utilize the existing `cmd/`, `internal/game/`, and `internal/ui/` directories as defined.

**Result**: The plan is in full compliance with the project constitution.

## 3. Phase 0: Research

No significant research was required for this feature. The problem domain (Klondike Solitaire) is well-defined, and the technology stack is pre-determined by the project constitution. A minimal `research.md` will be created to reflect this.

## 4. Phase 1: Design

The following design artifacts will be generated based on the feature specification:

- **`data-model.md`**: A document detailing the core Go structs that will represent the game state, such as `Card`, `Pile`, and the main `Game` model.
- **`quickstart.md`**: A guide explaining how to run the application and test the new gameplay features using both keyboard and mouse.
- **API Contracts**: Not applicable for this feature, as it is not an API-driven service.

## 5. Phase 2: Implementation Strategy

The implementation will focus on integrating the user-facing UI with the backend game logic.

1.  **Game State Enhancement (`internal/game`)**:
    -   The core `Game` struct in `internal/game/game.go` will be expanded to track the UI state, such as the currently selected pile and card.
    -   New methods will be added to handle user actions like selecting a card (`SelectCard()`) and attempting a move (`MoveSelectedCard()`). These methods will contain the core logic for validating moves against the game rules.

2.  **Visual Overhaul (`internal/ui/view.go`)**:
    -   The `View()` method will be rewritten to render the full game board.
    -   It will iterate through the `Game` state to draw each pile (Stock, Waste, Foundations, Tableaus).
    -   `lipgloss` will be used to render cards with their actual rank and suit. Red/black colors will be used for suits.
    -   A colored border will be drawn around the currently selected card/pile, based on the state in the `Game` struct.
    -   A "You Win!" message view will be created and displayed when the game is won.

3.  **Interaction Logic (`internal/ui/update.go`)**:
    -   The `Update()` method will be the central hub for interactivity. It will be expanded to handle `tea.MouseMsg` and `tea.KeyMsg`.
    -   **Mouse**: It will process clicks, determine which pile/card was clicked, and send messages to the `Game` model to select or move cards.
    -   **Keyboard**: It will process arrow keys to change the selection, `Enter`/`Space` to select/move cards, and `Escape` to deselect.
    -   The `Update` function will call the appropriate methods in the `internal/game` package and update its model based on the results.

4.  **Testing (`tests/game/`)**:
    -   New unit tests will be added to `game_test.go` to cover the new game logic:
        -   Test valid and invalid card moves between all pile types.
        -   Test the stock recycling logic.
        -   Test the win condition detection.
        -   Test flipping of tableau cards.