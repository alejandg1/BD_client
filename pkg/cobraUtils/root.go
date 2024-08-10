package cobrautils

import (
	"fmt"
	"os"

	"github.com/alejandg1/BDclient/internal/theme"
	"github.com/alejandg1/BDclient/internal/tui"
	"github.com/alejandg1/BDclient/pkg/config"
  tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "BDcli",
	Short: "BDcli is a CLI tool for database management",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	config.Checkdirs()
	var configs = config.GetConfig()
	var connections = config.GetConnections()
	config.GetHistory()
	var theme = theme.NewTheme(configs.Dark)
	var m = tui.NewConnections(connections, theme)
	program := tea.NewProgram(m)
	if _,err := program.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting program: %v", err)
		os.Exit(1)
	}
}
