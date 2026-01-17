package styles

import "github.com/charmbracelet/lipgloss"

const (
	CardWidth     = 11 // Width for realistic card appearance (support art)
	CardHeight    = 7  // Height for realistic card appearance
	OverlapHeight = 2  // Visible height when cards overlap
)

// ASCII Art Registry for Face Cards
var ArtRegistry = map[string]string{
	"K": `
 ♔ 
W|M
 | `,
	"Q": `
 ♕ 
(|)
 | `,
	"J": `
 ☺ 
/|\
 | `,
	"A": `
 ♠ 
 ^ 
 | `, // Default Ace (will be specialized),
	"Club":    "♣",
	"Diamond": "♦",
	"Heart":   "♥",
	"Spade":   "♠",
}

// Color palette constants
var (
	// App background - green felt table
	// App background - premium deep felt
	AppBackground = lipgloss.Color("#052d05") // More premium deep green

	// Suit colors - bright for visibility
	RedSuitColor   = lipgloss.Color("#FF0000") // Bright red for Hearts/Diamonds
	BlackSuitColor = lipgloss.Color("#000000") // Black for Spades/Clubs

	// Card colors - white background for contrast
	CardBackground     = lipgloss.Color("#FFFFFF") // White card
	CardForeground     = lipgloss.Color("#000000")
	FaceDownBackground = lipgloss.Color("#1a237e") // Dark blue for face-down
	FaceDownForeground = lipgloss.Color("#5c6bc0") // Lighter blue pattern
	FaceDownPattern    = lipgloss.Color("#3949ab")

	// Border colors
	NormalBorder   = lipgloss.Color("#333333")
	SelectedBorder = lipgloss.Color("#00FFFF") // Cyan for selection
	SourceBorder   = lipgloss.Color("#FFFF00") // Yellow for source card being moved

	// UI colors
	TitleBackground = lipgloss.Color("#1b5e20")
	TitleForeground = lipgloss.Color("#FFFFFF")
	ErrorColor      = lipgloss.Color("#FF5555")
	SuccessColor    = lipgloss.Color("#55FF55")
	HelpTextColor   = lipgloss.Color("#CCCCCC")
	OverlayBg       = lipgloss.Color("#1e1e1e")
)

var (
	// General app style with green felt background
	AppStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Background(AppBackground)

	// Title
	TitleStyle = lipgloss.NewStyle().
			Foreground(TitleForeground).
			Background(TitleBackground).
			Padding(0, 1).
			Bold(true)

	// Status text
	StatusStyle = lipgloss.NewStyle().
			Foreground(HelpTextColor).
			Render

	// Base card style - realistic dimensions
	BaseCard = lipgloss.NewStyle().
			Width(CardWidth).
			Height(CardHeight).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(NormalBorder).
			Background(CardBackground).
			Foreground(CardForeground)

	// Red suit card (Hearts, Diamonds) - red on white
	RedSuit = BaseCard.
		Foreground(RedSuitColor)

	// Black suit card (Spades, Clubs) - black on white
	BlackSuit = BaseCard.
			Foreground(BlackSuitColor)

	// Face-down card with pattern
	FaceDownCard = lipgloss.NewStyle().
			Width(CardWidth).
			Height(CardHeight).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(NormalBorder).
			Background(FaceDownBackground).
			Foreground(FaceDownForeground)

	// Empty pile placeholder
	EmptyPile = lipgloss.NewStyle().
			Width(CardWidth).
			Height(CardHeight).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.RoundedBorder()).
			Foreground(lipgloss.Color("#444444")).
			Background(lipgloss.Color("#0a3d0e")). // Slightly lighter than table
			BorderForeground(lipgloss.Color("#2d6a31"))

	// Selected card generic (cyan border, bold)
	SelectedCard = BaseCard.
			BorderForeground(SelectedBorder).
			BorderStyle(lipgloss.DoubleBorder())

	// Selected Red Card - preserves red color
	SelectedRedCard = RedSuit.
			BorderForeground(SelectedBorder).
			BorderStyle(lipgloss.DoubleBorder())

	// Selected Black Card - preserves black color
	SelectedBlackCard = BlackSuit.
				BorderForeground(SelectedBorder).
				BorderStyle(lipgloss.DoubleBorder())

	// Source card generic (yellow border)
	SourceCard = BaseCard.
			BorderForeground(SourceBorder).
			BorderStyle(lipgloss.ThickBorder())

	// Source Red Card
	SourceRedCard = RedSuit.
			BorderForeground(SourceBorder).
			BorderStyle(lipgloss.ThickBorder())

	// Source Black Card
	SourceBlackCard = BlackSuit.
			BorderForeground(SourceBorder).
			BorderStyle(lipgloss.ThickBorder())

	// Error text style
	ErrorStyle = lipgloss.NewStyle().
			Foreground(ErrorColor).
			Bold(true)

	// Success text style
	SuccessStyle = lipgloss.NewStyle().
			Foreground(SuccessColor).
			Bold(true)

	// Help text style
	HelpStyle = lipgloss.NewStyle().
			Foreground(HelpTextColor)

	// Help overlay style
	HelpOverlay = lipgloss.NewStyle().
			Background(OverlayBg).
			Padding(1, 2).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#555555"))
)
