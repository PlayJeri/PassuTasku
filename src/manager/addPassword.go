package manager

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"

	"github.com/playjeri/passutasku/src/utils"
)

type NewPasswordMessage struct {
	entry PasswordEntry
}

type CancelMessage struct{}

type AddPasswordModel struct {
	form *huh.Form // huh.Form is just a tea.Model
}

var service string
var username string
var password string
var confirm bool

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

				huh.NewConfirm().
					Title("Save the password?").
					Affirmative("Yes").
					Negative("No").
					Value(&confirm),
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
		if !confirm {
			log.Println("not saving")
			service, username, password, confirm = "", "", "", false
			m = NewModel()
			m.Init()
			return m, utils.TeaMessage(CancelMessage{})
		}

		entry := PasswordEntry{
			Service:  service,
			Username: username,
			Password: password,
		}

		// Reset the form
		service, username, password, confirm = "", "", "", false
		m = NewModel()
		m.Init()

		return m, utils.TeaMessage(NewPasswordMessage{entry})
	}

	return m, tea.Batch(cmds...)
}

func (m AddPasswordModel) View() string {
	return m.form.View()
}
