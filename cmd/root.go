package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd handle root comand
var RootCmd = &cobra.Command{
	Use: "galeras",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.AddCommand(BuildDockerCmd)
	RootCmd.AddCommand(NewNodeCommand())
	RootCmd.AddCommand(NewMonitorCommand())
	RootCmd.AddCommand(NewRunTestCommand())
	RootCmd.AddCommand(NewPullCommand())
}
