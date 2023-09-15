package cmd

import (
	"github.com/spf13/cobra"
	"github.com/raulcodes/logislim/lib"
)

// offCmd represents the off command
var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Turns off all available Litra Glow lights",
	Long: ``,
	Run: lib.AllDevicesOFF,
}

func init() {
	rootCmd.AddCommand(offCmd)
}
