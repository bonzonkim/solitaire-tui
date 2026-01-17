# Quickstart: Testing Solitaire Gameplay

- **Author**: AI Assistant
- **Created**: 2026-01-17

This guide provides instructions on how to run the application and test the new gameplay features.

## 1. Running the Application

1.  Navigate to the root directory of the project in your terminal.
2.  Execute the following command:
    ```bash
    go run ./cmd/solitaire
    ```
3.  The game board should appear in your terminal.

## 2. How to Play

The game can be controlled with either the keyboard or the mouse.

### Keyboard Controls

-   **Navigate Piles**: Use the **Arrow Keys** (`↑`, `↓`, `←`, `→`) to move the selection cursor between the different piles (Stock, Waste, Foundations, and Tableaus). The selected pile will have a colored border.
-   **Select a Card**: With a pile selected, press **Enter** or **Space** to select the top card (or the entire movable stack in a tableau). The selection border will now be on the card(s).
-   **Move a Card**: After selecting a card, navigate to a valid destination pile and press **Enter** or **Space** again to attempt the move.
-   **Draw from Stock**: Navigate to the Stock pile and press **Enter** or **Space**. If the stock is empty, this action will recycle the Waste pile.
-   **Deselect**: Press **Escape** to cancel your current card selection.
-   **Quit**: Press `q` or `Ctrl+C` to exit the application.

### Mouse Controls

-   **Select/Move**:
    1.  **Click** on a card or pile to select it. The item will be highlighted with a border.
    2.  **Click** on a valid destination pile to move the selected card there.
-   **Draw from Stock**: **Click** the Stock pile to draw a card. If the stock is empty, this will recycle the Waste pile.
-   **Drag and Drop**: You can also **click and drag** a card from its source pile and **release** it on a valid destination pile to move it.

## 3. Testing Scenarios

-   **[ ] Scenario 1: Draw and Move**
    -   Draw a card from the Stock to the Waste.
    -   Move that card from the Waste to a valid Foundation or Tableau pile.
-   **[ ] Scenario 2: Tableau Interaction**
    -   Move a card from one Tableau pile to another.
    -   Verify that the card underneath the moved card (if any) is flipped face-up.
-   **[ ] Scenario 3: Win the Game**
    -   Play the game until all cards are successfully moved to the Foundation piles.
    -   Verify that the "You Win!" message appears automatically.
-   **[ ] Scenario 4: Invalid Move**
    -   Attempt to move a card to an invalid location (e.g., a Red 5 onto a Red 6).
    -   Verify that the move is rejected and the game state does not change.
