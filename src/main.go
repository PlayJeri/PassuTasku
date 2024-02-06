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

	data := manager.LoadFile()
	model := manager.Model{Passwords: data}
	if len(model.Passwords) == 0 {
		model = manager.TestModel()
	}

	defer manager.SaveFile(model.Passwords)

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
