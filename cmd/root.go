package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
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
