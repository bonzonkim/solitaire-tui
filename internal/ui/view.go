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
	// Card width + spacing
	// We need careful horizontal layout

	// Helper to render a specific pile's top card or empty slot
	renderPile := func(pileIdx int, emptyText string, cards []*game.Card) string {
		isActive := m.game.ActivePile == pileIdx

		if len(cards) > 0 {
			topCard := cards[len(cards)-1]
			// Top row cards are simplified - just render the top one completely
			return m.renderCard(topCard, pileIdx, len(cards)-1, false)
		}

		// Empty pile
		style := styles.EmptyPile
		if isActive {
			style = style.BorderForeground(styles.SelectedBorder).BorderStyle(lipgloss.DoubleBorder())
		}

		// Adjust empty text vertical alignment manually if needed, or rely on style
		return style.Render(emptyText)
	}

	var parts []string

	// Stock
	stockActive := m.game.ActivePile == game.StockPile
	stockSource := m.sourcePileIndex == game.StockPile
	var stockStr string
	if len(m.game.Stock.Cards) > 0 {
		style := styles.FaceDownCard
		if stockSource {
			style = styles.SourceCard.Background(styles.FaceDownBackground).Foreground(styles.FaceDownForeground)
		} else if stockActive {
			style = styles.SelectedCard.Background(styles.FaceDownBackground).Foreground(styles.FaceDownForeground)
		}
		stockStr = style.Render("░░░░░░░░░\n░░░░░░░░░\n░░░░░░░░░\n░░░░░░░░░\n░░░░░░░░░")
	} else {
		style := styles.EmptyPile
		if stockActive {
			style = style.BorderForeground(styles.SelectedBorder).BorderStyle(lipgloss.DoubleBorder())
		}
		stockStr = style.Render("\n\n    ○    \n\n")
	}
	parts = append(parts, stockStr)
	parts = append(parts, "  ") // Space

	// Waste
	parts = append(parts, renderPile(game.WastePile, "\n\n         \n\n", m.game.Waste.Cards))
	parts = append(parts, "    ") // Gap

	// Foundations
	foundations := []string{"♠", "♥", "♦", "♣"}
	for i := 0; i < 4; i++ {
		pileIdx := game.FoundationPile1 + i
		emptyTxt := fmt.Sprintf("\n\n    %s    \n\n", foundations[i])
		parts = append(parts, renderPile(pileIdx, emptyTxt, m.game.Foundations[i].Cards))
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
			// Empty pile
			style := styles.EmptyPile
			if m.game.ActivePile == pileIdx {
				style = style.BorderForeground(styles.SelectedBorder).BorderStyle(lipgloss.DoubleBorder())
			}
			colBuilder.WriteString(style.Render("\n\n    K    \n\n"))
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

// renderCard creates the string for a single card
func (m model) renderCard(c *game.Card, pileIdx, cardIdx int, isOverlap bool) string {
	isActive := m.game.ActivePile == pileIdx && m.game.ActiveCard == cardIdx
	isSource := m.sourcePileIndex == pileIdx && m.sourceCardIndex == cardIdx

	// Determine base style
	var style lipgloss.Style

	if !c.FaceUp {
		style = styles.FaceDownCard
		if isSource {
			style = styles.SourceCard.Background(styles.FaceDownBackground).Foreground(styles.FaceDownForeground)
		} else if isActive {
			style = styles.SelectedCard.Background(styles.FaceDownBackground).Foreground(styles.FaceDownForeground)
		}

		if isOverlap {
			return style.Height(styles.OverlapHeight).Render("░░░░░░░░░")
		}
		return style.Render("░░░░░░░░░\n░░░░░░░░░\n░░░░░░░░░\n░░░░░░░░░\n░░░░░░░░░")
	}

	// Face Up
	// Face Up
	isRed := c.Suit.Color() == "Red"

	if isSource {
		if isRed {
			style = styles.SourceRedCard
		} else {
			style = styles.SourceBlackCard
		}
	} else if isActive {
		if isRed {
			style = styles.SelectedRedCard
		} else {
			style = styles.SelectedBlackCard
		}
	} else if isRed {
		style = styles.RedSuit
	} else {
		style = styles.BlackSuit
	}

	rank := c.Rank.String()
	// Adjust 10 to standard spacing
	rankStr := rank

	suitSym := styles.ArtRegistry[c.Suit.String()]

	var content string

	if isOverlap {
		// Overlap view: just top corners
		// Width 9 (inside 11 width border)
		// "K ♠      "
		content = fmt.Sprintf("%-2s%-1s      ", rankStr, suitSym)
		style = style.Height(styles.OverlapHeight) // Force height
	} else {
		// Full Card View (9x5 content area inside 11x7 border)
		// Top Line: "K ♠      " (Width 9)
		topLine := fmt.Sprintf("%-2s%-1s      ", rankStr, suitSym)

		// Bottom Line: "      ♠ K" (Width 9)
		botLine := fmt.Sprintf("      %-1s%2s", suitSym, rankStr)

		// Center Art (3 lines)
		var centerLines []string

		// Check for Art Registry (K, Q, J, A)
		if art, ok := styles.ArtRegistry[rank]; ok {
			// Split art into lines
			lines := strings.Split(strings.Trim(art, "\n"), "\n")
			// Pad to center
			for _, l := range lines {
				centerLines = append(centerLines, fmt.Sprintf("   %s   ", l)) // 3 padding + 3 art + 3 padding = 9
			}
			// Fill if missing lines
			for len(centerLines) < 3 {
				centerLines = append(centerLines, "         ")
			}
		} else {
			// Number cards - Draw Pips
			val := int(c.Rank)

			// Simple heuristics for pip placement (3 lines in center of 9 width)
			top := "         "
			mid := "         "
			bot := "         "

			switch val {
			case 10:
				top = " ♣  ♣  ♣ "
				mid = "  ♣   ♣  "
				bot = " ♣  ♣  ♣ "
			case 9:
				top = " ♣  ♣  ♣ "
				mid = "    ♣    "
				bot = " ♣  ♣  ♣ "
			case 8:
				top = " ♣  ♣  ♣ "
				mid = "   ♣ ♣   "
				bot = " ♣  ♣  ♣ "
			case 7:
				top = " ♣     ♣ "
				mid = " ♣  ♣  ♣ "
				bot = " ♣     ♣ "
			case 6:
				top = " ♣     ♣ "
				mid = " ♣     ♣ "
				bot = " ♣     ♣ "
			case 5:
				top = " ♣     ♣ "
				mid = "    ♣    "
				bot = " ♣     ♣ "
			case 4:
				top = " ♣     ♣ "
				mid = "         "
				bot = " ♣     ♣ "
			case 3:
				top = "    ♣    "
				mid = "    ♣    "
				bot = "    ♣    "
			case 2:
				top = "    ♣    "
				mid = "         "
				bot = "    ♣    "
			}

			replaceSuit := func(s string) string {
				return strings.ReplaceAll(s, "♣", suitSym)
			}

			if val >= 2 && val <= 10 {
				centerLines = []string{replaceSuit(top), replaceSuit(mid), replaceSuit(bot)}
			} else {
				// Fallback
				centerLines = []string{"         ", fmt.Sprintf("    %s    ", suitSym), "         "}
			}
		}

		content = fmt.Sprintf("%s\n%s\n%s\n%s\n%s",
			topLine,
			centerLines[0],
			centerLines[1],
			centerLines[2],
			botLine,
		)
	}

	return style.Render(content)
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
