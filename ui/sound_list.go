package ui

import (
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/manish-mehra/go-vibes/lib"
)

type soundModel struct {
	table         table.Model
	selectedSound string
}

func (m soundModel) Init() tea.Cmd {
	return nil
}

func (m soundModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.table.Focus()
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			row := m.table.SelectedRow()[1]
			m.selectedSound = row
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m soundModel) View() string {
	m.table.Focus()
	return m.table.View()
}

func sound_list() table.Model {

	columns := []table.Column{
		{Title: "Id", Width: 10},
		{Title: "Sounds", Width: 20},
	}

	rows := []table.Row{}
	sounds := lib.GetAudioFilesPath("./audio")
	index := 1
	for key := range sounds {
		rows = append(rows, table.Row{strconv.Itoa(index), key})
		index++
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(20),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.HiddenBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return t
}
