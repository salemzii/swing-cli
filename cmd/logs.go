package cmd

import (
	"log"
	"strconv"

	"github.com/salemzii/swing-cli/service"
	"github.com/spf13/cobra"
)

func init() {

	logsCmd.AddCommand(allCmd, lineCmd, lastXminCmd, last15MinCmd, levelCmd, functionCmd)
	logsCmd.PersistentFlags().Bool("json", false, "return response in json format")
	allCmd.PersistentFlags().Bool("tail", false, "maintain an open connection for all newly created logs")

}

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "logs helps you access all operations on your logs",
	Long: `logs provides an array of operations that you can use to analyze your logs data,
		if no sub command follows, it returns the all command by default`,
	Run: func(cmd *cobra.Command, args []string) {
		// if no command is specified after "logs", run logs.all
	},
}

var (
	allCmd = &cobra.Command{
		Use:   "all",
		Short: "fetch all log records",
		Long:  `all command allows you to fetch all your log records`,
		//Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			service.GetAllRecords(Swingtoken)
		},
	}

	lineCmd = &cobra.Command{
		Use:   "line",
		Short: "fetch all log records with a particular line number",
		Long:  `the line command allows you fetch all log records containing the specified line number`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			lineNum, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatalf("Error retrieving line number %v", err)
			}
			service.GetRecordsWithLineNum(args[0], lineNum)
		},
	}

	functionCmd = &cobra.Command{
		Use:   "function",
		Short: "fetch all log records with a particular function name",
		Long:  "the function command allows you fetch all log records containing the specified function name",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {

			service.GetRecordsWithFunction(args[0], args[1])
		},
	}

	levelCmd = &cobra.Command{
		Use:   "level",
		Short: "fetch all log records with a particular log level",
		Long:  "the level command allows you fetch all log records containing the specified log level",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			service.GetRecordsWithLogLevel(args[0], args[1])
		},
	}

	last15MinCmd = &cobra.Command{
		Use:   "last15",
		Short: "fetch all log records created within 15 minutes",
		Long:  "the last15 command allows you fetch all log records that were ingested within the last 15 minutes",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			service.GetRecordsLast15(args[0])
		},
	}

	lastXminCmd = &cobra.Command{
		Use:   "lastx",
		Short: "fetch all log records created within X minutes",
		Long:  "the lastx command allows you fetch all log records that were ingested within the last x minutes",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			lineNum, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatalf("Error retrieving line number %v", err)
			}
			service.GetRecordsLastX(args[0], lineNum)
		},
	}
)
