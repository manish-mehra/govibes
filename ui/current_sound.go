package ui

import (
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

	var ui = lipgloss.
		NewStyle().
		Align(lipgloss.Center).
		Height(2).
		Width(20).
		BorderStyle(lipgloss.NormalBorder())

	return ui.Render("cherrymx-black-abs")

}
