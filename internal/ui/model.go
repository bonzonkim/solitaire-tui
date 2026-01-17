package ui

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/solitaire-tui/solitaire-tui/internal/game"
)

type model struct {
	game *game.Game
	unwinnable bool
	won bool
	invalidMove bool
	invalidMoveTimer int
	width int
	height int
	cursor struct {
		X int
		Y int
	}
}

func NewModel() model {
	m := model{
		game: game.NewGame(),
	}
	m.cursor.X = 0
	m.cursor.Y = 0
	return m
}

func (m model) Init() tea.Cmd {
	return nil
}
