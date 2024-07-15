/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	r "carbon-intensity/net"

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
			// fmt.Println("Get general data")
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
		// fmt.Println(cmd.Flag("start-date").Changed)

		flagsValues := map[string]any{
			"id":         cmd.Flag("id").Value,
			"start-date": cmd.Flag("start-date").Value.String(),
			"end-date":   cmd.Flag("end-date").Value.String(),
			"postcode":   cmd.Flag("postcode").Value,
			"forecast":   cmd.Flag("forecast").Value,
			"window":     cmd.Flag("next").Value,
		}
		resp, err := r.DoRequest(r.GetEndpoint("regional", args, flagsValues))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(resp))
	},
	Example: "regional england --next 48",
}

func init() {
	rootCmd.AddCommand(regionalCmd)
	regionalCmd.Flags().StringVar(&From, "start-date", "", "Start date in YYYY-MM-DD format")
	regionalCmd.Flags().StringVar(&To, "end-date", "", "End date in YYYY-MM-DD format")
	regionalCmd.Flags().StringVar(&Next, "next", "24", "Forecast for the next hours for GB regions")
	regionalCmd.Flags().BoolVar(&Forecast, "forecast", false, "Forecast for the next hours for GB regions")
	regionalCmd.Flags().StringVar(&Postcode, "postcode", "", "Data for a region specified by postcode")
	regionalCmd.Flags().StringVar(&RegionId, "id", "", "Data for a region specified by region id")
	// regionalCmd.MarkFlagsRequiredTogether("start-date", "end-date")
	regionalCmd.MarkFlagsRequiredTogether("forecast", "next")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// regionalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// regionalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
