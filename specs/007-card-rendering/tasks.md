# Tasks: Card Visual Rendering

**Input**: Design documents from `/specs/007-card-rendering/`  
**Prerequisites**: plan.md, spec.md, research.md, data-model.md, quickstart.md

**Tests**: No automated tests requested. Manual visual verification only.

**Organization**: Tasks are grouped by user story to enable independent implementation and testing of each story.

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (e.g., US1, US2, US3)
- Include exact file paths in descriptions

## Path Conventions

- **UI**: `internal/ui/`
- **Styles**: `internal/ui/styles/`

---

## Phase 1: Setup (Shared Infrastructure)

**Purpose**: Add constants and update styles for new card rendering approach

- [x] T001 [P] Add Unicode border constants (`BorderTop`, `BorderBottom`, `BorderVert`, `FaceDownFill`, `InnerWidth`) to `internal/ui/styles/styles.go`
- [x] T002 [P] Remove `BorderStyle()` calls from card styles (`BaseCard`, `RedSuit`, `BlackSuit`, `FaceDownCard`, `EmptyPile`, `SelectedCard`, `SourceCard` variants) in `internal/ui/styles/styles.go`

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: No foundational tasks needed - all infrastructure already exists (Card model, Suit/Rank types, UI framework)

**⚠️ NOTE**: This feature modifies existing rendering logic only. No new infrastructure required.

**Checkpoint**: Setup complete - user story implementation can now begin

---

## Phase 3: User Story 1 - View Face-Up Cards (Priority: P1) 🎯 MVP

**Goal**: Render face-up cards with 11×7 Unicode box borders, rank in corners, suit centered

**Independent Test**: Launch game → Verify any face-up card shows Unicode borders (`┌┐└┘─│`), rank in top-left/bottom-right, suit at center

### Implementation for User Story 1

- [x] T003 [US1] Create `renderFaceUpCard()` helper function in `internal/ui/view.go` that builds 7-line string with box borders
- [x] T004 [US1] Implement rank padding logic in `renderFaceUpCard()` to handle both 1-char (A,2-9,J,Q,K) and 2-char (10) ranks in `internal/ui/view.go`
- [x] T005 [US1] Implement overlap mode for face-up cards (show only top 2 lines) in `internal/ui/view.go`
- [x] T006 [US1] Update `renderCard()` function to call `renderFaceUpCard()` when `c.FaceUp == true` in `internal/ui/view.go`
- [x] T007 [US1] Build and verify face-up cards render correctly with `go build -o solitaire ./cmd/solitaire && ./solitaire`

**Checkpoint**: Face-up cards should display with 11×7 Unicode borders, rank in corners, suit centered

---

## Phase 4: User Story 2 - View Face-Down Cards (Priority: P2)

**Goal**: Render face-down cards with 11×7 Unicode box borders filled with ░ pattern

**Independent Test**: Launch game → Verify face-down cards in tableaus show ░ fill pattern with same border style as face-up cards

### Implementation for User Story 2

- [x] T008 [US2] Create `renderFaceDownCard()` helper function in `internal/ui/view.go` that builds 7-line string with box borders and ░ fill
- [x] T009 [US2] Implement overlap mode for face-down cards (show only top 2 lines: border + one ░ line) in `internal/ui/view.go`
- [x] T010 [US2] Update `renderCard()` function to call `renderFaceDownCard()` when `c.FaceUp == false` in `internal/ui/view.go`
- [x] T011 [US2] Build and verify face-down cards render correctly with ░ pattern with `go build -o solitaire ./cmd/solitaire && ./solitaire`

**Checkpoint**: Face-down cards display with 11×7 Unicode borders filled with ░ pattern, visually distinct from face-up

---

## Phase 5: User Story 3 - Consistent Cross-Terminal Display (Priority: P3)

**Goal**: Verify cards render consistently across different terminal emulators

**Independent Test**: Run game in Terminal.app, iTerm2, and one Linux terminal → Verify identical character alignment

### Implementation for User Story 3

- [x] T012 [US3] Update `renderTopRow()` in `internal/ui/view.go` to use new card rendering for stock pile (with ░ pattern)
- [x] T013 [US3] Verify all empty pile placeholders use consistent box character styling in `internal/ui/view.go`
- [ ] T014 [US3] Manual verification: Test card rendering in Terminal.app, confirm 11×7 dimensions
- [ ] T015 [US3] Manual verification: Test card rendering in iTerm2, confirm visual parity with Terminal.app

**Checkpoint**: Cards render identically across multiple terminal emulators

---

## Phase 6: Polish & Cross-Cutting Concerns

**Purpose**: Final cleanup and verification

- [x] T016 Run `go build -o solitaire ./cmd/solitaire` to verify no build errors
- [x] T017 Run `go test ./...` to verify existing game logic tests still pass
- [ ] T018 Run quickstart.md verification steps to validate feature completeness

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: N/A - no foundational tasks
- **User Stories (Phase 3+)**: All depend on Setup (Phase 1) completion
  - US1 → US2 (US2 modifies same function, should wait)
  - US2 → US3 (US3 tests what US1/US2 built)
- **Polish (Phase 6)**: Depends on all user stories being complete

### User Story Dependencies

- **User Story 1 (P1)**: Can start after Setup - No dependencies on other stories
- **User Story 2 (P2)**: Depends on US1 completion (both modify `renderCard()`)
- **User Story 3 (P3)**: Depends on US1 + US2 (tests both card types)

### Within Each User Story

- Create helper function → Implement logic → Integrate → Verify
- Story complete before moving to next priority

### Parallel Opportunities

- T001 and T002 can run in parallel (different areas of same file)
- All verification tasks (T007, T011, T014-T015, T017-T018) are manual

---

## Parallel Example: Phase 1 Setup

```bash
# Both setup tasks can run in parallel since they modify different parts of styles.go:
Task T001: Add constants (new const block)
Task T002: Modify styles (existing style variables)
```

---

## Implementation Strategy

### MVP First (User Story 1 Only)

1. Complete Phase 1: Setup (T001-T002)
2. Complete Phase 3: User Story 1 (T003-T007)
3. **STOP and VALIDATE**: Face-up cards render correctly
4. Deploy/demo if ready

### Incremental Delivery

1. Complete Setup → Constants ready
2. Add User Story 1 → Test face-up cards → Demo (MVP!)
3. Add User Story 2 → Test face-down cards → Demo
4. Add User Story 3 → Cross-terminal verification → Demo
5. Each story adds value without breaking previous stories

---

## Notes

- [P] tasks = different files, no dependencies
- [Story] label maps task to specific user story for traceability
- No automated tests included (visual verification only as per spec)
- Commit after each task or logical group
- Stop at any checkpoint to validate story independently
