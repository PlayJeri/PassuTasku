package main

import (
	"fmt"
	"os"

	"github.com/playjeri/passutasku/src/manager"
	"github.com/playjeri/passutasku/src/utils"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	defer utils.ClearTerminal()
	utils.ClearTerminal()
	p := tea.NewProgram(manager.TestModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
