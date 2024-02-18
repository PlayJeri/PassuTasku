package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type State int

const (
	StateInitial State = iota
	StateOne
	StateTwo
)

type MainModel struct {
	state State
	view1 tea.Model
	view2 tea.Model
}

func (m *MainModel) Init() tea.Cmd {
	return nil
}

func (m *MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "1":
			m.state = StateOne
		case "2":
			m.state = StateTwo
		}
	}
	return m, nil
}

func (m *MainModel) View() string {
	switch m.state {
	case StateOne:
		return m.view1.View()
	case StateTwo:
		return m.view2.View()
	default:
		return "Press 1 or 2 to switch views"
	}
}

type ChildView1 struct{}

func (v *ChildView1) Init() tea.Cmd {
	return nil
}

func (v *ChildView1) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return v, nil
}

func (v *ChildView1) View() string {
	return "Viewing state one"
}

type ChildView2 struct{}

func (v *ChildView2) Init() tea.Cmd {
	return nil
}

func (v *ChildView2) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return v, nil
}

func (v *ChildView2) View() string {
	return "Viewing state two"
}

func main() {
	m := &MainModel{
		state: StateInitial,
		view1: &ChildView1{},
		view2: &ChildView2{},
	}

	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
	}
}
