package cmd

import (
	"fmt"
	"os"
	"time"

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
