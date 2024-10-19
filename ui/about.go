package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type aboutModel struct {
}

func (m aboutModel) Init() tea.Cmd {
	return nil
}
func (m aboutModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m aboutModel) View() string {

	var ui = lipgloss.
		NewStyle().
		Height(20).
		Align(lipgloss.Left).
		Width(59).
		Background(lipgloss.Color("#1a3c5a"))
		// BorderStyle(lipgloss.NormalBorder())

	return ui.Render("About Page")
}
