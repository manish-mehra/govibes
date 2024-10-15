package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type optionsModel struct{}

func (m optionsModel) Init() tea.Cmd {
	return nil
}

func (m optionsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	return m, nil
}

func (m optionsModel) View() string {
	var ui = lipgloss.
		NewStyle().
		Height(16).
		Align(lipgloss.Left).
		Width(20).
		BorderStyle(lipgloss.NormalBorder())

	options := fmt.Sprintf("%s \n\n %s \n %s \n %s \n %s", " OPTIONS", "[S] Sound", "[A] About", "[H] Help", "[Q] Quit")

	return ui.Render(options)
}
