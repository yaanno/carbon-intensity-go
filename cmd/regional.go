/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	s "carbon-intensity/services"

	"github.com/spf13/cobra"
)

var regionalCmd = &cobra.Command{
	Use:   "regional [region]",
	Short: "Regional Carbon Intensity data",
	Args:  cobra.OnlyValidArgs,
	ValidArgs: []string{
		"scotland", "england", "wales",
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Get general data")
		}
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
			"id":         cmd.Flag("id").Value.String(),
			"start-date": cmd.Flag("start-date").Value.String(),
			"end-date":   cmd.Flag("end-date").Value.String(),
			"postcode":   cmd.Flag("postcode").Value.String(),
			"forecast":   cmd.Flag("forecast").Value.String(),
			"window":     cmd.Flag("next").Value.String(),
		}

		request := s.NewIntensityAllRegionsRequest("regional")
		request.GetEndpoint(args, flagsValues)
		result, err := request.Get()
		if err != nil {
			fmt.Println(err)
			return
		}
		valid := request.Validate(result)
		if !valid {
			return
		}
		err = request.UnMarshal(result)
		if err != nil {
			return
		}
		fmt.Println(request.Response.Data)
	},
	Example: "regional england --next 48",
}

func init() {
	rootCmd.AddCommand(regionalCmd)
	regionalCmd.Flags().StringVarP(&From, "start-date", "s", "", "Start date in YYYY-MM-DD format")
	regionalCmd.Flags().StringVarP(&To, "end-date", "e", "", "End date in YYYY-MM-DD format")
	regionalCmd.Flags().StringVarP(&Next, "next", "n", "24", "Forecast for the next hours for GB regions")
	regionalCmd.Flags().BoolVarP(&Forecast, "forecast", "f", false, "Forecast for the next hours for GB regions")
	regionalCmd.Flags().StringVarP(&Postcode, "postcode", "p", "", "Data for a region specified by postcode")
	regionalCmd.Flags().StringVar(&RegionId, "id", "", "Data for a region specified by region id")
	regionalCmd.MarkFlagsRequiredTogether("start-date", "end-date")
	regionalCmd.MarkFlagsRequiredTogether("forecast", "next")
}

func validateDate(date string) bool {
	_, err := time.Parse(time.DateOnly, date)
	if err != nil {
		fmt.Printf(
			"Error: invalid date: `%v`. Please use the YYYY-MM-DD format\n",
			date,
		)
		return false
	}
	return true
}
