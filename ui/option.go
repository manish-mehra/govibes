package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var optionStr = []string{"[i]input devices", "[s]sounds", "[a]about", "[h]help", "[q]quit"}

type optionsModel struct {
	selected string
}

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
		Background(lipgloss.Color("#00ADD8")).
		Foreground(lipgloss.Color("#00000"))

	options := ""
	for _, str := range optionStr {

		if m.selected == string(str[1]) {
			options += buttonWrapper.Render(" ", str, " ") + " "
			continue
		}
		options += str + " "
	}

	return ui.Render(options)
}
