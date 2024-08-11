package tui

import (
	"fmt"
	"github.com/alejandg1/BDclient/internal/config"
	"github.com/alejandg1/BDclient/internal/theme"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

type menu struct {
	Cursor  int
	Choices []string
	Theme   theme.Theme
}

func NewMenu(colors theme.Theme) menu {
	choices := []string{"Nueva conexión", "Conexiones existentes", "Salir"}
	return menu{
		Choices: choices,
		Theme:   colors,
	}
}

func (m menu) Init() tea.Cmd {
	return nil
}

func (m menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}
		case "enter", " ":
			switch m.Cursor {
			case 0:
				fmt.Println("Nueva conexión seleccionada")
			case 1:
				fmt.Println("Conexiones existentes seleccionadas")
			case 2:
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m menu) View() string {
	s := ""
	title := m.Theme.TitleStyle.Render("Main Menu\n")
  keys := m.Theme.TitleStyle.Render("[j/k] to navigate, [enter] to select, [q] to quit\n")
  s += title + keys + "\n"
	for i, choice := range m.Choices {
		cursor := SetCursor(m.Cursor, i)

		info := fmt.Sprintf("%s %s", cursor, choice)
		if i == m.Cursor {
			s += m.Theme.SelectedListStyle.Render(info + "\n")
		} else {
			s += m.Theme.ListStyle.Render(info + "\n")
		}
		s += "\n"
	}
	return s
}

func MainMenu() {
	configs, err := config.GetData[config.Config]()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var theme = theme.NewTheme(configs.Dark)
	var m = NewMenu(theme)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
