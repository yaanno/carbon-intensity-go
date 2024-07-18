/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	s "carbon-intensity/services"

	"github.com/spf13/cobra"
)

// generationCmd represents the generation command
var generationCmd = &cobra.Command{
	Use:   "generation",
	Short: "Generation Mix for the GB power system",
	Long: `Generation Mix for the GB power system.
Contains the following fuel types:
gas, coal, nuclear, biomass, hydro, imports, solar, wind, other.`,
	Run: func(cmd *cobra.Command, args []string) {
		var dateValid bool
		if cmd.Flag("start-date").Changed {
			dateValid = validateDate(cmd.Flag("start-date").Value.String())
			if !dateValid {
				fmt.Println(cmd.UsageString())
				return
			}
		}
		if cmd.Flag("end-date").Changed {
			dateValid = validateDate(cmd.Flag("end-date").Value.String())
			if !dateValid {
				fmt.Println(cmd.UsageString())
				return
			}
		}
		flagsValues := map[string]string{
			"start-date": cmd.Flag("start-date").Value.String(),
			"end-date":   cmd.Flag("end-date").Value.String(),
			"past":       cmd.Flag("past").Value.String(),
		}
		request := s.NewGenerationMixRequest("generation")
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
	rootCmd.AddCommand(generationCmd)
	generationCmd.Flags().StringVarP(&From, "start-date", "s", "", "Start date in YYYY-MM-DD format")
	generationCmd.Flags().StringVarP(&To, "end-date", "e", "", "End date in YYYY-MM-DD format")
	generationCmd.Flags().BoolVarP(&Past, "past", "p", false, "Show data for the past 24 hours")
	generationCmd.MarkFlagRequired("start-date")
}
