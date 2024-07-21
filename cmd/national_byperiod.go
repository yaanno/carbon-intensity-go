/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	s "carbon-intensity/services"
	"fmt"

	"github.com/spf13/cobra"
)

var Future string
var Back string
var Hours uint

var nationalByPeriodCmd = &cobra.Command{
	Use:   "period",
	Short: "Carbon Intensity data by specified period",
	Run: func(cmd *cobra.Command, args []string) {
		pastVal, _ := cmd.Flags().GetBool("past")
		hoursVal, _ := cmd.Flags().GetUint("hours")
		futureVal, _ := cmd.Flags().GetBool("future")

		flagsValues := map[string]interface{}{
			"hours":  hoursVal,
			"past":   pastVal,
			"future": futureVal,
		}

		if cmd.Flag("start-date").Changed {
			startVal := cmd.Flag("start-date").Value.String()

			if validateDate(startVal) {
				flagsValues["from"] = startVal
			} else {
				return
			}
		}
		if cmd.Flag("end-date").Changed {
			endVal := cmd.Flag("end-date").Value.String()
			if validateDate(endVal) {
				flagsValues["to"] = endVal
			} else {
				return
			}
		}
		nationalByPeriodStartDate(flagsValues)

	},
}

func init() {
	nationalCmd.AddCommand(nationalByPeriodCmd)
	nationalByPeriodCmd.Flags().StringVarP(&From, "start-date", "s", "", "Start date in YYYY-MM-DD format")
	nationalByPeriodCmd.Flags().StringVarP(&To, "end-date", "e", "", "End date in YYYY-MM-DD format")
	nationalByPeriodCmd.Flags().BoolP("future", "f", false, "Prediction for the specified period")
	nationalByPeriodCmd.Flags().BoolP("past", "p", false, "History for the specified period")
	nationalByPeriodCmd.Flags().UintVarP(&Hours, "hours", "t", 24, "Period (24 or 48) in hours")
	nationalByPeriodCmd.MarkFlagRequired("start-date")
	nationalByPeriodCmd.MarkFlagsMutuallyExclusive("future", "past")
	nationalByPeriodCmd.MarkFlagsMutuallyExclusive("future", "end-date")
	nationalByPeriodCmd.MarkFlagsMutuallyExclusive("past", "end-date")
}

func nationalByPeriodStartDate(flags map[string]interface{}) {
	request := s.NewIntensityPeriodRequest("intensity")
	request.GetEndpoint(flags)
	_, err := request.Get()
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
	}
	fmt.Println(request.Response.Data)
}
