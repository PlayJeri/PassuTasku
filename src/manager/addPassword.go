package manager

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type NewPasswordMessage struct {
	entry PasswordEntry
}

type AddPasswordModel struct {
	form *huh.Form // huh.Form is just a tea.Model
}

var service string
var username string
var password string

func NewModel() AddPasswordModel {
	return AddPasswordModel{
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Service").
					Value(&service).
					Prompt("?"),

				huh.NewInput().
					Title("Username").
					Value(&username).
					Prompt("?"),

				huh.NewInput().
					Title("password").
					Password(true).
					Value(&password).
					Prompt("?"),
			),
		),
	}
}

func (m AddPasswordModel) Init() tea.Cmd {
	return m.form.Init()
}

func (m AddPasswordModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		// Quit when the form is done.
		return m, func() tea.Msg {
			return NewPasswordMessage{entry: PasswordEntry{
				Service:  service,
				Username: username,
				Password: password,
			}}
		}
	}

	return m, tea.Batch(cmds...)
}

func (m AddPasswordModel) View() string {
	return m.form.View()
}
