package ui

// TODO: add keyboard listener package
// TODO: highlight current selected item in sound list
// TODO: Add about, help and sound in main model

import (
	"context"
	"log"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/manish-mehra/go-vibes/lib"
)

var paths = lib.GetAudioFilesPath("./audio")

var wg sync.WaitGroup

var cancel context.CancelFunc // holds the cancel function of the previous sound
var ctx context.Context

type model struct {
	header       headerModel
	currentSound currentSoundModel
	options      optionsModel
	about        aboutModel
	sounds       soundsModel
	currentView  string // s, h, a
	width        int
	height       int
}

func initModel() model {
	return model{
		header: headerModel{},
		currentSound: currentSoundModel{
			sound: "No sound selected",
		},
		sounds:      soundsModel{list: load_sounds()},
		currentView: "s",
		options:     optionsModel{},
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
			return m, nil
		case "s":
			m.currentView = "s"
			return m, nil
		default:
			if m.currentView == "s" {
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
					go lib.ListenKeyboardInput(ctx, configPaths.ConfigJson, configPaths.SoundFilePath)

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

	var render = lipgloss.
		NewStyle().
		PaddingLeft(2)

	var header = lipgloss.
		JoinHorizontal(lipgloss.Center, m.header.View(), lipgloss.NewStyle().MarginRight(2).Render(""), m.currentSound.View())

	var aboutLayout = lipgloss.
		JoinVertical(lipgloss.Top, header, m.options.View(), m.about.View())
	if m.currentView == "a" {
		return render.Render(aboutLayout)
	}

	var soundLayout = lipgloss.
		JoinVertical(lipgloss.Top, header, m.options.View(), m.sounds.View())
	if m.currentView == "s" {
		return render.Render(soundLayout)
	}

	return render.Render(header, "unknow view")
}
