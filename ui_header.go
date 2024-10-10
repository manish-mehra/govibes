package main

import (
	"fmt"

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

	var render = lipgloss.
		NewStyle().
		Align(lipgloss.Center).
		Background(lipgloss.Color("#7D56F4")).
		Height(5).
		Width(50)

	return fmt.Sprintf(
		render.Render(asciiTitle),
	)
}
