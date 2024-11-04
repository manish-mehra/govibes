package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var AsciiTitle = `
█▀▀ █▀█ █ █ █ █▄▄ █▀▀ █▀
█▄█ █▄█ ▀▄▀ █ █▄█ ██▄ ▄█
`

func TitleStyle(title string) string {
	return lipgloss.
		NewStyle().
		Background(lipgloss.Color("#50C878")).
		Foreground(lipgloss.Color("#00000")).
		Render(" " + title + " ")
}

func InputDeviceStyle(device string) string {
	return lipgloss.
		NewStyle().
		Background(lipgloss.Color("#FF69B4")).
		Foreground(lipgloss.Color("#00000")).
		Render(" 🖮", device, " ")
}

func SoundStyle(sound string) string {
	return lipgloss.
		NewStyle().
		Align(lipgloss.Left).
		Foreground(lipgloss.Color("#00000")).
		Background(lipgloss.Color("#FFC300 ")).
		Render(" 🎧", sound, " ")
}

func AlertStyle(message string) string {
	return lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("#D2042D")).
		Render(message)
}
