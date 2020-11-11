package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runAlg)
}

var runAlg = &cobra.Command{
	Use:   "run",
	Short: "Run an algorithm",
	Run: func(cmd *cobra.Command, args []string) {
		// Get algorithm name from first argument
		alg := args[0]
		args = args[1:]

		// TODO implement function mirroring
		fmt.Printf("Algorithm called: %s(%v)\n", alg, args)
	},
}
