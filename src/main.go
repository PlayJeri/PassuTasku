package main

import (
	"fmt"
	"log"
	"os"

	"github.com/playjeri/passutasku/src/manager"
	"github.com/playjeri/passutasku/src/utils"

	tea "github.com/charmbracelet/bubbletea"
	m "github.com/playjeri/passutasku/src/manager/models"
)

func main() {
	defer utils.ClearTerminal()
	utils.ClearTerminal()

	file := manager.OpenLogFile()
	log.SetOutput(file)
	defer file.Close()

	data := manager.LoadFile()
	MainModel := m.MainModel{
		State: m.ShowPasswords,
		ShowPasswordsModel: m.ShowPasswordsModel{
			Passwords: data,
		},
		AddPasswordModel: m.NewModel(),
	}
	defer log.Println(MainModel.ShowPasswordsModel.Passwords)

	defer manager.SaveFile(MainModel.ShowPasswordsModel.Passwords)

	if len(MainModel.ShowPasswordsModel.Passwords) == 0 {
		log.Println("No passwords found in the file")
		MainModel.ShowPasswordsModel = m.TestModel()
	}

	p := tea.NewProgram(MainModel)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
