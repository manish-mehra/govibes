package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var help = []string{
	"Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
	"Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
	"Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
}

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

	var helpStr = ""
	for _, str := range help {
		helpStr += str + "\n"
	}

	var title = titleStyle.MarginBottom(1).MarginLeft(1).Render(" Help ")

	var layout = lipgloss.
		JoinVertical(lipgloss.Left, title, helpStr)

	return ui.Render(layout)
}
