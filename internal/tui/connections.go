package tui

import (
	"fmt"

	"github.com/alejandg1/BDclient/internal/theme"
	"github.com/alejandg1/BDclient/pkg/config"
	tea "github.com/charmbracelet/bubbletea"
)

type Connections struct {
	Connections []config.Connection
	Cursor      int
	Selected    int
	Theme       theme.Theme
}

func NewConnections(list []config.Connection, colors theme.Theme) Connections {
	return Connections{
		Connections: list,
		Cursor:      0,
		Selected:    0,
		Theme:       colors,
	}
}

func (c Connections) Init() tea.Cmd {
	return nil
}

func (m Connections) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "down":
			if m.Cursor < len(m.Connections)-1 {
				m.Cursor++
			}
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "enter":
			m.Selected = m.Cursor
		}
	}

	return m, nil
}

func (m Connections) View() string {
	s := ""
	title := m.Theme.TitleStyle.Render("Connections")
	s += title + "\n\n"
	for i, conn := range m.Connections {
    info := fmt.Sprintf("name: %s database: %s",conn.Name, conn.Database)
		if i == m.Cursor {
			s += m.Theme.SelectedListStyle.Render("  " + info + "\n")
		} else {
			s += m.Theme.ListStyle.Render("  " + info + "\n")
		}
		s += "\n"
	}
	return s
}
