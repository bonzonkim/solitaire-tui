package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/solitaire-tui/solitaire-tui/internal/ui/styles"
)

func (m model) View() string {
	if m.won {
		return styles.AppStyle.Render("You Win!\n\nPress 'q' to quit.\n")
	}
	if m.unwinnable {
		return styles.AppStyle.Render("Unwinnable game\n\nPress 'q' to quit.\n")
	}
	var s strings.Builder
	s.WriteString(styles.TitleStyle.Render("Solitaire TUI"))
	s.WriteString("\n\n")

	if m.invalidMove {
		s.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render("Invalid Move"))
		s.WriteString("\n\n")
	}

	// This is a very basic view, will be improved later
	s.WriteString("Stock: ")
	if len(m.game.Stock.Cards) > 0 {
		s.WriteString("[?]")
	} else {
		s.WriteString("[ ]")
	}
	s.WriteString("\n")

	s.WriteString("Waste: ")
	if len(m.game.Waste.Cards) > 0 {
		lastCard := m.game.Waste.Cards[len(m.game.Waste.Cards)-1]
		s.WriteString(string(lastCard.Rank) + string(lastCard.Suit))
	} else {
		s.WriteString("[ ]")
	}
	s.WriteString("\n\n")

	s.WriteString("Foundation:\n")
	for i, pile := range m.game.Foundation.Piles {
		s.WriteString(string(i + 1))
		s.WriteString(": ")
		if len(pile) > 0 {
			lastCard := pile[len(pile)-1]
			s.WriteString(string(lastCard.Rank) + string(lastCard.Suit))
		} else {
			s.WriteString("[ ]")
		}
		s.WriteString("\n")
	}
	s.WriteString("\n")

	s.WriteString("Tableau:\n")
	for i, col := range m.game.Tableau.Columns {
		s.WriteString(string(i + 1))
		s.WriteString(": ")
		for _, card := range col {
			if card.FaceUp {
				s.WriteString(string(card.Rank) + string(card.Suit) + " ")
			} else {
				s.WriteString("[?] ")
			}
		}
		s.WriteString("\n")
	}

	s.WriteString("\nPress 'q' to quit.\n")

	return styles.AppStyle.Render(s.String())
}
