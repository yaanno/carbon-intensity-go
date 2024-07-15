/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// statisticsCmd represents the statistics command
var statisticsCmd = &cobra.Command{
	Use:   "statistics",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("statistics called")
	},
}

func init() {
	rootCmd.AddCommand(statisticsCmd)
	statisticsCmd.Flags().StringVar(&From, "start-date", "", "Start date in YYYY-MM-DD format")
	statisticsCmd.Flags().StringVar(&To, "end-date", "", "End date in YYYY-MM-DD format")
	statisticsCmd.MarkFlagRequired("start-date")
	statisticsCmd.MarkFlagRequired("end-date")
	// statisticsCmd.MarkFlagsRequiredTogether("start-date", "end-date")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statisticsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statisticsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
