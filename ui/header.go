package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var asciiTitle = `
█▀▀ █▀█ █ █ █ █▄▄ █▀▀ █▀
█▄█ █▄█ ▀▄▀ █ █▄█ ██▄ ▄█
`

type headerModel struct{}

func (m headerModel) Init() tea.Cmd {
	return nil
}

func (m headerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	return m, nil
}

func (m headerModel) View() string {

	var ui = lipgloss.
		NewStyle().
		Border(lipgloss.NormalBorder())

	return ui.Render(asciiTitle)
}
