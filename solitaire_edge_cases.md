# Context: Klondike Solitaire Edge Cases & Unwinnable Scenarios

## 1. Objective
We are building a TUI (Text-based User Interface) Klondike Solitaire game. The agent must understand that not all random deals are solvable. Approximately 20% of random deals are theoretically unwinnable.

The code must handle these "Game Over" states gracefully, either by detecting them, preventing them (via solvable seed generation), or allowing user mitigation (Undo/Reshuffle).

## 2. Unwinnable Scenarios (Edge Cases)

### A. The King Bottleneck (Empty Column Deadlock)
* **Rule:** Only a King (K) can be placed in an empty Tableau column.
* **Scenario:** A player clears a column, creating an empty slot. However, all Kings are either:
    1.  Buried deep in the Stockpile (and inaccessible due to current card ordering).
    2.  Located within "face-down" cards in other Tableau columns.
* **Result:** The empty space cannot be utilized to unblock other columns, halting progress.

### B. Color/Parity Lock (The "Cycle of Doom")
* **Rule:** Tableau builds must be in descending order and alternating colors (Red/Black).
* **Scenario:**
    * Player needs to move a **Black 7**.
    * To move it, a **Red 8** is required.
    * However, both Red 8s are currently inaccessible:
        * One is buried *underneath* the Black 7 (in a face-down pile).
        * The other is in the Stockpile, blocked by cards that require the Black 7 to be moved first.
* **Result:** A circular dependency prevents any further moves.

### C. Buried Aces (Foundation Block)
* **Rule:** Foundations must be built up from Ace to King.
* **Scenario:** An Ace (e.g., Ace of Spades) is the very bottom card of a Tableau pile (index 0, face-down).
* **Condition:** High-ranking cards (Kings, Queens) are stacked on top of this pile.
* **Result:** To retrieve the Ace, the player must move the entire stack to another column. If no other King-ready empty columns exist (see "King Bottleneck"), the Ace remains trapped forever.

### D. The "Draw 3" Loop
* **Rule:** (If implementing Draw 3 rules) The player draws 3 cards at a time.
* **Scenario:** The Stockpile is cycled through multiple times without any cards being played to the Tableau.
* **Result:** Since the order of cards in the Stockpile does not change unless a card is removed, the visible cards in the "fan" of 3 remain identical in every pass. The player enters an infinite loop.

### E. Zero Moves Initial Deal (Dead on Arrival)
* **Probability:** Rare (~0.025%) but possible.
* **Scenario:**
    * The initial 7 Tableau piles have no valid moves (no cards can be moved to Foundations or other Tableau piles).
    * The first cycle of the Stockpile offers no playable cards.
* **Result:** The game is lost immediately after the deal.

---

## 3. Implementation Directives for the Agent

When writing the Game Logic, please consider the following approaches:

### Option A: Unsolvable State Detection (Game Over Check)
Implement a function `CheckGameState()` that runs after every move:
1.  Are there any moves available from Tableau to Foundation?
2.  Are there any moves available between Tableau columns?
3.  Are there any moves available from Stock/Waste to Tableau/Foundation?
4.  **IF** all are FALSE **AND** Stockpile has been cycled completely **THEN** trigger `GameOver(Win=False)`.

### Option B: "Solvable Deck" Generator (Advanced)
Instead of `rand.Shuffle`, implement a Solver/Simulation step during initialization:
1.  Generate a random deck.
2.  Run a fast internal simulation (DFS or Greedy algorithm) to see if the deck *can* be solved.
3.  If unsolvable, discard and regenerate.
4.  Serve only solvable seeds to the user.

### Option C: Mitigation Features
If we stick to random deals, implement these features in the TUI:
* **Unlimited Undo:** Store game states in a stack to allow backtracking from deadlocks.
* **Hint System:** Highlight available moves to confirm to the user that no moves exist.
