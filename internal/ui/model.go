package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/solitaire-tui/solitaire-tui/internal/game"
)

// Messages
type clearInvalidMoveMsg struct{}
type clearLastKeyMsg struct{}

// Command timeout
const commandTimeout = 300 * time.Millisecond

// Layout constants
const (
	minHeight = 24
	minWidth  = 80
)

type model struct {
	game     *game.Game
	viewport viewport.Model
	ready    bool

	// Window dimensions
	width  int
	height int

	// Selection state for moving cards
	sourcePileIndex int
	sourceCardIndex int

	// UI state
	showInvalidMove bool
	showHelp        bool
	lastKey         string
	lastKeyTime     time.Time
}

func NewModel() model {
	return model{
		game:            game.NewGame(),
		sourcePileIndex: -1,
		sourceCardIndex: -1,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

// Timer commands
func clearInvalidMoveAfter(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return clearInvalidMoveMsg{}
	})
}

func clearLastKeyAfter(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return clearLastKeyMsg{}
	})
}
