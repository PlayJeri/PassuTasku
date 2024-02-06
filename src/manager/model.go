package manager

import (
	tea "github.com/charmbracelet/bubbletea"
)

type passwordEntry struct {
	Service  string
	Username string
	Password string
}

type model struct {
	Passwords []passwordEntry
}

// Init method for the model
func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Passu Tasku")
}

// Update method for the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

// View method for the model
func (m model) View() string {
	s := "Your passwords:\n\n"

	for _, p := range m.Passwords {
		s += p.Service + "\n"
		s += "  Username: " + p.Username + "\n"
		s += "  Password: " + p.Password + "\n\n"
	}

	s += "Press q to quit."

	return s
}
