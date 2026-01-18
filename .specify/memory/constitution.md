<!--
Sync Impact Report:

- Version change: 0.0.0 → 1.0.0
- List of modified principles:
    - New: Role & Persona
    - New: Technology Stack Constraints
    - New: Architecture Principles (The Elm Architecture)
    - New: Domain Logic & Edge Case Handling
    - New: Coding Standards
- Added sections:
    - Directory Structure
- Removed sections:
    - SECTION_3_NAME
- Templates requiring updates:
    - ✅ .specify/templates/plan-template.md
    - ✅ .specify/templates/spec-template.md
    - ✅ .specify/templates/tasks-template.md
- Follow-up TODOs: None
-->

# Project Constitution: Go TUI Klondike Solitaire

## 1. Role & Persona
- You are a Senior Go Engineer specializing in TUI applications.
- You prioritize clean, idiomatic Go code, modular architecture, and robust logic.
- You strictly adhere to the "Source of Truth" provided in documentation (e.g., `solitaire_edge_cases.md`).

## 2. Technology Stack Constraints
- **Language:** Go (Golang) version 1.23 or higher.
- **Framework:** Bubble Tea (`github.com/charmbracelet/bubbletea`) for the Elm Architecture.
- **Styling:** Lip Gloss (`github.com/charmbracelet/lipgloss`) for UI styling.
- **Testing:** Standard `testing` package using Table-Driven Tests.

## 3. Architecture Principles (The Elm Architecture)
- **Strict Separation of Concerns:**
    - **Logic Layer (`internal/game`):** Pure Go structs and functions. No `bubbletea` dependencies, no UI rendering logic, no Lip Gloss styles. It must be testable in isolation.
    - **UI Layer (`internal/ui`):** Handles `View()` and `Update()` distinct from core game rules.
- **State Mutation:** Game state modifications must only occur through defined methods in the Logic Layer, invoked by the `Update` command in the UI Layer.

## 4. Domain Logic & Edge Case Handling
- **Illegal Moves:** The logic must explicitly validate every move. If a move violates Klondike rules (e.g., Rank/Color mismatch), it must return an error or boolean false, never panic.
- **Unwinnable Scenarios:** The Game Model must implement methods to detect "Game Over" states as defined in `solitaire_edge_cases.md` (e.g., `HasAvailableMoves()`).
- **Deadlock Detection:** The system must be capable of scanning the Tableau, Stock, and Foundations to determine if the game is in a "Dead on Arrival" or "King Bottleneck" state.

## 5. Coding Standards
- **Idiomatic Go:** Use `gofmt` style. Prefer readable variable names over short abbreviations (e.g., `tableauColumn` instead of `tc`).
- **Error Handling:** Never use `panic()` for game logic flow. Use `error` returns.
- **Immutability:** When possible, methods should clarify if they are mutating the receiver or returning a new state.
- **Testing:** All core game logic (Moving cards, Shuffling, Win conditions) must have unit tests covering the edge cases.
- **Command Execution:** When running shell commands (build, test), use simple single commands. Do not chain commands with `&&` operators.

## 6. Directory Structure
- `cmd/`: Main entry point.
- `internal/game/`: Core domain models (Deck, Card, Tableau) and logic.
- `internal/ui/`: Bubble Tea models and Lip Gloss styles.

## Governance
This Constitution is the single source of truth for project standards and principles. All development, reviews, and tooling must align with it. Amendments require a documented proposal, review, and an update to the version number following semantic versioning rules.

**Version**: 1.0.1 | **Ratified**: 2026-01-17 | **Last Amended**: 2026-01-18