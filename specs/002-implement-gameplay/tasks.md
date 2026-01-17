# Tasks for Feature: Implement Solitaire Gameplay

- **Feature**: [Implement Solitaire Gameplay and Controls](./spec.md)
- **Author**: AI Assistant
- **Created**: 2026-01-17
- **Status**: IN_PROGRESS (MVP Complete)

## Implementation Strategy

The implementation will be phased to ensure incremental progress. We will start with the foundational data models and visual rendering, then layer on gameplay logic, keyboard controls, and finally mouse controls. Each phase, especially those tied to a user story, will result in a testable state.

## Phase 1: Foundational - Visuals and Game State

**Goal**: Render a static but visually correct game board. All cards should be visible with their ranks and suits as specified, but there will be no interactivity.

- [X] T001 [P] In `internal/game/card.go`, define the `Suit` and `Rank` enumerations and the `Card` struct with `Suit`, `Rank`, and `FaceUp` fields.
- [X] T002 [P] In `internal/game/piles.go`, define the `Pile` struct as a slice of `*Card` and add helper methods for basic stack operations (e.g., `Push`, `Pop`, `Peek`).
- [X] T003 In `internal/game/game.go`, modify the `Game` struct to hold the complete board state: `Stock`, `Waste`, `Foundations[4]`, and `Tableaus[7]`. Also add fields for UI state: `ActivePile`, `ActiveCard`, and `IsWon`.
- [X] T004 [P] In `internal/ui/styles/styles.go`, create `lipgloss` styles for card rendering: a base style, styles for red and black suits, a style for face-down cards, and a style for the selection border.
- [X] T005 In `internal/ui/view.go`, update the `View()` function to render the entire game board by iterating through the `Game` state structs. Use the styles from `styles.go` to draw each card correctly.

## Phase 2: User Story 1 - Core Gameplay Actions

**Goal**: Implement the ability to draw from the stock, recycle the waste pile, and automatically win the game.

- [X] T006 [US1] In `internal/game/game.go`, implement a `DrawCard()` method that moves a card from the `Stock` pile to the `Waste` pile.
- [X] T007 [US1] In `internal/game/game.go`, implement a `RecycleWaste()` method that moves all cards from the `Waste` pile back to the `Stock`, ensuring they are face-down.
- [X] T008 [US1] In `internal/game/game.go`, implement `CheckWinCondition()` to verify if all cards are in the foundation piles and set the `IsWon` flag.
- [X] T009 [US1] In `internal/ui/update.go`, add logic to the `Update` function to call `DrawCard()` or `RecycleWaste()` when the stock is activated (via keyboard or mouse).
- [X] T010 [US1] In `internal/ui/view.go`, add logic to display the "You Win!" message when `Game.IsWon` is true.
- [X] T011 [P] [US1] In `tests/game/game_test.go`, write unit tests for the `DrawCard`, `RecycleWaste`, and `CheckWinCondition` methods.

## Phase 3: User Story 2 - Keyboard Interaction

**Goal**: Enable a user to play a complete game using only the keyboard.

- [X] T012 [US2] In `internal/game/game.go`, implement the core `Move()` method, which takes source and destination info, validates the move against Klondike rules, and executes it. This method should also handle flipping newly revealed tableau cards.
- [X] T013 [US2] In `internal/game/game.go`, implement methods to manage the UI selection state, such as `SetSelection()` and `ClearSelection()`.
- [X] T014 [US2] In `internal/ui/update.go`, add logic to handle `tea.KeyMsg` for arrow keys to call `SetSelection()`, changing the active pile.
- [X] T015 [US2] In the keyboard handler in `internal/ui/update.go`, implement the logic for the **Enter/Space** keys to either select a card or call the `Move()` method.
- [X] T016 [US2] In the keyboard handler in `internal/ui/update.go`, implement the logic for the **Enter/Space** keys to either select a card or, if a card is already selected, call the `Move()` method.
- [X] T017 [US2] In the keyboard handler in `internal/ui/update.go`, implement the logic for the **Escape** key to call `ClearSelection()`.
- [X] T018 [P] [US2] In `tests/game/game_test.go`, write comprehensive table-driven tests for the `Move()` method, covering all valid and invalid moves described in the spec.

## Phase 4: User Story 3 - Mouse Interaction

**Goal**: Enable a user to play a complete game using the mouse.

- [X] T019 [US3] In `internal/ui/update.go`, add a `case tea.MouseMsg` to the `Update` function to handle mouse input.
- [X] T020 [US3] In the mouse handler in `internal/ui/update.go`, add logic to determine which pile/card was clicked based on the mouse event's coordinates.
- [X] T021 [US3] In the mouse handler in `internal/ui/update.go`, implement the "click-select, click-move" logic by calling the `SetSelection()` and `Move()` methods from the game package.
- [ ] T022 [P] [US3] In `internal/ui/update.go`, add handlers for `MouseDrag` and `MouseRelease` events to implement drag-and-drop functionality. (Note: This is more complex and can be considered a stretch goal if time is a constraint).

## Phase 5: Polish & Finalization

**Goal**: Ensure the feature is complete, correct, and adheres to quality standards.

- [X] T023 In `internal/ui/view.go`, add visual feedback for invalid moves (e.g., a brief flash of color).
- [X] T024 Review all new code for style, clarity, and adherence to the project constitution.
- [ ] T025 Manually execute all test scenarios outlined in `quickstart.md` to confirm functionality.

## Dependencies

- **Phase 2** depends on **Phase 1**.
- **Phase 3** depends on **Phase 2**.
- **Phase 4** depends on **Phase 3** (specifically, the `Move` and selection logic in `internal/game`).
- **Phase 5** depends on all previous phases.

**MVP Scope**: Completing Phase 1, 2, and 3 would deliver a complete, keyboard-playable game, which constitutes the Minimum Viable Product for this feature.
