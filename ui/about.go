package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var about = []string{
	"Govibes - an unnecessary rewrite of mechvibes.com (mechanical keyboard \nsound simulator) disguised as a cli tool",
}

type aboutModel struct{}

func (m aboutModel) Init() tea.Cmd {
	return nil
}
func (m aboutModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m aboutModel) View() string {

	var ui = lipgloss.
		NewStyle().MarginTop(1)

	var aboutStr = ""
	for _, str := range about {
		aboutStr += str + "\n"
	}

	var title = titleStyle.MarginBottom(1).MarginLeft(1).Render(" About ")

	var layout = lipgloss.
		JoinVertical(lipgloss.Left, title, aboutStr)

	return ui.Render(layout)
}
