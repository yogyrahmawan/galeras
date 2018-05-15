package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yogyrahmawan/galeras/websocket"
)

// NewRunWsServer run wsserverr
func NewRunWsServer() *cobra.Command {
	runWsCmd := &cobra.Command{
		Use: "runwsserver",
		Run: runWsServer,
	}

	runWsCmd.Flags().StringP("path", "p", "path", "path of public dir")

	return runWsCmd
}

func runWsServer(cmd *cobra.Command, args []string) {
	path, err := cmd.Flags().GetString("path")
	if err != nil {
		ExitWithError(128, err)
	}

	websocket.Start(path)
}
