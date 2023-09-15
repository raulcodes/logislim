package cmd

import (
	"github.com/spf13/cobra"
	"github.com/raulcodes/logislim/lib"
)

// onCmd represents the on command
var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Turns on all available Litra Glow lights",
	Long: ``,
	Run: lib.AllDevicesON,
}

func init() {
	rootCmd.AddCommand(onCmd)
}
