package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/solitaire-tui/solitaire-tui/internal/game"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()

		// Global keys
		switch key {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "?":
			m.showHelp = !m.showHelp
			return m, nil
		}

		// If help is open, only allow closing
		if m.showHelp {
			if key == "esc" {
				m.showHelp = false
			}
			return m, nil
		}

		// Game interaction - handling multi-key commands and navigation
		switch key {
		case "gg":
			m.game.SetSelection(game.StockPile, 0)
			m.lastKey = ""
			m.scrollToTop() // Helper to scroll viewport to top
			return m, nil
		case "G":
			// Jump to last tableau
			m.game.SetSelection(game.TableauPile7, m.game.GetActiveCardIndex(game.TableauPile7))
			m.scrollToBottom() // Helper to scroll viewport
			return m, nil
		case "dd":
			// Draw command
			m.handleDraw()
			m.lastKey = ""
			return m, nil
		}

		// Check for multi-key start
		if key == "g" || key == "d" {
			if m.lastKey == "" {
				m.lastKey = key
				m.lastKeyTime = time.Now()
				return m, clearLastKeyAfter(commandTimeout)
			} else if m.lastKey == key {
				// Double key pressed (gg or dd) - handled above ideally, but let's handle here if missed
				if key == "d" {
					m.handleDraw()
				} else if key == "g" {
					m.game.SetSelection(game.StockPile, 0)
					m.scrollToTop()
				}
				m.lastKey = ""
				return m, nil
			}
		}

		// Navigation
		switch key {
		case "h", "left":
			m.moveSelection(-1, 0)
		case "l", "right":
			m.moveSelection(1, 0)
		case "k", "up":
			m.moveSelection(0, -1)
		case "j", "down":
			m.moveSelection(0, 1)
		case "enter", "space":
			m, cmd = m.handleSelectOrMove()
			return m, cmd
		case "esc":
			m.game.ClearSelection()
			m.sourcePileIndex = -1
			m.sourceCardIndex = -1
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		if !m.ready {
			// Initialize viewport on first size msg
			// Header height (3 approx) + Footer (1) + Padding (2) = ~6 lines reserved
			viewportHeight := msg.Height - 6
			if viewportHeight < 10 {
				viewportHeight = 10 // Minimum sensible height
			}
			m.viewport = viewport.New(msg.Width, viewportHeight)
			m.viewport.YPosition = 4 // Offset for header?
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - 6
		}

	case clearInvalidMoveMsg:
		m.showInvalidMove = false

	case clearLastKeyMsg:
		// Execute single key action if timeout
		if m.lastKey == "d" {
			m.handleDraw()
		}
		m.lastKey = ""
	}

	// Update viewport (handles mouse wheel, etc.)
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// handleDraw logic refactored for clarity and bug fixing
func (m *model) handleDraw() {
	if len(m.game.Stock.Cards) > 0 {
		m.game.DrawCard()
		// BUG FIX: Explicitly move selection to Waste pile
		m.game.SetSelection(game.WastePile, len(m.game.Waste.Cards)-1)
	} else {
		m.game.RecycleWaste()
		// Keep selection on stock
		m.game.SetSelection(game.StockPile, 0)
	}
}

// Helper to move selection and potentially scroll
func (m *model) moveSelection(dx, dy int) {
	// Logic to calculate jumping between piles...
	// Reusing existing game navigation logic but ensuring index validity

	if m.game.ActivePile == -1 {
		m.game.SetSelection(game.StockPile, 0)
		return
	}

	// Simplify: use existing rudimentary navigation or reimplement?
	// Existing navigation was a bit implicit in Update. Let's make it explicit.

	currentPile := m.game.ActivePile

	if dx != 0 {
		// Horizontal move: Cycle through piles (0-12)
		// Stock(0) -> Waste(1) -> F1-4(2-5) -> T1-7(6-12)
		// Top row logic vs Tableau logic handled by simple index wrapping?
		// Previous implementation used mod 13.

		nextPile := (currentPile + dx + 13) % 13

		// If moving to a tableau, select the last card (or same index?)
		// Usually selecting the bottom-most card is best for navigation
		cardIdx := m.game.GetActiveCardIndex(nextPile)
		m.game.SetSelection(nextPile, cardIdx)
	}

	if dy != 0 {
		// Vertical move: Only valid in Tableaus (6-12)
		if currentPile >= game.TableauPile1 && currentPile <= game.TableauPile7 {
			pile := m.game.Tableaus[currentPile-game.TableauPile1]
			currentIdx := m.game.ActiveCard

			newIdx := currentIdx + dy
			if newIdx < 0 {
				newIdx = 0
			}
			if newIdx >= len(pile.Cards) {
				newIdx = len(pile.Cards) - 1
			}
			// Only update if valid card
			if len(pile.Cards) > 0 {
				m.game.SetSelection(currentPile, newIdx)
			}

			// SCROLLING: If moving down, check if we need to scroll viewport
			// Simplified scroll logic:
			// If moving down, and index is high, scroll viewport down?
			// Since we don't know exact line numbers easily here without rendering,
			// letting viewport handle mouse/pgup/pgdn is safer for now.
		}
	}
}

func (m *model) scrollToTop() {
	m.viewport.GotoTop()
}

func (m *model) scrollToBottom() {
	m.viewport.GotoBottom()
}

// Copied from previous Update: Handle selection/move logic
func (m model) handleSelectOrMove() (model, tea.Cmd) {
	if m.sourcePileIndex == -1 {
		// Select current card
		pileIdx := m.game.ActivePile
		cardIdx := m.game.ActiveCard

		// Validate selection
		isValid := false
		if pileIdx == game.StockPile || pileIdx == game.WastePile || (pileIdx >= game.FoundationPile1 && pileIdx <= game.FoundationPile4) {
			isValid = true
		} else if pileIdx >= game.TableauPile1 && pileIdx <= game.TableauPile7 {
			// For tableau, card must be face up
			pile := m.game.GetPile(pileIdx)
			if pile != nil && cardIdx >= 0 && cardIdx < len(pile.Cards) {
				if pile.Cards[cardIdx].FaceUp {
					isValid = true
				}
			}
		}

		if isValid {
			m.sourcePileIndex = pileIdx
			m.sourceCardIndex = cardIdx
		} else {
			m.showInvalidMove = true
			return m, clearInvalidMoveAfter(2 * time.Second)
		}
	} else {
		// Try to move to current target
		targetPile := m.game.ActivePile
		if targetPile != -1 {
			success := m.game.Move(m.sourcePileIndex, m.sourceCardIndex, targetPile)
			if success {
				m.sourcePileIndex = -1
				m.sourceCardIndex = -1

				// Check victory
				if m.game.HasWon() {
					m.game.IsWon = true
				}
			} else {
				m.showInvalidMove = true
				return m, clearInvalidMoveAfter(2 * time.Second)
			}
		}
	}
	return m, nil
}
