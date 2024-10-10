package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type wrapperModel struct{}

func (m wrapperModel) Init() tea.Cmd {
	return nil
}

func (m wrapperModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	return m, nil
}

func (m wrapperModel) View() string {

	var render = lipgloss.
		NewStyle().
		Align(lipgloss.Center).
		Background(lipgloss.Color("#e8590c")).
		Height(20).
		Width(30)

	return fmt.Sprintf(
		render.Render("wrapper view"),
	)
}
