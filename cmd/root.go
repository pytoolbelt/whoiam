/*
Copyright © 2024 Jesse Maitland jesse@pytoolbelt.com
*/
package cmd

import (
	"github.com/pytoolbelt/whoiam/internal"
	"os"

	"github.com/spf13/cobra"
)

func rootEntrypoint(cmd *cobra.Command, args []string) {
	client, err := internal.NewStsClient()
	internal.HandelError(err)

	identity, err := client.GetCallerIdentity()
	internal.HandelError(err)

	cfgPath, err := internal.NewConfigPath()
	internal.HandelError(err)

	cfg, err := cfgPath.LoadConfig()
	internal.HandelError(err)

	accountName := cfg.GetAccountByNumber(*identity.Account)

	if accountName == "" {
		accountName = "Unknown"
	}

	internal.PrintCallerIdentityTable(identity, accountName)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "whoiam",
	Short: "Check your current AWS IAM Role",
	Long:  ``,
	Run:   rootEntrypoint,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
