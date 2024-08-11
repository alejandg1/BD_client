package tui

import (
	"fmt"
	"github.com/alejandg1/BDclient/internal/config"
	"github.com/alejandg1/BDclient/internal/theme"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type newConnectionForm struct {
	Cursor  int
	Choices []string
  Selected int
	Theme   theme.Theme
}

func NewConnectionForm(colors theme.Theme) newConnectionForm {
	//NOTE: los valores de opciones podria ponerlos en un archivo de configuracion
	choices := []string{" PostgreSQL", " MySQL", " SQLite", "󰈆 Salir"}
	return newConnectionForm{
		Choices: choices,
    Selected: 0,
		Theme:   colors,
	}
}

func (m newConnectionForm) Init() tea.Cmd {
	return nil
}

func (m newConnectionForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "down":
			if m.Cursor < len(m.Choices)-1 {
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

func (m newConnectionForm) View() string {
	s := ""
	title := m.Theme.TitleStyle.Render("Database Engines")
	s += title + "\n\n"
	for i, bd := range m.Choices {

		cursor := ""
		if m.Cursor == i {
			cursor = "❯"
		}

		info := fmt.Sprintf("%s %s",cursor, bd)
		if i == m.Cursor {
			s += m.Theme.SelectedListStyle.Render("  " + info + "\n")
		} else {
			s += m.Theme.ListStyle.Render("  " + info + "\n")
		}
		s += "\n"
	}
	return s
}

func SelectEngineMenu() {
	configs, err := config.GetData[config.Config]()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var theme = theme.NewTheme(configs.Dark)
	var m = NewConnectionForm(theme)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
