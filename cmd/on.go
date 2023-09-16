package cmd

import (
	"github.com/spf13/cobra"
	"github.com/raulcodes/logislim/lib"
)

var brightness int

// onCmd represents the on command
var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Turns on all available Litra Glow lights",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		lib.AllDevicesON(cmd, args, brightness)
	},
}

func init() {
	onCmd.PersistentFlags().IntVarP(&brightness, "brightness", "b", 80, "set brightness (1-100)")
	rootCmd.AddCommand(onCmd)
}
