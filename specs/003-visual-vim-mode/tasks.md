# Tasks: Card Visual Update & Vim Mode Support

**Input**: Design documents from `/specs/003-visual-vim-mode/`
**Prerequisites**: plan.md ✅, spec.md ✅, research.md ✅, data-model.md ✅, quickstart.md ✅

## Format: `[ID] [P?] [Story] Description`

- **[P]**: Can run in parallel (different files, no dependencies)
- **[Story]**: Which user story this task belongs to (US1, US2, US3, US4)
- Include exact file paths in descriptions

## Path Conventions

- **Project Root**: `cmd/`, `internal/`, `tests/`
- **Game Logic**: `internal/game/`
- **UI**: `internal/ui/`
- **Tests**: `tests/`

---

## Phase 1: Setup

**Purpose**: Verify existing code structure and prepare for visual updates

- [X] T001 Verify current card rendering logic in `internal/ui/view.go`
- [X] T002 Verify current keyboard handling in `internal/ui/update.go`
- [X] T003 [P] Review existing styles in `internal/ui/styles/styles.go`

---

## Phase 2: Foundational (Blocking Prerequisites)

**Purpose**: Update core styles that all visual improvements depend on

**⚠️ CRITICAL**: No user story work can begin until this phase is complete

- [X] T004 Define card dimension constants (width=5, height=3) in `internal/ui/styles/styles.go`
- [X] T005 [P] Define color palette constants (RedSuit, BlackSuit, Selection, Source) in `internal/ui/styles/styles.go`
- [X] T006 [P] Create card border style with rounded corners in `internal/ui/styles/styles.go`
- [X] T007 [P] Create face-down card style with pattern in `internal/ui/styles/styles.go`
- [X] T008 Optionally update `Rank.String()` method in `internal/game/card.go` to return "10" instead of "T" for Ten

**Checkpoint**: Foundation ready - card visual work can now begin

---

## Phase 3: User Story 1 - Realistic Card Display (Priority: P1) 🎯 MVP

**Goal**: Cards display with proper borders, rank/suit symbols, and color distinguishing red vs black suits

**Independent Test**: Launch game, visually confirm cards have borders, rank in corner, suit symbols visible, red/black colors distinct

### Implementation for User Story 1

- [X] T009 [US1] Update `renderCard()` function to use new bordered style in `internal/ui/view.go`
- [X] T010 [US1] Add logic to apply red color for Hearts/Diamonds, white for Spades/Clubs in `internal/ui/view.go`
- [X] T011 [US1] Create face-down card rendering with pattern (░░░) in `internal/ui/view.go`
- [X] T012 [US1] Update selection highlight to use cyan border in `internal/ui/view.go`
- [X] T013 [US1] Update source card (being moved) highlight to use yellow border in `internal/ui/view.go`
- [X] T014 [US1] Add empty pile placeholders (suit symbols for foundations, "K" for tableaus) in `internal/ui/view.go`

**Checkpoint**: Cards should now look like real playing cards with visible rank, suit, and colors

---

## Phase 4: User Story 2 - Vim-Style Keyboard Navigation (Priority: P1) 🎯 MVP

**Goal**: Navigate using h/j/k/l keys with arrow keys still working as fallback

**Independent Test**: Use h/l to move between piles, j/k to move within tableau, confirm arrow keys also work

### Implementation for User Story 2

- [X] T015 [US2] Verify `h` key moves selection left in `internal/ui/update.go` (already implemented, may just need testing)
- [X] T016 [US2] Verify `l` key moves selection right in `internal/ui/update.go` (already implemented, may just need testing)
- [X] T017 [US2] Verify `k` key moves selection up in tableau in `internal/ui/update.go` (already implemented, may just need testing)
- [X] T018 [US2] Verify `j` key moves selection down in tableau in `internal/ui/update.go` (already implemented, may just need testing)
- [X] T019 [US2] Ensure arrow keys continue to work alongside vim keys in `internal/ui/update.go`

**Checkpoint**: Full vim navigation (h/j/k/l) and arrow keys both functional

---

## Phase 5: User Story 3 - Vim Command Mode Actions (Priority: P2)

**Goal**: Support multi-key commands (dd, gg, G) and help overlay (?)

**Independent Test**: Press `dd` to draw, `gg` to jump to Stock, `G` to jump to last tableau, `?` to show help

### Implementation for User Story 3

- [X] T020 [US3] Add `lastKey` and `lastKeyTime` fields to model struct in `internal/ui/model.go`
- [X] T021 [US3] Add `showHelp` field to model struct in `internal/ui/model.go`
- [X] T022 [US3] Implement command buffer logic with 300ms timeout in `internal/ui/update.go`
- [X] T023 [US3] Implement `gg` command to jump to Stock pile in `internal/ui/update.go`
- [X] T024 [US3] Implement `G` command to jump to last Tableau pile in `internal/ui/update.go`
- [X] T025 [US3] Implement `dd` command to draw card from stock in `internal/ui/update.go`
- [X] T026 [US3] Implement `?` key to toggle help overlay in `internal/ui/update.go`
- [X] T027 [US3] Create help overlay rendering function in `internal/ui/view.go`
- [X] T028 [US3] Style help overlay with border and proper layout in `internal/ui/styles/styles.go`

