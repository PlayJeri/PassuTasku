package main

import (
	"log"
	"os"

	m "github.com/playjeri/passutasku/src/manager"
	"github.com/playjeri/passutasku/src/utils"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	defer utils.ClearTerminal()
	utils.ClearTerminal()

	file := m.OpenLogFile()
	log.SetOutput(file)

	data := m.LoadFile()
	MainModel := m.MainModel{
		State: m.ShowPasswords,
		ShowPasswordsModel: m.ShowPasswordsModel{
			Passwords: data,
		},
		AddPasswordModel: m.NewModel(),
	}

	if len(MainModel.ShowPasswordsModel.Passwords) == 0 {
		os.Exit(1)
		MainModel.ShowPasswordsModel.Passwords = []m.PasswordEntry{
			{Service: "Twitter", Username: "user1", Password: "password1"},
			{Service: "Facebook", Username: "user2", Password: "password2"},
			{Service: "Instagram", Username: "user3", Password: "password3"},
		}
	}

	p := tea.NewProgram(MainModel)
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
	defer file.Close()
}
