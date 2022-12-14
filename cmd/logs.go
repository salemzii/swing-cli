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
		// if no command is specified after "logs",run logs.all
	},
}

var (
	allCmd = &cobra.Command{
		Use:   "all",
		Short: "fetch all log records",
		Long:  `all command allows you to fetch all your log records`,
		//Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			val, err := logsCmd.Flags().GetBool("json")
			if err != nil {
				log.Printf("error parsing flags %v", err)
			}
			if val == true {
				service.Tojson = true
			}
			service.GetAllRecords(Swingtoken)
		},
	}

	lineCmd = &cobra.Command{
		Use:   "line",
		Short: "fetch all log records with a particular line number",
		Long:  `the line command allows you fetch all log records containing the specified line number`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			lineNum, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalf("Error retrieving line number %v", err)
			}
			val, err := logsCmd.Flags().GetBool("json")
			if err != nil {
				log.Printf("error parsing flags %v", err)
			}
			if val == true {
				service.Tojson = true
			}
			service.GetRecordsWithLineNum(Swingtoken, lineNum)
		},
	}

	functionCmd = &cobra.Command{
		Use:   "function",
		Short: "fetch all log records with a particular function name",
		Long:  "the function command allows you fetch all log records containing the specified function name",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			val, err := logsCmd.Flags().GetBool("json")
			if err != nil {
				log.Printf("error parsing flags %v", err)
			}
			if val == true {
				service.Tojson = true
			}
			service.GetRecordsWithFunction(Swingtoken, args[0])
		},
	}

	levelCmd = &cobra.Command{
		Use:   "level",
		Short: "fetch all log records with a particular log level",
		Long:  "the level command allows you fetch all log records containing the specified log level",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			val, err := logsCmd.Flags().GetBool("json")
			if err != nil {
				log.Printf("error parsing flags %v", err)
			}
			if val == true {
				service.Tojson = true
			}
			service.GetRecordsWithLogLevel(Swingtoken, args[0])
		},
	}

	last15MinCmd = &cobra.Command{
		Use:   "last15",
		Short: "fetch all log records created within 15 minutes",
		Long:  "the last15 command allows you fetch all log records that were ingested within the last 15 minutes",
		//Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			val, err := logsCmd.Flags().GetBool("json")
			if err != nil {
				log.Printf("error parsing flags %v", err)
			}
			if val == true {
				service.Tojson = true
			}
			service.GetRecordsLast15(Swingtoken)
		},
	}

	lastXminCmd = &cobra.Command{
		Use:   "lastx",
		Short: "fetch all log records created within X minutes",
		Long:  "the lastx command allows you fetch all log records that were ingested within the last x minutes",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			lineNum, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatalf("Error retrieving line number %v", err)
			}
			val, err := logsCmd.Flags().GetBool("json")
			if err != nil {
				log.Printf("error parsing flags %v", err)
			}
			if val == true {
				service.Tojson = true
			}
			service.GetRecordsLastX(Swingtoken, lineNum)
		},
	}
)
