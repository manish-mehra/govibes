package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/manish-mehra/govibes/lib"
)

type helpModel struct{}

func (m helpModel) Init() tea.Cmd {
	return nil
}
func (m helpModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m helpModel) View() string {

	var ui = lipgloss.
		NewStyle().MarginTop(1)

	var title = titleStyle.MarginBottom(1).MarginLeft(1).Render(" Help ")

	var layout = lipgloss.
		JoinVertical(lipgloss.Left, title, lib.PrintHelp())

	return ui.Render(layout)
}
