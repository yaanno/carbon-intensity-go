package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var Verbose bool

// Start date
var From string

// End date
var To string
var Next string
var Past bool
var Postcode string
var RegionId string

// Turn forecast reporting on
var Forecast bool

var rootCmd = &cobra.Command{
	Use:   "carbon-intensity",
	Short: "Display Carbon Intensity data",
	Long: `This CLI is a query and display interface for the Carbon Intensity API for Great Britain 
developed by National Grid. You can find out more about carbon intensity at http://carbonintensity.org.uk.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// regionalCmd.PersistentFlags().StringVar(&From, "start-date", "", "Start date in YYYY-MM-DD format")
	// regionalCmd.PersistentFlags().StringVar(&To, "end-date", "", "End date in YYYY-MM-DD format")
	// regionalCmd.PersistentFlags().StringVar(&Next, "next", "24", "Forecast for the next hours for GB regions")
	// regionalCmd.PersistentFlags().BoolVar(&Forecast, "forecast", false, "Forecast for the next hours for GB regions")
	// regionalCmd.PersistentFlags().StringVar(&Postcode, "postcode", "", "Data for a region specified by postcode")
	// regionalCmd.PersistentFlags().StringVar(&RegionId, "id", "", "Data for a region specified by region id")
	// regionalCmd.MarkFlagsRequiredTogether("start-date", "end-date")
	// regionalCmd.MarkFlagsRequiredTogether("forecast", "next")

	// statsCmd.PersistentFlags().StringVar(&From, "start-date", "", "Start date in YYYY-MM-DD format")
	// statsCmd.PersistentFlags().StringVar(&To, "end-date", "", "End date in YYYY-MM-DD format")
	// statsCmd.MarkFlagRequired("start-date")
	// statsCmd.MarkFlagRequired("end-date")
	// statsCmd.MarkFlagsRequiredTogether("start-date", "end-date")

	// generationCmd.PersistentFlags().StringVar(&Past, "window", "24", "Get generation mix for the previous specified hours")
	// rootCmd.AddCommand(regionalCmd, statsCmd, generationCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".carbon-intensity" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".carbon-intensity")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
