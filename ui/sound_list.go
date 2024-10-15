package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type soundModel struct {
	table table.Model
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
			fmt.Printf("\r play %s", row)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m soundModel) View() string {
	m.table.Focus()
	return baseStyle.Render(m.table.View())
}

func sound_list() table.Model {

	columns := []table.Column{
		{Title: "Country", Width: 10},
		{Title: "Population", Width: 25},
	}

	rows := []table.Row{
		{"1", "Tokyo"},
		{"2", "Delhi"},
		{"4", "Dhaka"},
		{"5", "São Paulo"},
		{"6", "Mexico City"},
		{"7", "Cairo"},
		{"8", "Beijing"},
		{"9", "Mumbai"},
		{"10", "Osaka"},
		{"11", "Chongqing"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
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
