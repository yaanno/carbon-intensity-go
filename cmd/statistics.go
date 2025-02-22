/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	s "carbon-intensity/services"
	"fmt"

	"github.com/spf13/cobra"
)

// statisticsCmd represents the statistics command
var statisticsCmd = &cobra.Command{
	Use:   "statistics",
	Short: "Carbon Intensity statistics up to a month period",
	Run: func(cmd *cobra.Command, args []string) {
		var dateValid bool
		if cmd.Flag("start-date").Changed {
			dateValid = validateDate(cmd.Flag("start-date").Value.String())
			if !dateValid {
				fmt.Println(cmd.UsageString())
				return
			}
		}
		if cmd.Flag("start-date").Changed {
			dateValid = validateDate(cmd.Flag("end-date").Value.String())
			if !dateValid {
				fmt.Println(cmd.UsageString())
				return
			}
		}
		flagsValues := map[string]string{
			"start-date": cmd.Flag("start-date").Value.String(),
			"end-date":   cmd.Flag("end-date").Value.String(),
		}
		request := s.NewIntensityIntervalRequest("intensity")
		(&request).GetEndpoint(flagsValues)
		_, err := request.Get()
		if err != nil {
			fmt.Println("Error:")
			fmt.Println(&err)
			return
		}
		fmt.Println(&request.Response.Data)
	},
}

func init() {
	rootCmd.AddCommand(statisticsCmd)
	statisticsCmd.Flags().StringVarP(&From, "start-date", "s", "", "Start date in YYYY-MM-DD format")
	statisticsCmd.Flags().StringVarP(&To, "end-date", "e", "", "End date in YYYY-MM-DD format")
	statisticsCmd.MarkFlagRequired("start-date")
	statisticsCmd.MarkFlagRequired("end-date")
}