**Checkpoint**: All vim commands (gg, G, dd, ?) working, help overlay displays correctly

---

## Phase 6: User Story 4 - Card Stacking Visual (Priority: P2)

**Goal**: Tableau cards display with visible overlap showing rank/suit of underlying cards

**Independent Test**: View tableau with multiple cards, confirm all face-up cards have visible rank/suit

### Implementation for User Story 4

- [X] T029 [US4] Update `renderTableaus()` to display cards with vertical overlap in `internal/ui/view.go`
- [X] T030 [US4] Show face-down cards as thin lines (single-line representation) in `internal/ui/view.go`
- [X] T031 [US4] Ensure top card is fully visible while others show partial in `internal/ui/view.go`
- [X] T032 [US4] Handle empty tableau display with "K" placeholder in `internal/ui/view.go`

**Checkpoint**: Tableau stacking correctly shows all cards with overlapping visibility

---

## Phase 7: Polish & Cross-Cutting Concerns

**Purpose**: Final improvements affecting multiple user stories

- [X] T033 [P] Update status bar with better help text showing key controls in `internal/ui/view.go`
- [X] T034 [P] Add visual feedback for successful moves in `internal/ui/view.go`
- [ ] T035 Run all test scenarios from `quickstart.md` for manual validation
- [ ] T036 Test card rendering on terminals with 60+ columns width
- [X] T037 [P] Code cleanup and documentation updates in `internal/ui/`

---

## Dependencies & Execution Order

### Phase Dependencies

- **Setup (Phase 1)**: No dependencies - can start immediately
- **Foundational (Phase 2)**: Depends on Setup - BLOCKS all user stories
- **User Story 1 (Phase 3)**: Depends on Foundational - Card visuals
- **User Story 2 (Phase 4)**: Depends on Foundational - Can run parallel with US1
- **User Story 3 (Phase 5)**: Depends on Foundational - Can run parallel with US1/US2
- **User Story 4 (Phase 6)**: Depends on US1 completion (needs new card rendering)
- **Polish (Phase 7)**: Depends on all user stories being complete

### User Story Dependencies

- **User Story 1 (P1)**: ✅ Independent - MVP card visuals
- **User Story 2 (P1)**: ✅ Independent - MVP vim navigation (mostly already done)
- **User Story 3 (P2)**: ✅ Independent - vim commands
- **User Story 4 (P2)**: ⚠️ Depends on US1 (uses new card rendering)

### Within Each User Story

- Styles/Models before rendering logic
- Rendering logic before interaction logic
- Core implementation before polish

### Parallel Opportunities

- T003, T005, T006, T007 can run in parallel (different style definitions)
- US1, US2, US3 can run in parallel after foundational phase
- T033, T034, T037 can run in parallel (different concerns)

---

## Parallel Example: Foundational Phase

```bash
# Launch style tasks in parallel:
Task: "Define color palette constants in internal/ui/styles/styles.go"
Task: "Create card border style in internal/ui/styles/styles.go"
Task: "Create face-down card style in internal/ui/styles/styles.go"
```

---

## Implementation Strategy

### MVP First (User Stories 1 + 2 Only)

1. Complete Phase 1: Setup (verify existing code)
2. Complete Phase 2: Foundational (card styles)
3. Complete Phase 3: User Story 1 (realistic cards)
4. Complete Phase 4: User Story 2 (vim navigation - mostly verification)
5. **STOP and VALIDATE**: Test card rendering and vim navigation independently
6. Deploy/demo if ready

### Full Delivery

1. Complete MVP (above)
2. Add User Story 3 (vim commands + help overlay)
3. Add User Story 4 (card stacking)
4. Complete Polish phase

---

## Notes

- [P] tasks = different files, no dependencies
- [Story] label maps task to specific user story for traceability
- User Story 2 tasks are mostly verification since h/j/k/l already implemented
- Commit after each task or logical group
- Stop at any checkpoint to validate story independently
- Avoid: vague tasks, same file conflicts

## Summary

| Metric | Count |
|--------|-------|
| Total Tasks | 37 |
| Setup Tasks | 3 |
| Foundational Tasks | 5 |
| User Story 1 Tasks | 6 |
| User Story 2 Tasks | 5 |
| User Story 3 Tasks | 9 |
| User Story 4 Tasks | 4 |
| Polish Tasks | 5 |
| Parallel Opportunities | 11 tasks marked [P] |
| MVP Scope | User Stories 1 + 2 (14 core tasks) |
