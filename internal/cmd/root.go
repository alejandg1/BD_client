package cobrautils

import (
	"fmt"
	"github.com/alejandg1/BDclient/internal/config"
	"github.com/alejandg1/BDclient/internal/tui"
	"github.com/spf13/cobra"
	"os"
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
  // tui.MainMenu()
  tui.ConnectionsMenu()
  // tui.SelectEngineMenu()
}
