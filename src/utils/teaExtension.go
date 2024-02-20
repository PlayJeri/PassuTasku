package utils

import tea "github.com/charmbracelet/bubbletea"

func TeaMessage(msg tea.Msg) tea.Cmd {
	return func() tea.Msg { return msg }
}
