/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Date string
var Period uint

// natCmd represents the nat command
var nationalByDateCmd = &cobra.Command{
	Use:   "date",
	Short: "Carbon Intensity data for a specific date",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nationalByDateCmd called")
		var dateValid bool
		isToday := cmd.Flag("today").Value
		period := cmd.Flag("period").Value
		date := cmd.Flag("date")

		if cmd.Flag("date").Changed {
			dateValid = validateDate(cmd.Flag("date").Value.String())
			if !dateValid {
				fmt.Println(cmd.UsageString())
				return
			}
		}
		fmt.Println(date.Value, isToday, period)
		// flagsValues := map[string]string{
		// 	"start-date": cmd.Flag("start-date").Value.String(),
		// 	"end-date":   cmd.Flag("end-date").Value.String(),
		// }
		cmd.Usage()
	},
}

func init() {
	nationalCmd.AddCommand(nationalByDateCmd)
	nationalByDateCmd.Flags().StringVarP(&Date, "date", "d", "", "Data for a specific date in YYYY-MM-DD format")
	nationalByDateCmd.Flags().BoolP("today", "t", false, "Data for today")
	nationalByDateCmd.Flags().UintVarP(&Period, "period", "p", 1, "Data for a specific date (YYYY-MM-DD) and period (in half hour settlements: 1-48)")
}
