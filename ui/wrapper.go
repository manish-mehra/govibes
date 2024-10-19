package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var s = sound_list()

type wrapperModel struct {
	currentView string
	sounds      soundModel
	about       aboutModel
}

func (m wrapperModel) Init() tea.Cmd {
	return nil
}

func (m wrapperModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Handle the view and pass the keystrokes to the correct model
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit

		case "s":
			if m.currentView == "s" {
				// Update sounds and reassign to m.sounds
				updatedSounds, cmd := m.sounds.Update(msg)
				m.sounds = updatedSounds.(soundModel) // Reassign the updated soundModel
				return m, cmd
			}

			// Initialize sound list and set currentView to "s"
			m.sounds = soundModel{table: s}
			m.currentView = "s"
			return m, nil

		case "a":
			if m.currentView == "a" {
				updatedAbout, cmd := m.about.Update(msg)
				m.about = updatedAbout.(aboutModel)
				return m, cmd
			}
			m.about = aboutModel{}
			m.currentView = "a"
			return m, nil

		case "h":
			m.currentView = "h"
			return m, nil

		default:
			if m.currentView == "s" {
				updatedSounds, cmd := m.sounds.Update(msg)
				m.sounds = updatedSounds.(soundModel)
				return m, cmd
			}
			if m.currentView == "a" {
				updatedAbout, cmd := m.about.Update(msg)
				m.about = updatedAbout.(aboutModel)
				return m, cmd
			}
			return m, nil
		}
	}
	return m, nil
}
func (m wrapperModel) View() string {

	var ui = lipgloss.
		NewStyle().
		Align(lipgloss.Center).
		Height(20).
		Width(59)
		//BorderStyle(lipgloss.NormalBorder())
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
