# Research: Solitaire TUI Game

## 1. Best Practices for Bubble Tea

- **Decision**: Follow the Model-View-Update (MVU) architecture provided by Bubble Tea.
- **Rationale**: The MVU architecture is the core of Bubble Tea and provides a clear and scalable structure for TUI applications. It promotes a unidirectional data flow, making the application easier to reason about and debug.
- **Alternatives Considered**: None, as this is the fundamental paradigm of the chosen framework.

## 2. Best Practices for Lip Gloss

- **Decision**: Define a consistent color and style palette in a central `styles` package.
- **Rationale**: Centralizing styles will make the application's appearance consistent and easy to modify. It also keeps styling logic separate from the UI components.
- **Alternatives Considered**: Defining styles inline within each component, which would lead to code duplication and inconsistency.

## 3. Algorithms for Detecting Unwinnable States

- **Decision**: Implement a heuristic-based approach to detect unwinnable states. The initial implementation will check for the following conditions:
    - No more moves from the stock to the waste.
    - No more moves from the waste to the tableau or foundations.
    - No more moves within the tableau.
    - No more moves from the tableau to the foundations.
- **Rationale**: A full "dead on arrival" or "king bottleneck" analysis is complex and may not be necessary for the initial version. A heuristic-based approach provides a good balance of performance and accuracy for common unwinnable scenarios. More advanced detection can be added later if needed.
- **Alternatives Considered**: A full graph-based analysis of the game state, which would be more accurate but also significantly more complex to implement.
