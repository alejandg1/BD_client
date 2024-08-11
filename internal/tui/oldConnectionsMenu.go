package tui

import (
	"fmt"
	"github.com/alejandg1/BDclient/internal/config"
	"github.com/alejandg1/BDclient/internal/theme"
	tea "github.com/charmbracelet/bubbletea"
	"os"
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
		case "e":
			fmt.Println("Edit connection")
		case "d":
			fmt.Println("Delete connection")
		case "r":
			fmt.Println("return to main menu")
		}
	}

	return m, nil
}

func (m Connections) View() string {
	s := ""
	title := m.Theme.TitleStyle.Render("Connections")
	s += title + "\n"
	for i, conn := range m.Connections {
		var icon string
    icon = SetIcon(conn.Engine)
    cursor := SetCursor(m.Cursor, i)
		info := fmt.Sprintf("%s name: %s %s",cursor, conn.Name, icon)
		if i == m.Cursor {
			s += m.Theme.SelectedListStyle.Render(info + "\n")
		} else {
			s += m.Theme.ListStyle.Render(info + "\n")
		}
		s += "\n"
	}
	return s
}

func ConnectionsMenu() {
	configs, err := config.GetData[config.Config]()
	connec, err := config.GetData[[]config.Connection]()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var theme = theme.NewTheme(configs.Dark)
	var m = NewConnections(connec, theme)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
