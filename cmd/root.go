package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goalg",
	Short: "goalg is a CLI algorithm caller.",
}

func Execute() error {
	return rootCmd.Execute()
}