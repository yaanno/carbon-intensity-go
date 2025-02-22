/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	s "carbon-intensity/services"
	"fmt"

	"github.com/spf13/cobra"
)

// factorsCmd represents the factors command
var factorsCmd = &cobra.Command{
	Use:   "factors",
	Short: "Get Carbon Intensity factors for each fuel type",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		request := s.NewFactorsRequest("intensity")
		request.GetEndpoint()
		_, err := request.Get()
		if err != nil {
			fmt.Println("Error:")
			fmt.Println(err)
			return
		}
		fmt.Println(request.Response)
	},
}

func init() {
	rootCmd.AddCommand(factorsCmd)
}
