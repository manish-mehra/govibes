package main

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

	var render = lipgloss.
		NewStyle().
		Align(lipgloss.Center).
		Background(lipgloss.Color("#099268")).
		Height(10).
		Width(20)

	return fmt.Sprintf(
		render.Render("options"),
	)
}
