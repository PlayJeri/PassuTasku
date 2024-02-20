package manager

import (
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

	switch m.State {
	case ShowPasswords:
		var updatedModel tea.Model
		updatedModel, cmd = m.ShowPasswordsModel.Update(msg)
		m.ShowPasswordsModel = updatedModel.(ShowPasswordsModel)
	case AddPassword:
		var updatedModel tea.Model
		updatedModel, cmd = m.AddPasswordModel.Update(msg)
		m.AddPasswordModel = updatedModel.(AddPasswordModel)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.State = ShowPasswords
		case "a":
			m.State = AddPassword
		case "ctrl+c":
			return m, tea.Quit
		}

	case NewPasswordMessage:
		m.ShowPasswordsModel.Passwords = append(m.ShowPasswordsModel.Passwords, msg.entry)
		SaveFile(m.ShowPasswordsModel.Passwords)
		m.State = ShowPasswords

	case CancelMessage:
		m.State = ShowPasswords
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
