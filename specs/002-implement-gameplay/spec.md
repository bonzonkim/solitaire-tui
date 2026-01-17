# Feature Specification: Implement Solitaire Gameplay

- **Author**: AI Assistant
- **Stakeholders**: @b9
- **Created**: 2026-01-17
- **Status**: DRAFT

## 1. Feature Name

Implement Solitaire Gameplay and Controls

## 2. Description

The application currently displays a static, non-interactive board where card faces are hidden. This feature will transform the static display into a fully playable game of Klondike Solitaire. It involves rendering cards correctly, implementing game logic, and enabling user interaction via both mouse and keyboard.

## 3. Problem Statement

Users cannot play the game because card faces are not visible, and there are no controls to interact with the game elements. The application is currently a visual placeholder and lacks the core functionality of a Solitaire game.

## Clarifications

### Session 2026-01-17

- Q: How should the stock pile behave when it becomes empty? → A: Recycle the waste pile back into the stock an unlimited number of times.
- Q: How should the game officially end? → A: The game automatically ends and shows a "You Win!" message as soon as all cards are in the foundation piles.
- Q: How should a "selected" card or pile be visually indicated to the user? → A: Draw a colored border (e.g., cyan or yellow) around the selected item.

## 4. User Scenarios

- **As a player**, I want to see the rank and suit of each card so I can plan my moves and make strategic decisions.
- **As a player**, I want to use my mouse to click, drag, and drop cards between piles to organize them according to game rules.
- **As a player**, I want to use my keyboard to navigate between cards and piles and execute moves, providing an alternative to mouse-based interaction.
- **As a player**, I want to draw new cards from the stock pile when I have no other available moves.
- **As a player**, I want to move cards to the foundation piles to progress towards winning the game.
- **As a player**, I want to be able to move cards between the different tableau columns to reveal face-down cards and organize my layout.
- **As a player**, I want to receive clear feedback when I attempt a move that is not allowed by the game rules, so I understand my mistake.

## 5. Functional Requirements

### 5.1. Card Visibility

- FR1: All cards on the game board must be rendered to clearly display their rank (e.g., 'A', 'K', '7') and suit (e.g., '♠', '♥', '♦', '♣').
- FR2: Face-down cards in the tableau should be visually distinct from face-up cards.
- FR3: When a face-down card in a tableau pile becomes the top card, it must be automatically flipped to be face-up.

### 5.2. Game Rules & Logic

- FR4: The game must correctly implement the rules of Klondike Solitaire for moving cards:
    - **Stock to Waste**: Clicking the stock pile moves one card to the waste pile. When a player clicks the empty stock pile, the entire waste pile is moved back to the stock (face-down). This can be done an unlimited number of times.
    - **Tableau Moves**: A card or a valid stack of cards can be moved from one tableau pile to another if the move follows the descending rank, alternating color rule.
    - **Foundation Moves**: A single card can be moved to a foundation pile if it is the next card in ascending order for that suit (starting with an Ace).
    - **Waste/Tableau to Foundation**: Cards can be moved from the waste pile or a tableau pile to a foundation pile if the move is valid.
    - **Waste/Foundation to Tableau**: Cards can be moved from the waste pile or a foundation pile back to a tableau pile if the move is valid.
- FR5: The game must prevent invalid moves. When a user attempts an invalid move, the game state must not change.
- FR6: The game should provide visual feedback to the user indicating that an attempted move is invalid.
- FR7: The game will automatically detect a win condition (all 52 cards correctly placed in the foundation piles). Upon winning, the game will end and display a "You Win!" message to the player.

### 5.3. User Input

- FR8: The game must be controllable via the mouse.
    - A left-click on a card or pile should select it as the source for a move.
    - A subsequent left-click on a valid destination pile should complete the move.
    - Drag-and-drop functionality should be supported for moving cards.
- FR9: The game must be controllable via the keyboard.
    - Arrow keys (Up, Down, Left, Right) should allow the user to navigate and change the active selection between the stock, waste, foundation, and tableau piles.
    - A dedicated key (e.g., `Enter` or `Space`) should be used to select a source card/pile.
    - A subsequent press of the action key on a valid destination pile should complete the move.
    - A key (e.g., `Escape`) should be used to deselect a currently selected card.
- FR10: A selected card or pile must be clearly indicated to the user by drawing a colored border (e.g., cyan or yellow) around it.

## 6. Data Model (Key Entities)

- **Card**: Represents a standard playing card with properties for suit, rank, and face-up/face-down state.
- **Pile**: A collection of cards. The game will have several types of piles:
    - **Stock**: Holds face-down cards to be drawn.
    - **Waste**: Holds face-up cards drawn from the stock.
    - **Foundation (4x)**: Piles for building up suits from Ace to King.
    - **Tableau (7x)**: The main playing area piles, built down by alternating colors.

## 7. Assumptions

- The game variant is the standard "Klondike" Solitaire (draw one card).
- The initial shuffling and dealing of cards onto the board are handled correctly by the existing game logic.
- For keyboard controls, a simple and intuitive mapping (Arrows, Enter/Space, Escape) will be sufficient for a good user experience.

## 8. Out of Scope

- A scoring system.
- Game timers or statistics tracking.
- "Undo" or "Redo" functionality.
- Alternative game modes (e.g., Draw Three, Vegas scoring).
- Saving and loading game state.
- Animations for card movements.
- Sound effects.

## 9. Success Criteria

- SC1: A user can successfully start and complete a game of Solitaire using only the mouse for all interactions.
- SC2: A user can successfully start and complete a game of Solitaire using only the keyboard for all interactions.
- SC3: All 52 cards are correctly rendered on the board, showing their respective suits and ranks when face-up.
- SC4: The game correctly enforces all fundamental rules of Klondike Solitaire, blocking any invalid moves.
- SC5: User inputs (mouse clicks, keyboard presses) result in a visual response or state change within 200ms.
- SC6: Upon placing the final card correctly, the game automatically displays a "You Win!" message.