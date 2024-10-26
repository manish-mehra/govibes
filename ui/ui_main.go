package ui

// TODO: store default sound in a cache file
// TODO: fix change input/sound bug
// TODO: play sound by args on cli

import (
	"context"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/manish-mehra/go-vibes/lib"
	"log"
	"sync"
)

var paths = lib.GetAudioFilesPath("./audio")

var wg sync.WaitGroup

var cancel context.CancelFunc // holds the cancel function of the previous sound
var ctx context.Context

type model struct {
	header            headerModel
	currentSound      currentSoundModel
	options           optionsModel
	about             aboutModel
	sounds            soundsModel
	help              helpModel
	inputDevices      inputDevicesModel
	alert             string
	keyboardInputPath string
	currentView       string // s, h, a
	width             int
	height            int
}

func initModel() model {

	inputDevLs, err := lib.GetDeviceInfoFromProcBusInputDevices()
	if err != nil {
		log.Fatal(err)
	}

	return model{
		header:       headerModel{},
		currentSound: currentSoundModel{},
		inputDevices: inputDevicesModel{list: load_devices(), paths: inputDevLs, choice: ""},
		sounds:       soundsModel{list: load_sounds()},
		currentView:  "i", // i, s, h
		help:         helpModel{},
		options:      optionsModel{selected: "i"},
	}
}

func Ui_Main() {
	p := tea.NewProgram(initModel())
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
			return m, tea.Batch(
				tea.ClearScreen, // TODO: Not clearning screen
				tea.Quit,
			)
		case "a":
			m.currentView = "a"
			m.options.selected = "a"
			return m, nil
		case "h":
			m.currentView = "h"
			m.options.selected = "h"
		case "s":
			m.currentView = "s"
			m.options.selected = "s"
			return m, nil
		case "i":
			m.currentView = "i"
			m.options.selected = "i"
			return m, nil
		default:
			// BUG: on changing inputDevices, it doesnt' change update the inputdevicepath unless we change views

			if m.currentView == "i" {
				updatedInputDevices, _ := m.inputDevices.Update(msg)
				m.inputDevices = updatedInputDevices.(inputDevicesModel)

				if m.inputDevices.choice != "" {
					for path, devName := range m.inputDevices.paths {
						if m.inputDevices.choice == devName {
							/**	if cancel != nil {
									cancel()
								}
							**/
							m.keyboardInputPath = path
							m.alert = ""
							// m.alert = "Selected" + " " + m.inputDevices.choice
						}
					}

				}
			}
			if m.currentView == "s" {
				if m.keyboardInputPath == "" {
					m.alert = "Please select an input channel first"
					return m, nil
				}

				updatedSounds, _ := m.sounds.Update(msg)
				m.sounds = updatedSounds.(soundsModel) // Reassign the updated soundModel
				if m.sounds.choice != "" {
					m.currentSound.sound = m.sounds.choice
					// get config json & sound file path based on selected sound
					configPaths, err := lib.GetConfigPaths(paths[m.sounds.choice])
					if err != nil {
						panic(err)
					}
					// Cancel previous sound if it's playing
					if cancel != nil {
						cancel()
					}
					ctx, cancel = context.WithCancel(context.Background())
					wg.Add(1)
					go lib.ListenKeyboardInput(ctx, configPaths.ConfigJson, configPaths.SoundFilePath, m.keyboardInputPath)

				}
			}

		}
		return m, nil
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func (m model) View() string {

	var ui = lipgloss.
		NewStyle().
		PaddingLeft(2).
		MarginBottom(2)

	var header = lipgloss.
		JoinHorizontal(lipgloss.Left, m.header.View())

	var alert = lipgloss.
		NewStyle().
		Foreground(lipgloss.Color("#D2042D")).
		Render(m.alert)

	var inputDeviceUI = lipgloss.
		NewStyle().
		Background(lipgloss.Color("#FF69B4")).
		Foreground(lipgloss.Color("#00000")).
		Render(" ", m.inputDevices.choice, " ")
	if m.inputDevices.choice == "" {
		inputDeviceUI = ""
	}

	var footer = lipgloss.
		NewStyle().
		Align(lipgloss.Left).
		Render(inputDeviceUI, m.currentSound.View())

	var aboutLayout = lipgloss.
		JoinVertical(lipgloss.Left, header, m.options.View(), m.about.View(), footer, alert)
	if m.currentView == "a" {
		return ui.Render(aboutLayout)
	}

	var soundLayout = lipgloss.
		JoinVertical(lipgloss.Left, header, m.options.View(), m.sounds.View(), footer, alert)
	if m.currentView == "s" {
		return ui.Render(soundLayout)
	}
	var inputDevicesLayout = lipgloss.JoinVertical(lipgloss.Left, header, m.options.View(), m.inputDevices.View(), footer, alert)
	if m.currentView == "i" {
		return ui.Render(inputDevicesLayout)
	}
	var helpLayout = lipgloss.
		JoinVertical(lipgloss.Left, header, m.options.View(), m.help.View(), footer, alert)

	if m.currentView == "h" {
		return ui.Render(helpLayout)
	}

	return ui.Render(header, "unknow view")
}
