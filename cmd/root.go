package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var Verbose bool
var From string
var To string
var Next int
var Past string

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
	// regionCmd.PersistentFlags().StringVar(&From, "from", time.Now().Format(time.RFC3339), "Start datetime in YYYY-MM-DDThh:mmZ format")
	// regionCmd.PersistentFlags().StringVar(&To, "to", "", "End datetime in in YYYY-MM-DDThh:mmZ format")
	regionCmd.PersistentFlags().IntVar(&Next, "next", 24, "Get Carbon Intensity data for the next specified hours for GB regions")
	generationCmd.PersistentFlags().StringVar(&Past, "window", "24", "Get generation mix for the previous specified hours")

	// rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.carbon-intensity.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(regionCmd, statCmd, generationCmd)
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

var regionCmd = &cobra.Command{
	Use:   "region [region]",
	Short: "Regional Carbon Intensity data",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	ValidArgs: []string{
		"all", "england",
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Echo: " + strings.Join(args, " "))
	},
	Example: "region england --next 48",
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
