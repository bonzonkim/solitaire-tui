package styles

import "github.com/charmbracelet/lipgloss"

var (
	// General.
	AppStyle = lipgloss.NewStyle().Padding(1, 2)

	// Title.
	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	// Status.
	StatusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Render
)
