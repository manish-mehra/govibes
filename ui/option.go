package ui

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
	var ui = lipgloss.
		NewStyle()

	var buttonWrapper = lipgloss.
		NewStyle().
		Background(lipgloss.Color("#00ADD8"))

	// TODO: only highlight the current selected tab
	options := fmt.Sprintf(
		"%s %s  %s %s",
		buttonWrapper.Render("[S]Sounds"),
		"[A]About",
		"[H]Help",
		"[Q]Quit",
	)

	return ui.Render(options)
}
