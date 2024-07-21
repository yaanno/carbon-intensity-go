/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	s "carbon-intensity/services"
	"fmt"

	"github.com/spf13/cobra"
)

var Date string
var Period string

// natCmd represents the nat command
var nationalByDateCmd = &cobra.Command{
	Use:   "date",
	Short: "Carbon Intensity data for a specific date",
	Run: func(cmd *cobra.Command, args []string) {
		isToday, _ := cmd.Flags().GetBool("today")
		date := cmd.Flag("date")
		period := cmd.Flag("period")

		if isToday {
			nationalTodayCmd()
			return
		}

		if date.Changed {
			dateValid := validateDate(date.Value.String())
			if !dateValid {
				cmd.Usage()
				return
			}
			flagsValues := map[string]string{
				"date":   date.Value.String(),
				"period": period.Value.String(),
			}
			nationalDateWithPeriodCmd(flagsValues)
			return
		} else {
			cmd.Usage()
		}
	},
}

func init() {
	nationalCmd.AddCommand(nationalByDateCmd)
	nationalByDateCmd.Flags().StringVarP(&Date, "date", "d", "", "Data for a specific date in YYYY-MM-DD format")
	nationalByDateCmd.Flags().BoolP("today", "t", false, "Data for today")
	nationalByDateCmd.Flags().StringVarP(&Period, "period", "p", "", "Data for a specific date (YYYY-MM-DD) and period (in half hour settlements: 1-48)")
}

func nationalTodayCmd() {
	request := s.NewIntensityTodayRequest("intensity")
	request.GetEndpoint()
	_, err := request.Get()
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
	}
	fmt.Println(&request.Response.Data)
}

func nationalDateWithPeriodCmd(flags map[string]string) {
	request := s.NewIntensityDateAndPeriodRequest("intensity")
	request.GetEndpoint(flags)
	_, err := (&request).Get()
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(&err)
		return
	}

	fmt.Println(&request.Response.Data)

}
