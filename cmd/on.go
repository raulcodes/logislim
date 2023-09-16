package cmd

import (
	"github.com/spf13/cobra"
	"github.com/raulcodes/logislim/lib"
)

var brightness int
var temperature int

// onCmd represents the on command
var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Turns on all available Litra Glow lights",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		lib.AllDevicesON(cmd, args, brightness, temperature)
	},
}

func init() {
	onCmd.PersistentFlags().IntVarP(&brightness, "brightness", "b", 80, "set brightness (1 - 100)")
	onCmd.PersistentFlags().IntVarP(&temperature, "temperature", "t", 2700, "set temperature (2700K - 6500K)")
	rootCmd.AddCommand(onCmd)
}
