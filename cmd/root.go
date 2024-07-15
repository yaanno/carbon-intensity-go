package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	r "carbon-intensity/net"

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
var Past string
var Postcode string
var RegionId string

// Turn forecast reporting on
var Forecast bool

var rootCmd = &cobra.Command{
	Use:   "carbon-intensity",
	Short: "Display Carbon Intensity data",
	// 	Long: `A longer description that spans multiple lines and likely contains
	// examples and usage of using your application. For example:
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	regionalCmd.PersistentFlags().StringVar(&From, "start-date", "", "Start date in YYYY-MM-DD format")
	regionalCmd.PersistentFlags().StringVar(&To, "end-date", "", "End date in YYYY-MM-DD format")
	regionalCmd.PersistentFlags().StringVar(&Next, "next", "24", "Forecast for the next hours for GB regions")
	regionalCmd.PersistentFlags().BoolVar(&Forecast, "forecast", false, "Forecast for the next hours for GB regions")
	regionalCmd.PersistentFlags().StringVar(&Postcode, "postcode", "", "Data for a region specified by postcode")
	regionalCmd.PersistentFlags().StringVar(&RegionId, "id", "", "Data for a region specified by region id")
	regionalCmd.MarkFlagsRequiredTogether("start-date", "end-date")
	regionalCmd.MarkFlagsRequiredTogether("forecast", "next")
	// generationCmd.PersistentFlags().StringVar(&Past, "window", "24", "Get generation mix for the previous specified hours")
	rootCmd.AddCommand(regionalCmd, statCmd, generationCmd)
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

		flagsValues := map[string]any{
			"id":         RegionId,
			"start-date": From,
			"end-date":   To,
			"postcode":   Postcode,
			"forecast":   Forecast,
			"window":     Past,
			"hours":      Next,
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

var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "Get Carbon Intensity data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Echo: " + strings.Join(args, " "))
	},
}

var generationCmd = &cobra.Command{
	Use:   "generation [now | today | past]",
	Short: "Generation Mix for the GB power system",
	ValidArgs: []string{
		"now", "today", "past",
	},
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Echo: " + strings.Join(args, time.Now().Format(time.RFC3339)))
	},
	Example: "generation past --window 24",
}
