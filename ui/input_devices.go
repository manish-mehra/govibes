package ui

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/manish-mehra/govibes/lib"
)

type inputDevicesModel struct {
	list   list.Model
	choice string
	paths  map[string]string
}

func (m inputDevicesModel) Init() tea.Cmd {
	return nil
}

func (m inputDevicesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			inputDev := m.list.SelectedItem().(item)
			m.choice = string(inputDev)
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m inputDevicesModel) View() string {
	return "\n" + m.list.View()
}

func load_devices() list.Model {
	items := []list.Item{}

	inputDevLs, err := lib.GetDeviceInfoFromProcBusInputDevices()
	if err != nil {
		log.Fatal("failed to get input devices list", err)
	}

	for _, val := range inputDevLs {
		items = append(items, item(val))
	}

	const defaultWidth = 70

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = " Available input devices "
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return l
}
