/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"autogit/actions"

	"github.com/spf13/cobra"
)

// commitMsgCmd represents the commitMsg command
var commitMsgCmd = &cobra.Command{
	Use:   "commitMsg",
	Short: "MACHINE ONLY: git hook for commit-msg. Not for human usage.",
	Run: func(cmd *cobra.Command, args []string) {
		actions.CommmitMsg(args)
	},
}

func init() {
	hookCmd.AddCommand(commitMsgCmd)
}
