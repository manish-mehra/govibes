package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type currentSoundModel struct {
	sound string
}

func (m currentSoundModel) Init() tea.Cmd {
	return nil
}

func (m currentSoundModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m currentSoundModel) View() string {

	if m.sound == "" {
		return lipgloss.NewStyle().Render("")
	}
	return SoundStyle(m.sound)
}
