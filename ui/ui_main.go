package ui

// A simple program that opens the alternate screen buffer then counts down
// from 5 and then exits.

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	header       headerModel
	currentSound currentSoundModel
	wrapper      wrapperModel
	options      optionsModel
	width        int
	height       int
}

func Ui_Main() {
	p := tea.NewProgram(model{}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		default:
			var updatedWrapper tea.Model
			updatedWrapper, _ = m.wrapper.Update(message)
			m.wrapper = updatedWrapper.(wrapperModel) // Reassign the updated wrapper model
		}
		return m, nil
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m model) View() string {

	var render = lipgloss.
		NewStyle().
		Align(lipgloss.Center)

	var sidebar = lipgloss.
		JoinVertical(lipgloss.Left, m.currentSound.View(), m.options.View())

	var mainContainer = lipgloss.
		JoinHorizontal(lipgloss.Left, sidebar, m.wrapper.View())

	var layout = lipgloss.
		JoinVertical(lipgloss.Left, m.header.View(), mainContainer)

	return fmt.Sprint(
		lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, render.Render(layout)),
	)
}
