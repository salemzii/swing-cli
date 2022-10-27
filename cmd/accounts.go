package cmd

import "github.com/spf13/cobra"

func init() {

}

var accounts = &cobra.Command{
	Use:   "accounts",
	Short: "accounts allows all operations upon a specific account",
	Long: `accounts allows all operations upon a specific account like 
	creating an account	signing into an account, and finding details about an account`,
}

var create = &cobra.Command{
	Use:   "create",
	Short: "create an account on swing",
	Long:  `create allow you to sign up for an account on swing, with your username, email and password`,
}
