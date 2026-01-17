# Tasks: Solitaire TUI Game

**Input**: Design documents from `/specs/001-solitaire-tui-game/`
**Prerequisites**: plan.md (required), spec.md (required for user stories), research.md, data-model.md

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Project initialization and basic structure

- [x] T001 Initialize Go module in the project root (`go mod init github.com/solitaire-tui/solitaire-tui`)
- [x] T002 Add dependencies (`go get github.com/charmbracelet/bubbletea github.com/charmbracelet/lipgloss`)
- [x] T003 [P] Create directory structure (`cmd/solitaire`, `internal/game`, `internal/ui`, `tests/game`)

---

## Phase 2: Foundational (Game Logic)

**Purpose**: Core game logic that MUST be complete before ANY user story can be implemented

- [x] T004 [P] Implement `Card`, `Deck`, `Tableau`, `Stock`, `Waste`, and `Foundation` structs in `internal/game/`
- [x] T005 [P] Implement card shuffling and initial dealing logic in `internal/game/deck.go`
- [x] T006 Implement game state initialization in `internal/game/game.go`
- [x] T007 Implement move validation logic in `internal/game/game.go` for all move types (tableau to tableau, tableau to foundation, waste to tableau, waste to foundation)
- [x] T008 [P] Write unit tests for game logic in `tests/game/game_test.go`

---

## Phase 3: User Story 1 - Play a game of Solitaire (Priority: P1) 🎯 MVP

**Goal**: A user can play a complete game of solitaire from start to finish.

**Independent Test**: Run the application, move cards according to the rules, and see the game end with a win or lose message.

### Implementation for User Story 1

- [x] T009 [P] Create the main application model in `internal/ui/model.go`
- [x] T010 Implement the `Init` and `Update` functions in `internal/ui/update.go`
- [x] T011 Implement the `View` function in `internal/ui/view.go` to render the game board.
- [x] T012 Implement the main application entry point in `cmd/solitaire/main.go`
- [x] T013 Integrate the game logic from `internal/game` with the UI layer in `internal/ui`

---

## Phase 4: User Story 4 - Handle Unwinnable Games (Priority: P1)

**Goal**: The game automatically ends when it becomes unwinnable.

**Independent Test**: Set up a game state that is known to be unwinnable and verify that the game ends and displays the correct message.

### Implementation for User Story 4

- [x] T014 Implement the unwinnable state detection logic in `internal/game/game.go` based on the research in `research.md`.
- [x] T015 Integrate the unwinnable state detection into the UI in `internal/ui/update.go`.
- [x] T016 Display an "Unwinnable game" message in `internal/ui/view.go` when the game is unwinnable.

---

## Phase 5: User Story 2 - Control the game with the keyboard (Priority: P2)

**Goal**: A user can play the game using only the keyboard with Vim-like bindings.

**Independent Test**: Navigate the game board, select cards, and move them using only the keyboard.

### Implementation for User Story 2

- [x] T017 [P] Add key handling to the `Update` function in `internal/ui/update.go` to process 'h', 'j', 'k', 'l', and 'enter' keys.
- [x] T018 Implement cursor movement between different parts of the game board (tableau, stock, waste, foundation).
- [x] T019 Implement card selection and movement logic based on keyboard input.

---

## Phase 6: User Story 3 - Control the game with the mouse (Priority: P3)

**Goal**: A user can play the game using the mouse to click and drag cards.

**Independent Test**: Navigate the game board, select cards, and move them using only the mouse.

### Implementation for User Story 3

- [x] T020 [P] Add mouse handling to the `Update` function in `internal/ui/update.go`.
- [x] T021 Implement card selection based on mouse clicks.
- [x] T022 Implement card movement based on mouse clicks and drags.

---

## Phase 7: Polish & Cross-Cutting Concerns

**Purpose**: Improvements that affect multiple user stories

- [x] T023 [P] Implement the "You Win!" message when the game is won.
- [x] T024 Implement the visual feedback for invalid moves (shake card, show message) as specified in the clarifications.
- [x] T025 Implement the terminal size check on startup.
- [x] T026 [P] Refine the styles and colors in a central `styles` package.

---

## Dependencies & Execution Order

- **Phase 1 & 2**: Must be completed before any user story work can begin.
- **Phase 3 (US1)**: Can begin after Phase 2 is complete.
- **Phase 4 (US4)**: Depends on Phase 3.
- **Phase 5 (US2) & 6 (US3)**: Can be worked on in parallel after Phase 3.
- **Phase 7**: Can be worked on throughout the project, but should be finalized after all user stories are complete.
