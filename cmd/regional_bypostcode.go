/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	s "carbon-intensity/services"
	"fmt"

	"github.com/spf13/cobra"
)

var Postcode string

// regionalByPostCodeCmd represents the regionalByPostCode command
var regionalByPostCodeCmd = &cobra.Command{
	Use:   "postcode",
	Short: "Carbon Intensity data for specified postcode",
	Run: func(cmd *cobra.Command, args []string) {
		current, _ := cmd.Flags().GetBool("current")
		postcode, _ := cmd.Flags().GetString("postcode")
		if postcode == "" && !current {
			cmd.Help()
		} else if current && postcode == "" {
			cmd.Usage()
		} else if current && postcode != "" {
			regionalbyPostcodeCurrent(&postcode)
		} else {
			pastVal, _ := cmd.Flags().GetBool("past")
			futureVal, _ := cmd.Flags().GetBool("future")

			flagsValues := map[string]interface{}{
				"past":   &pastVal,
				"future": &futureVal,
			}
			if cmd.Flag("start-date").Changed {
				startVal := cmd.Flag("start-date").Value.String()

				if validateDate(startVal) {
					flagsValues["from"] = &startVal
				} else {
					return
				}
			}
			if cmd.Flag("end-date").Changed {
				endVal := cmd.Flag("end-date").Value.String()
				if validateDate(endVal) {
					flagsValues["to"] = &endVal
				} else {
					return
				}
			}
			regionalbyPostcodeAndDate(flagsValues)
			// flagsValues := map[string]string{
			// 	"start-date": cmd.Flag("start-date").Value.String(),
			// 	"end-date":   cmd.Flag("end-date").Value.String(),
			// 	"postcode":   cmd.Flag("postcode").Value.String(),
			// 	// "forecast":   cmd.Flag("forecast").Value.String(),
			// 	// "window":     cmd.Flag("next").Value.String(),
			// }
		}
	},
}

func init() {
	regionalCmd.AddCommand(regionalByPostCodeCmd)
	regionalByPostCodeCmd.Flags().BoolP("current", "c", false, "Carbon Intensity data for current half hour for specified postcode")
	regionalByPostCodeCmd.Flags().StringVarP(&Postcode, "postcode", "p", "", "Data for a region specified by postcode")
	regionalByPostCodeCmd.Flags().StringVarP(&From, "start-date", "s", "", "Start date in YYYY-MM-DD format")
	regionalByPostCodeCmd.Flags().StringVarP(&To, "end-date", "e", "", "End date in YYYY-MM-DD format")
	// regionalByPostCodeCmd.MarkFlagRequired("postcode")
}

func regionalbyPostcodeCurrent(postcode *string) {
	request := s.NewIntensityRegionsPostcodeRequest("regional")
	request.GetEndpoint(postcode)
	_, err := request.Get()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(&request.Response.Data)
}

func regionalbyPostcodeAndDate(flags map[string]interface{}) {
	fmt.Println(flags)
}
