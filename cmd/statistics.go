/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	r "carbon-intensity/net"
	"fmt"

	"github.com/spf13/cobra"
)

// statisticsCmd represents the statistics command
var statisticsCmd = &cobra.Command{
	Use:   "statistics",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("statistics called")
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
		request := r.NewIntensityIntervalRequest("intensity")
		request.GetEndpoint(args, flagsValues)
		result, err := request.Get()
		if err != nil {
			fmt.Println("Error:")
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
}

func init() {
	rootCmd.AddCommand(statisticsCmd)
	statisticsCmd.Flags().StringVar(&From, "start-date", "", "Start date in YYYY-MM-DD format")
	statisticsCmd.Flags().StringVar(&To, "end-date", "", "End date in YYYY-MM-DD format")
	statisticsCmd.MarkFlagRequired("start-date")
	statisticsCmd.MarkFlagRequired("end-date")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statisticsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statisticsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
