package cmd

import "github.com/spf13/cobra"

func init() {

}

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "logs helps you access all operations on your logs",
	Long:  `logs provides an array of operations that you can use to analyze your logs data`,
	Run: func(cmd *cobra.Command, args []string) {
		// if no command is specified after "logs", run logs.all
	},
}
