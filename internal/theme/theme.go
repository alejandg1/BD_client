package theme

import (
	"github.com/charmbracelet/lipgloss"
)

type Theme struct {
	TitleStyle         lipgloss.Style
	ListStyle          lipgloss.Style
	SelectedListStyle  lipgloss.Style
	SelectedTitleStyle lipgloss.Style
	BorderStyle        lipgloss.Style
	PromptStyle        lipgloss.Style
	ErrorStyle         lipgloss.Style
}

func NewTheme(darkTheme bool) Theme {
	if darkTheme {
		return Theme{
			TitleStyle:         lipgloss.NewStyle().Foreground(lipgloss.Color("#C0CAF5")).Background(lipgloss.Color("#1A1B26")).Bold(true),
			ListStyle:          lipgloss.NewStyle().Foreground(lipgloss.Color("#A9B1D6")).Background(lipgloss.Color("#1A1B26")),
			SelectedListStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#1A1B26")).Background(lipgloss.Color("#7AA2F7")),
			SelectedTitleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#1A1B26")).Background(lipgloss.Color("#7AA2F7")).Bold(true),
			BorderStyle:        lipgloss.NewStyle().Foreground(lipgloss.Color("#7AA2F7")).Background(lipgloss.Color("#1A1B26")),
			PromptStyle:        lipgloss.NewStyle().Foreground(lipgloss.Color("#7AA2F7")).Background(lipgloss.Color("#1A1B26")),
			ErrorStyle:         lipgloss.NewStyle().Foreground(lipgloss.Color("#F7768E")).Background(lipgloss.Color("#1A1B26")),
		}
	} else {
		return Theme{
			TitleStyle:         lipgloss.NewStyle().Foreground(lipgloss.Color("#4C566A")).Background(lipgloss.Color("#E4E8F0")).Bold(true),
			ListStyle:          lipgloss.NewStyle().Foreground(lipgloss.Color("#7C818C")).Background(lipgloss.Color("#E4E8F0")),
			SelectedListStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("#E4E8F0")).Background(lipgloss.Color("#7AA2F7")),
			SelectedTitleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#E4E8F0")).Background(lipgloss.Color("#7AA2F7")).Bold(true),
			BorderStyle:        lipgloss.NewStyle().Foreground(lipgloss.Color("#7AA2F7")).Background(lipgloss.Color("#E4E8F0")),
			PromptStyle:        lipgloss.NewStyle().Foreground(lipgloss.Color("#7AA2F7")).Background(lipgloss.Color("#E4E8F0")),
			ErrorStyle:         lipgloss.NewStyle().Foreground(lipgloss.Color("#F7768E")).Background(lipgloss.Color("#E4E8F0")),
		}
	}
}
