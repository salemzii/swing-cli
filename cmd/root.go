package cmd

import (
	"fmt"
	"os"

	"github.com/TwiN/go-color"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Swingtoken string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(logsCmd, accountsCmd)
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if Swingtoken != "" {
		// Use config file from the flag.
		viper.SetConfigFile(Swingtoken)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.SetConfigType("env")
		viper.AddConfigPath(home)
		viper.SetConfigName("swing")
	}

	if err := viper.ReadInConfig(); err != nil {
		//fmt.Println("Can't read config:", err)
		println(color.Colorize(color.Red, "Unable to retrieve your token at the moment; try signing in"))
	}

	tkIntf := viper.Get("TOKEN")
	if tkIntf == nil {
		println(color.Colorize(color.Red, "Unable to retrieve your token at the moment; try signing in. Token is nil"))
		return
	} else {
		Swingtoken = tkIntf.(string)
	}
}

var rootCmd = &cobra.Command{
	Use:   "swing",
	Short: "distributed log as a service with ssdb",
	Long:  `A Fast and Flexible cloud based log as a service software built on SingleStoreDb and jsonRPC`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
