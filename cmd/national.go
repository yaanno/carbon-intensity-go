/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// nationalCmd represents the national command
var nationalCmd = &cobra.Command{
	Use:   "national",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func init() {
	rootCmd.AddCommand(nationalCmd)
	nationalCmd.Flags().BoolP("current", "c", false, "Get Carbon Intensity data for current half hour")
}
