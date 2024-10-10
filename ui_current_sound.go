package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type currentSoundModel struct{}

func (m currentSoundModel) Init() tea.Cmd {
	return nil
}

func (m currentSoundModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	return m, nil
}

func (m currentSoundModel) View() string {

	var render = lipgloss.
		NewStyle().
		Align(lipgloss.Center).
		Background(lipgloss.Color("#c2255c")).
		Height(3).
		Width(20)

	return fmt.Sprintf(
		render.Render("Currently Playing..."),
	)
}
