package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var ledsFile string
var ledCount int
var treeWidth int
var treeHeight int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "v2",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.xmaspi.yaml)")

	rootCmd.PersistentFlags().Int("port", 3000, "Port to use for GRPC connection")
	rootCmd.PersistentFlags().String("host", "localhost", "Host to use for GRPC connection")

	rootCmd.PersistentFlags().StringVar(&ledsFile, "leds", "leds.json", "")
	rootCmd.PersistentFlags().IntVar(&ledCount, "led-count", 50, "")

	rootCmd.PersistentFlags().IntVar(&treeWidth, "tree-width", 100, "")
	rootCmd.PersistentFlags().IntVar(&treeHeight, "tree-height", 180, "")

	viper.BindPFlag("grpc.port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("grpc.host", rootCmd.PersistentFlags().Lookup("host"))

	viper.BindPFlag("leds.file", rootCmd.PersistentFlags().Lookup("leds"))
	viper.BindPFlag("leds.count", rootCmd.PersistentFlags().Lookup("led-count"))

	viper.BindPFlag("tree.width", rootCmd.PersistentFlags().Lookup("tree-width"))
	viper.BindPFlag("tree.height", rootCmd.PersistentFlags().Lookup("tree-height"))
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

		viper.AddConfigPath(home)
		viper.AddConfigPath("./")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".xmaspi.yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
	}
}
