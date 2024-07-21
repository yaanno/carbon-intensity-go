/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	s "carbon-intensity/services"

	"github.com/spf13/cobra"
)

type Flags map[string]interface{}

var regionalCmd = &cobra.Command{
	Use:   "regional",
	Short: "Regional Carbon Intensity data",
	Args:  cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Move this to a subcommand
		current, _ := cmd.Flags().GetBool("current")
		startVal := cmd.Flag("start-date").Value.String()
		endVal := cmd.Flag("end-date").Value.String()
		// past, _ := cmd.Flags().GetBool("past")
		// future, _ := cmd.Flags().GetBool("future")
		// hours, _ := cmd.Flags().GetUint("hours")

		flagsValues := Flags{}

		if startVal != "" || endVal != "" {
			if validateDate(startVal) {
				flagsValues["from"] = startVal
			}
			if validateDate(endVal) {
				flagsValues["to"] = endVal
			}
			regionalByDateCmd(flagsValues)

		} else if current {
			request := s.NewIntensityAllRegionsRequest("regional")
			_, err := request.Get()
			if err != nil {
				fmt.Println(&err)
				return
			}
			fmt.Println(request.Response.Data)
		} else {
			cmd.Help()
		}
	},
	Example: "regional -c",
}

func init() {
	rootCmd.AddCommand(regionalCmd)
	regionalCmd.Flags().BoolP("current", "c", false, "Carbon Intensity data for current half hour")
	regionalCmd.Flags().StringVarP(&From, "start-date", "s", "", "Start date in YYYY-MM-DD format")
	regionalCmd.Flags().StringVarP(&To, "end-date", "e", "", "End date in YYYY-MM-DD format")

	// regionalCmd.Flags().BoolP("future", "f", false, "Prediction for the specified period")
	// regionalCmd.Flags().BoolP("past", "p", false, "History for the specified period")
	// regionalCmd.Flags().UintVarP(&Hours, "hours", "t", 24, "Period (24 or 48) in hours")

	regionalCmd.MarkFlagsMutuallyExclusive("current", "start-date")
	// regionalCmd.MarkFlagsMutuallyExclusive("future", "past")
	regionalCmd.MarkFlagsMutuallyExclusive("current", "end-date")

	regionalCmd.MarkFlagsRequiredTogether("start-date", "end-date")
	// regionalCmd.MarkFlagsRequiredTogether("hours", "future")
}

func regionalByDateCmd(flags Flags) {
	fmt.Println(flags)
	request := s.NewIntensityDateRequest("regional/intensity")
	request.GetEndpoint(flags)
	_, err := (&request).Get()
	if err != nil {
		fmt.Println("Error:")
		fmt.Println(err)
		return
	}

	fmt.Println(&request.Response.Data)
}
