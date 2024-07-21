/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	s "carbon-intensity/services"

	"github.com/spf13/cobra"
)

var regionalCmd = &cobra.Command{
	Use:   "regional [region]",
	Short: "Regional Carbon Intensity data",
	Args:  cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		current, _ := cmd.Flags().GetBool("current")
		if !current {
			cmd.Help()
		} else {
			request := s.NewIntensityAllRegionsRequest("regional")
			request.GetEndpoint()
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
		}
		// if len(args) == 0 {
		// 	fmt.Println("Get general data")
		// }
		// var dateValid bool
		// if cmd.Flag("start-date").Changed {
		// 	dateValid = validateDate(cmd.Flag("start-date").Value.String())
		// 	if !dateValid {
		// 		fmt.Println(cmd.UsageString())
		// 		return
		// 	}
		// }
		// if cmd.Flag("start-date").Changed {
		// 	dateValid = validateDate(cmd.Flag("end-date").Value.String())
		// 	if !dateValid {
		// 		fmt.Println(cmd.UsageString())
		// 		return
		// 	}
		// }

		// flagsValues := map[string]string{
		// 	"id":         cmd.Flag("id").Value.String(),
		// 	"start-date": cmd.Flag("start-date").Value.String(),
		// 	"end-date":   cmd.Flag("end-date").Value.String(),
		// 	"postcode":   cmd.Flag("postcode").Value.String(),
		// 	"forecast":   cmd.Flag("forecast").Value.String(),
		// 	"window":     cmd.Flag("next").Value.String(),
		// }

	},
	Example: "regional -c",
}

func init() {
	rootCmd.AddCommand(regionalCmd)
	regionalCmd.Flags().BoolP("current", "c", false, "Carbon Intensity data for current half hour")

	// regionalCmd.Flags().StringVarP(&From, "start-date", "s", "", "Start date in YYYY-MM-DD format")
	// regionalCmd.Flags().StringVarP(&To, "end-date", "e", "", "End date in YYYY-MM-DD format")
	// regionalCmd.Flags().StringVarP(&Next, "next", "n", "24", "Forecast for the next hours for GB regions")
	// regionalCmd.Flags().BoolVarP(&Forecast, "forecast", "f", false, "Forecast for the next hours for GB regions")

	// regionalCmd.Flags().StringVar(&RegionId, "id", "", "Data for a region specified by region id")
	// regionalCmd.MarkFlagsRequiredTogether("start-date", "end-date")
	// regionalCmd.MarkFlagsRequiredTogether("forecast", "next")
}
