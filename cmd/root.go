package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	verbose bool
	config  string
	listen  string
)

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Debug logs")
	RootCmd.Flags().StringVarP(&config, "config", "c", "", "path and name of configuration file")
	RootCmd.Flags().StringVarP(&listen, "listen", "l", "127.0.0.1:8080", "API listenning address")
	viper.BindPFlags(RootCmd.Flags())
}

// RootCmd launch the API.
var RootCmd = &cobra.Command{
	Use:   "cataclysm",
	Short: "Cataclysm aims to be the simpliest Torrent client with an API",
	Long:  "",
	Run:   run,
}

func initConfig() {

	// Debug logs
	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	// Bind environment variables
	viper.SetEnvPrefix("cataclysm")
	viper.AutomaticEnv()

	// Set config search path
	viper.AddConfigPath("/etc/cataclysm/")
	viper.AddConfigPath("$HOME/.cataclysm")
	viper.AddConfigPath(".")

	// Load config
	viper.SetConfigName("cataclysm.config")
	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Debug("No config file found")
		} else {
			log.Panicf("Fatal error in config file: %v \n", err)
		}
	}

	// Load user defined config
	if config != "" {
		viper.SetConfigFile(config)
		err := viper.ReadInConfig()
		if err != nil {
			log.Panicf("Fatal error in config file: %v \n", err)
		}
	}
}

func run(cmd *cobra.Command, args []string) {

	/*err := core.APIStart(listen)
	if err != nil {
		log.panic(err)
	}*/
}
