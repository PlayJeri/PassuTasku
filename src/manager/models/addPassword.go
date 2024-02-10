package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type AddPasswordModel struct {
	form *huh.Form // huh.Form is just a teaFormModel.
}

func NewModel() AddPasswordModel {
	return AddPasswordModel{
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Key("class").
					Options(huh.NewOptions("Warrior", "Mage", "Rogue")...).
					Title("Choose your class"),

				huh.NewSelect[int]().
					Key("level").
					Options(huh.NewOptions(1, 20, 9999)...).
					Title("Choose your level"),
			),
		),
	}
}

func (m AddPasswordModel) Init() tea.Cmd {
	return m.form.Init()
}

func (m AddPasswordModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// ...

	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	return m, cmd
}

func (m AddPasswordModel) View() string {
	if m.form.State == huh.StateCompleted {
		class := m.form.GetString("class")
		level := m.form.GetString("level")
		return fmt.Sprintf("You selected: %s, Lvl. %s", class, level)
	}
	return m.form.View()
}
