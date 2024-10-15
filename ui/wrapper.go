package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type wrapperModel struct {
	currentView string
	sounds      soundModel
	about       aboutModel
}

func (m wrapperModel) Init() tea.Cmd {
	return nil
}

func (m wrapperModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	fmt.Printf("--> \r wrapper %s", msg)
	// check for current view
	// if already view set, pass down the the event
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		case "s":
			// check for current view
			// if already view set, pass down the the event
			if m.currentView == "s" {
				m.sounds.Update(msg)
				return m, nil
			}
			s := sound_list()
			m.sounds = soundModel{table: s}
			m.currentView = "s"
			return m, nil
		case "a":
			if m.currentView == "a" {
				m.about.Update(msg)
				return m, nil
			}
			m.about = aboutModel{}
			m.currentView = "a"
			return m, nil
		case "h":
			m.currentView = "h"
			return m, nil
		default:
			return m, nil
		}
	}
	return m, nil
}

func (m wrapperModel) View() string {

	var ui = lipgloss.
		NewStyle().
		Align(lipgloss.Center).
		Height(10).
		Width(59).
		BorderStyle(lipgloss.NormalBorder())
	// Background(lipgloss.Color("#008000"))

	switch m.currentView {
	case "s":
		return ui.Render(m.sounds.View())
	case "a":
		return ui.Render(m.about.View())
	default:
		return ui.Render("unknown view")
	}
}
