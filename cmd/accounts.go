package cmd

import "github.com/spf13/cobra"

func init() {
	accountsCmd.AddCommand(create, login)
}

var accountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "accounts allows all operations upon a specific account",
	Long: `accounts allows all operations upon a specific account like 
	creating an account	signing into an account, and finding details about an account`,
}

var create = &cobra.Command{
	Use:   "signup",
	Short: "signup an account on swing",
	Long:  `signup allow you to sign up for an account on swing, with your username, email and password`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		// service.CreateAccount(args)
	},
}

var login = &cobra.Command{
	Use:   "login",
	Short: "login to your swing account",
	Long:  `login allows you log into your swing account, with your email and password`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// service.Login(args)
	},
}

var details = &cobra.Command{
	Use:   "info",
	Short: "get more information about an account",
	Long:  `get detailed information about a logged in account like tokens details, account creation date, number of records logged etc`,
	Run: func(cmd *cobra.Command, args []string) {
		//service.Details()
	},
}
