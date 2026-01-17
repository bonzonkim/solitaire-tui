package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/term"
	"github.com/solitaire-tui/solitaire-tui/internal/ui"
)

func main() {
	// Initialize term
	physicalWidth, physicalHeight, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		log.Fatal(err)
	}

	if physicalWidth < 80 || physicalHeight < 24 {
		fmt.Println("Terminal is too small to run this application.")
		fmt.Println("Please resize your terminal to at least 80x24.")
		os.Exit(1)
	}

	p := tea.NewProgram(ui.NewModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
