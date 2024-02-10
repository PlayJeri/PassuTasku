package main

import (
	"fmt"
	"os"

	"github.com/playjeri/passutasku/src/manager"
	"github.com/playjeri/passutasku/src/utils"

	tea "github.com/charmbracelet/bubbletea"
	m "github.com/playjeri/passutasku/src/manager/models"
)

func main() {
	defer utils.ClearTerminal()
	utils.ClearTerminal()

	data := manager.LoadFile()
	model := m.MainModel{
		State: m.ShowPasswords,
		ShowPasswordsModel: m.ShowPasswordsModel{
			Passwords: data,
		},
		AddPasswordModel: m.NewModel(),
	}

	if model.ShowPasswordsModel.Passwords == nil {
		model.ShowPasswordsModel = m.TestModel()
	}

	defer manager.SaveFile(model.ShowPasswordsModel.Passwords)

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
