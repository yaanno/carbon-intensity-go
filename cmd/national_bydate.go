/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Date string
var Period string

// natCmd represents the nat command
var nationalByDateCmd = &cobra.Command{
	Use:   "date",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nat called")
		cmd.Usage()
	},
}

func init() {
	nationalCmd.AddCommand(nationalByDateCmd)
	nationalByDateCmd.Flags().StringVarP(&Date, "date", "d", "", "Data for a specific date in YYYY-MM-DD format")
	nationalByDateCmd.Flags().BoolP("today", "t", false, "Data for today")
	nationalByDateCmd.Flags().StringVarP(&Period, "period", "p", "", "Data for specific date (YYYY-MM-DD) and period (half hour settlement: 1-48)")
}
