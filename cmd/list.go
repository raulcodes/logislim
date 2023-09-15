package cmd

import (
	"github.com/raulcodes/logislim/lib"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Returns a list of compatible Logitech Litra Glow lights",
	Long: ``,
	Run: lib.ListDevices,
}


func init() {
	rootCmd.AddCommand(listCmd)
}
