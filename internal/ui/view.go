package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/solitaire-tui/solitaire-tui/internal/game"
	"github.com/solitaire-tui/solitaire-tui/internal/ui/styles"
)

// View renders the game using a viewport for scrolling
func (m model) View() string {
	if !m.ready {
		return "\n  Initializing Solitaire..."
	}

	if m.showHelp {
		return styles.AppStyle.Render(m.renderHelpOverlay())
	}

	// Calculate content for the viewport
	content := m.renderGameContent()
	m.viewport.SetContent(content)

	// Build the final view with header and footer
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

// renderGameContent generates the full game board string
func (m model) renderGameContent() string {
	if m.game.IsWon {
		return lipgloss.Place(m.width, m.height-6, // Adjust for header/footer
			lipgloss.Center, lipgloss.Center,
			lipgloss.NewStyle().
				Foreground(lipgloss.Color("#00FF00")).
				Bold(true).
				Render("🎉 YOU WIN! 🎉\n\nPress 'q' to quit"),
			lipgloss.WithWhitespaceBackground(styles.AppBackground),
		)
	}

	var b strings.Builder

	// Top row: Stock, Waste, gap, Foundations
	// Add some padding at top
	b.WriteString("\n")
	b.WriteString(m.renderTopRow())
	b.WriteString("\n\n")

	// Tableaus
	b.WriteString(m.renderTableaus())
	b.WriteString("\n") // Bottom padding

	// Apply background style to the whole content
	// We need to ensure the width matches the viewport to prevent weird wrapping if logical lines are short
	contentStyle := lipgloss.NewStyle().
		Background(styles.AppBackground).
		Width(max(m.width, 80)) // Minimum width to cover game board

	return contentStyle.Render(b.String())
}

// headerView renders the title bar
func (m model) headerView() string {
	title := styles.TitleStyle.Render("♠ Solitaire TUI ♥")
	line := strings.Repeat("─", max(0, m.width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title,
		lipgloss.NewStyle().Background(styles.TitleBackground).Foreground(styles.TitleForeground).Render(line))
}

// footerView renders the status bar
func (m model) footerView() string {
	var status strings.Builder

	if m.showInvalidMove {
		status.WriteString(styles.ErrorStyle.Render("✗ Invalid move "))
	}

	if m.sourcePileIndex != -1 {
		status.WriteString(lipgloss.NewStyle().Foreground(styles.SourceBorder).Render("📌 Card selected "))
	}

	// Current pile indicator
	pileNames := []string{"Stock", "Waste", "F1", "F2", "F3", "F4", "T1", "T2", "T3", "T4", "T5", "T6", "T7"}
	if m.game.ActivePile != -1 && m.game.ActivePile < len(pileNames) {
		status.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#88FF88")).Render(fmt.Sprintf("► %s ", pileNames[m.game.ActivePile])))
	}

	status.WriteString(styles.HelpStyle.Render("│ hjkl:move Enter:select d:draw ?:help q:quit"))

	// Ensure background covers full width
	bar := lipgloss.NewStyle().
		Background(styles.TitleBackground). // Reuse darker green
		Foreground(styles.HelpTextColor).
		Width(m.width).
		Render(status.String())

	return bar
}

// renderTopRow renders Stock, Waste, and Foundations
func (m model) renderTopRow() string {
	// Helper to render an empty pile with box borders
	renderEmptyPile := func(centerText string, isActive, isSource bool) string {
		style := styles.EmptyPile

		// Select border based on state and apply color
		var borderTop, borderBottom, borderVert string
		if isSource {
			borderStyle := lipgloss.NewStyle().Foreground(styles.SourceBorder)
			borderTop = borderStyle.Render(styles.SourceBorderTop)
			borderBottom = borderStyle.Render(styles.SourceBorderBottom)
			borderVert = borderStyle.Render(styles.SourceBorderVert)
		} else if isActive {
			borderStyle := lipgloss.NewStyle().Foreground(styles.SelectedBorder)
			borderTop = borderStyle.Render(styles.SelectedBorderTop)
			borderBottom = borderStyle.Render(styles.SelectedBorderBottom)
			borderVert = borderStyle.Render(styles.SelectedBorderVert)
		} else {
			borderTop = styles.BorderTop
			borderBottom = styles.BorderBottom
			borderVert = styles.BorderVert
		}

		// Build 7-line empty pile with box borders
		line2 := borderVert + "         " + borderVert
		// Center line with text (e.g., "  ○  " or "  ♠  ")
		line4 := borderVert + "    " + centerText + "    " + borderVert
		content := borderTop + "\n" +
			line2 + "\n" +
			line2 + "\n" +
			line4 + "\n" +
			line2 + "\n" +
			line2 + "\n" +
			borderBottom
		return style.Render(content)
	}

	// Helper to render a specific pile's top card or empty slot
	renderPile := func(pileIdx int, emptyCenterText string, cards []*game.Card) string {
		isActive := m.game.ActivePile == pileIdx
		isSource := m.sourcePileIndex == pileIdx

		if len(cards) > 0 {
			topCard := cards[len(cards)-1]
			// Top row cards are simplified - just render the top one completely
			return m.renderCard(topCard, pileIdx, len(cards)-1, false)
		}

		// Empty pile with box borders
		return renderEmptyPile(emptyCenterText, isActive, isSource)
	}

	var parts []string

	// Stock
	stockActive := m.game.ActivePile == game.StockPile
	stockSource := m.sourcePileIndex == game.StockPile
	var stockStr string
	if len(m.game.Stock.Cards) > 0 {
		style := styles.FaceDownCard

		// Select border based on state and apply color
		var borderTop, borderBottom, borderVert string
		if stockSource {
			style = styles.SourceCard.Background(styles.FaceDownBackground).Foreground(styles.FaceDownForeground)
			borderStyle := lipgloss.NewStyle().Foreground(styles.SourceBorder)
			borderTop = borderStyle.Render(styles.SourceBorderTop)
			borderBottom = borderStyle.Render(styles.SourceBorderBottom)
			borderVert = borderStyle.Render(styles.SourceBorderVert)
		} else if stockActive {
			style = styles.SelectedCard.Background(styles.FaceDownBackground).Foreground(styles.FaceDownForeground)
			borderStyle := lipgloss.NewStyle().Foreground(styles.SelectedBorder)
			borderTop = borderStyle.Render(styles.SelectedBorderTop)
			borderBottom = borderStyle.Render(styles.SelectedBorderBottom)
			borderVert = borderStyle.Render(styles.SelectedBorderVert)
		} else {
			borderTop = styles.BorderTop
			borderBottom = styles.BorderBottom
			borderVert = styles.BorderVert
		}

		// Stock with box borders and ░ fill
		content := borderTop + "\n" +
			borderVert + styles.FaceDownFill + borderVert + "\n" +
			borderVert + styles.FaceDownFill + borderVert + "\n" +
			borderVert + styles.FaceDownFill + borderVert + "\n" +
			borderVert + styles.FaceDownFill + borderVert + "\n" +
			borderVert + styles.FaceDownFill + borderVert + "\n" +
			borderBottom
		stockStr = style.Render(content)
	} else {
		stockStr = renderEmptyPile("○", stockActive, stockSource)
	}
	parts = append(parts, stockStr)
	parts = append(parts, "  ") // Space

	// Waste
	parts = append(parts, renderPile(game.WastePile, " ", m.game.Waste.Cards))
	parts = append(parts, "    ") // Gap

	// Foundations
	foundations := []string{"♠", "♥", "♦", "♣"}
	for i := 0; i < 4; i++ {
		pileIdx := game.FoundationPile1 + i
		parts = append(parts, renderPile(pileIdx, foundations[i], m.game.Foundations[i].Cards))
		if i < 3 {
			parts = append(parts, " ")
		}
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, parts...)
}

// renderTableaus renders the 7 tableau piles
func (m model) renderTableaus() string {
	// We need to render columns, then join horizontally
	columns := make([]string, 7)

	for col := 0; col < 7; col++ {
		pileIdx := game.TableauPile1 + col
		pile := m.game.Tableaus[col]
		var colBuilder strings.Builder

		if len(pile.Cards) == 0 {
			// Empty pile with box borders
			style := styles.EmptyPile
			isActive := m.game.ActivePile == pileIdx
			isSource := m.sourcePileIndex == pileIdx

			// Select border based on state and apply color
			var borderTop, borderBottom, borderVert string
			if isSource {
				borderStyle := lipgloss.NewStyle().Foreground(styles.SourceBorder)
				borderTop = borderStyle.Render(styles.SourceBorderTop)
				borderBottom = borderStyle.Render(styles.SourceBorderBottom)
				borderVert = borderStyle.Render(styles.SourceBorderVert)
			} else if isActive {
				borderStyle := lipgloss.NewStyle().Foreground(styles.SelectedBorder)
				borderTop = borderStyle.Render(styles.SelectedBorderTop)
				borderBottom = borderStyle.Render(styles.SelectedBorderBottom)
				borderVert = borderStyle.Render(styles.SelectedBorderVert)
			} else {
				borderTop = styles.BorderTop
				borderBottom = styles.BorderBottom
				borderVert = styles.BorderVert
			}

			line2 := borderVert + "         " + borderVert
			line4 := borderVert + "    K    " + borderVert // K centered for King placement
			content := borderTop + "\n" +
				line2 + "\n" +
				line2 + "\n" +
				line4 + "\n" +
				line2 + "\n" +
				line2 + "\n" +
				borderBottom
			colBuilder.WriteString(style.Render(content))
		} else {
			// Stack of cards
			for i, card := range pile.Cards {
				isLast := i == len(pile.Cards)-1
				isOverlap := !isLast

				cardStr := m.renderCard(card, pileIdx, i, isOverlap)
				colBuilder.WriteString(cardStr)

				if !isLast {
					colBuilder.WriteString("\n") // Newline between stacked cards
				}
			}
		}
		columns[col] = colBuilder.String()
	}

	// Join columns with spacing
	// We can't just use JoinHorizontal because the newlines in overlapping cards
	// might mess up alignment if not careful. Lipgloss JoinHorizontal aligns by top.
	// But we need spacing between columns.

	var finalTableau []string
	for i, colStr := range columns {
		finalTableau = append(finalTableau, colStr)
		if i < 6 {
			// Spacer column
			// We effectively need a blank column or just join with margin
			// Using a 2-space string in JoinHorizontal works
			finalTableau = append(finalTableau, " ")
		}
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, finalTableau...)
}

// renderCard creates the string for a single card using Unicode Box Drawing characters
// Dimensions: 11 chars wide × 7 lines high
func (m model) renderCard(c *game.Card, pileIdx, cardIdx int, isOverlap bool) string {
	isActive := m.game.ActivePile == pileIdx && m.game.ActiveCard == cardIdx
	isSource := m.sourcePileIndex == pileIdx && m.sourceCardIndex == cardIdx

	// Create border style based on state
	var borderStyle lipgloss.Style
	var borderTopStr, borderBottomStr, borderVertStr string

	if isSource {
		// SOURCE card: thick borders with YELLOW color
		borderStyle = lipgloss.NewStyle().Foreground(styles.SourceBorder)
		borderTopStr = styles.SourceBorderTop
		borderBottomStr = styles.SourceBorderBottom
		borderVertStr = styles.SourceBorderVert
	} else if isActive {
		// ACTIVE card: double-line borders with CYAN color
		borderStyle = lipgloss.NewStyle().Foreground(styles.SelectedBorder)
		borderTopStr = styles.SelectedBorderTop
		borderBottomStr = styles.SelectedBorderBottom
		borderVertStr = styles.SelectedBorderVert
	} else {
		// NORMAL card: single-line borders (no extra color)
		borderStyle = lipgloss.NewStyle()
		borderTopStr = styles.BorderTop
		borderBottomStr = styles.BorderBottom
		borderVertStr = styles.BorderVert
	}

	// Pre-render colored borders
	borderTop := borderStyle.Render(borderTopStr)
	borderBottom := borderStyle.Render(borderBottomStr)
	borderL := borderStyle.Render(borderVertStr)
	borderR := borderStyle.Render(borderVertStr)

	if !c.FaceUp {
		// Face-down card - use face-down colors for fill
		fillStyle := styles.FaceDownCard
		fill := fillStyle.Render(styles.FaceDownFill)

		if isOverlap {
			// Overlap mode: show only top 2 lines
			return borderTop + "\n" + borderL + fill + borderR
		}

		// Full 7-line card
		return borderTop + "\n" +
			borderL + fill + borderR + "\n" +
			borderL + fill + borderR + "\n" +
			borderL + fill + borderR + "\n" +
			borderL + fill + borderR + "\n" +
			borderL + fill + borderR + "\n" +
			borderBottom
	}

	// Face Up card - determine content style
	isRed := c.Suit.Color() == "Red"
	var contentStyle lipgloss.Style
	if isRed {
		contentStyle = styles.RedSuit
	} else {
		contentStyle = styles.BlackSuit
	}

	rank := c.Rank.String()
	suitSym := c.Suit.String()

	// Build face-up card with colored content
	var rankL, rankR string
	if len(rank) == 2 {
		rankL = rank
		rankR = rank
	} else {
		rankL = rank + " "
		rankR = " " + rank
	}

	// Style the inner content (rank, suit, spaces) with card colors
	inner2 := contentStyle.Render(" " + rankL + "      ")
	inner3 := contentStyle.Render("         ")
	inner4 := contentStyle.Render("    " + suitSym + "    ")
	inner5 := contentStyle.Render("         ")
	inner6 := contentStyle.Render("      " + rankR + " ")

	if isOverlap {
		// Overlap mode: show only top 2 lines
		return borderTop + "\n" + borderL + inner2 + borderR
	}

	// Full 7-line card
	return borderTop + "\n" +
		borderL + inner2 + borderR + "\n" +
		borderL + inner3 + borderR + "\n" +
		borderL + inner4 + borderR + "\n" +
		borderL + inner5 + borderR + "\n" +
		borderL + inner6 + borderR + "\n" +
		borderBottom
}

// renderHelpOverlay renders the help popup
func (m model) renderHelpOverlay() string {
	help := `
  ♠ SOLITAIRE HELP ♥
  ──────────────────

  NAVIGATION
  h / ←     Move left
  l / →     Move right  
  k / ↑     Move up
  j / ↓     Move down
  gg        Jump to Stock
  G         Jump to Tableau 7

  ACTIONS
  Enter     Select / Move
  d / dd    Draw from Stock
  Esc       Cancel selection
  q         Quit

  Press ? or Esc to close
`
	return styles.HelpOverlay.Render(help)
}
