package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yogyrahmawan/galeras/app"
)

// NewMonitorCommand create new monitor command
func NewMonitorCommand() *cobra.Command {
	monitorCmd := &cobra.Command{
		Use: "monitor --username <username> --password <password> --node <node>",
		Run: runMonitor,
	}

	monitorCmd.Flags().StringP("username", "u", "username", "usernamedb")
	monitorCmd.Flags().StringP("password", "p", "password", "passworddb")
	monitorCmd.Flags().StringP("node", "n", "node", "node")

	return monitorCmd
}

func runMonitor(cmd *cobra.Command, args []string) {
	username, err := cmd.Flags().GetString("username")
	if err != nil {
		ExitWithError(128, err)
	}

	password, err := cmd.Flags().GetString("password")
	if err != nil {
		ExitWithError(128, err)
	}

	node, err := cmd.Flags().GetString("node")
	if err != nil {
		ExitWithError(128, err)
	}

	app.MonitorNode(username, password, node, "SHOW STATUS LIKE 'wsrep_cluster_size';")
	app.MonitorNode(username, password, node, "SHOW STATUS LIKE 'wsrep_incoming_addresses';")
}
