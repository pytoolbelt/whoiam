/*
Copyright © 2024 Jesse Maitland jesse@pytoolbelt.com
*/
package cmd

import (
	"fmt"
	"github.com/pytoolbelt/whoiam/internal"
	"github.com/spf13/cobra"
	"os"
)

func execEntrypoint(cmd *cobra.Command, args []string) {
	accountName, _ := cmd.Flags().GetString("account")

	if accountName == "" {
		fmt.Println("Account Name is required")
		os.Exit(1)
	}

	if len(args) == 0 {
		fmt.Println("No command provided starting subshell. Type 'exit' to return to the parent shell")
	}

	cfgPath, err := internal.NewConfigPath()
	internal.HandelError(err)

	cfg, err := cfgPath.LoadConfig()
	internal.HandelError(err)

	if !cfg.AccountExists(accountName) {
		fmt.Println("Account does not exist")
		os.Exit(1)
	}

	client, err := internal.NewStsClient()
	internal.HandelError(err)

	identity, err := client.GetCallerIdentity()
	internal.HandelError(err)

	err = internal.AssertAccountAsExpected(identity, cfg.Accounts[accountName])
	internal.HandelError(err)
	fmt.Printf("Assert AWS account %s is %s\n", accountName, cfg.Accounts[accountName])

	shell, err := internal.NewSubShell(args...)
	internal.HandelError(err)

	err = shell.Run()
	internal.HandelError(err)
}

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "A brief description of your command",
	Long:  ``,
	Run:   execEntrypoint,
}

func init() {
	rootCmd.AddCommand(execCmd)
	execCmd.Flags().StringP("account", "a", "", "Account Name")
}
