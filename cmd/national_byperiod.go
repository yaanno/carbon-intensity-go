/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Future string
var Back string

// natCmd represents the nat command
var nationalByPeriodCmd = &cobra.Command{
	Use:   "period",
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
	nationalCmd.AddCommand(nationalByPeriodCmd)
	nationalByPeriodCmd.Flags().StringVarP(&From, "start-date", "s", "", "Start date in YYYY-MM-DD format")
	nationalByPeriodCmd.Flags().StringVarP(&To, "end-date", "e", "", "End date in YYYY-MM-DD format")
	nationalByPeriodCmd.Flags().StringVarP(&Future, "future", "f", "", "Get data x hrs forwards from specific datetime")
	nationalByPeriodCmd.Flags().StringVarP(&Back, "past", "p", "", "Get data x hrs backwards from specific datetime")
}
