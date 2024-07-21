/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	s "carbon-intensity/services"
	"fmt"

	"github.com/spf13/cobra"
)

// nationalCmd represents the national command
var nationalCmd = &cobra.Command{
	Use:   "national",
	Short: "National Carbon Intensity data",
	Run: func(cmd *cobra.Command, args []string) {
		current, _ := cmd.Flags().GetBool("current")
		if !current {
			cmd.Help()
		} else {
			request := s.NewIntensityRecentRequest("intensity")
			request.GetEndpoint()
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
		}
	},
}

func init() {
	rootCmd.AddCommand(nationalCmd)
	nationalCmd.Flags().BoolP("current", "c", false, "Carbon Intensity data for current half hour")
}
