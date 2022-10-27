package cmd

import "github.com/spf13/cobra"

var logsCmd = cobra.Command{
	Use:   "logs",
	Short: "logs helps you access all operations on your logs",
	Long:  `logs provides an array of operations that you can use to analyze your logs data`,
}
