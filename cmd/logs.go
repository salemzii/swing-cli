package cmd

import "github.com/spf13/cobra"

func init() {

	logsCmd.PersistentFlags().Bool("json", false, "Author name for copyright attribution")

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

var all = &cobra.Command{
	Use:   "all",
	Short: "fetch all log records",
	Long:  `all command allows you to fetch all your log records`,
	Run: func(cmd *cobra.Command, args []string) {
		// service.GetAllRecords()
	},
}

var line = &cobra.Command{
	Use:   "line",
	Short: "fetch all log records with a particular line number",
	Long:  `the line command allows you fetch all log records containing the specified line number`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// service.GetRecordsWithLineNum(args[0])
	},
}

var function = &cobra.Command{
	Use:   "function",
	Short: "fetch all log records with a particular function name",
	Long:  "the function command allows you fetch all log records containing the specified function name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// service.GetRecordsWithLineNum(args[0])
	},
}

var level = &cobra.Command{
	Use:   "level",
	Short: "fetch all log records with a particular log level",
	Long:  "the level command allows you fetch all log records containing the specified log level",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// service.GetRecordsWithLogLevel(args[0])
	},
}

var last15Min = &cobra.Command{
	Use:   "last15",
	Short: "fetch all log records created within 15 minutes",
	Long:  "the last15 command allows you fetch all log records that were ingested within the last 15 minutes",
	Run: func(cmd *cobra.Command, args []string) {
		// service.GetRecordsLast15()
	},
}

var lastXmin = &cobra.Command{
	Use:   "lastx",
	Short: "fetch all log records created within X minutes",
	Long:  "the lastx command allows you fetch all log records that were ingested within the last x minutes",
	Run: func(cmd *cobra.Command, args []string) {
		// service.GetRecordsLastX()
	},
}
