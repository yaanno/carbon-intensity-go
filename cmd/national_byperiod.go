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
	Short: "Carbon Intensity data by specified period",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nationalByPeriodCmd called")
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
