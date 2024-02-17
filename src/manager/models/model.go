package models

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

func (m ShowPasswordsModel) Init() tea.Cmd {
	return tea.SetWindowTitle("PassuTasku")
}

// Init method for the model
func (m MainModel) Init() tea.Cmd {
	return tea.Batch(m.ShowPasswordsModel.Init(), m.AddPasswordModel.Init())
}

// Update method for the model
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.State {
	case AddPassword:
		return m.AddPasswordModel.Update(msg)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "a":
			m.State = AddPassword

		}
	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width, msg.Height
	}
	return m, nil
}

// View method for the model
func (m MainModel) View() string {
	switch m.State {
	case AddPassword:
		return m.AddPasswordModel.View()
	case ShowPasswords:
		return m.ShowPasswordsModel.View()
	}

	panic("unknown state")
}
