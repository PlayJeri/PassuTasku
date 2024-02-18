package models

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type State int

const (
	ShowPasswords State = iota
	AddPassword
)

type MainModel struct {
	State              State
	ShowPasswordsModel ShowPasswordsModel
	AddPasswordModel   AddPasswordModel
	Width, Height      int
}

// Init method for the model
func (m MainModel) Init() tea.Cmd {
	return tea.Batch(m.ShowPasswordsModel.Init(), m.AddPasswordModel.Init())
}

// Update method for the model
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var juttu = msg
	log.Println(juttu)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "1":
			m.State = ShowPasswords
		case "2":
			m.State = AddPassword
		case "esc", "ctrl+c":
			return m, tea.Quit
		}
	default:
		switch m.State {
		case ShowPasswords:
			var updatedModel tea.Model
			updatedModel, cmd = m.ShowPasswordsModel.Update(msg)
			m.ShowPasswordsModel = updatedModel.(ShowPasswordsModel)
		case AddPassword:
			cmd = nil
			log.Println("AddPassword")
			return m.AddPasswordModel.Update(msg)
			// var updatedModel tea.Model
			// updatedModel, cmd = m.AddPasswordModel.Update(msg)
			// m.AddPasswordModel = updatedModel.(AddPasswordModel)
		}
	}

	return m, cmd
}

// View method for the model
func (m MainModel) View() string {
	switch m.State {
	case AddPassword:
		return m.AddPasswordModel.View()
	case ShowPasswords:
		return m.ShowPasswordsModel.View()
	}
	return "Press 1 or 2 to switch views"
}
