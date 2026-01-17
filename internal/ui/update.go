package ui

import (
	"time"

	"github.com/charmbracelet/bubbletea"
)

type invalidMoveMsg struct{}

func clearInvalidMove() tea.Msg {
	return invalidMoveMsg{}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.game.IsWon() {
		m.won = true
		return m, nil
	}
	if m.game.IsUnwinnable() {
		m.unwinnable = true
		return m, nil
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "h":
			if m.cursor.X > 0 {
				m.cursor.X--
			}
		case "l":
			if m.cursor.X < 6 {
				m.cursor.X++
			}
		case "k":
			if m.cursor.Y > 0 {
				m.cursor.Y--
			}
		case "j":
			// This is a simplified check, will be improved later
			m.cursor.Y++
		case "enter":
			// This is a simplified check, will be improved later
			// For now, just trigger the invalid move message
			m.invalidMove = true
			m.invalidMoveTimer = 2
			return m, tea.Tick(time.Second, func(t time.Time) tea.Msg {
				return clearInvalidMove()
			})
		}
	case tea.MouseMsg:
		m.cursor.X = msg.X
		m.cursor.Y = msg.Y
	case invalidMoveMsg:
		m.invalidMoveTimer--
		if m.invalidMoveTimer <= 0 {
			m.invalidMove = false
		}
		return m, tea.Tick(time.Second, func(t time.Time) tea.Msg {
			return clearInvalidMove()
		})
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}
