package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yogyrahmawan/galeras/app"
)

// NewNodeCommand create new node command
func NewPullCommand() *cobra.Command {
	pullCmd := &cobra.Command{
		Use:   "pull --image <imagename>",
		Short: "pull --image <imagename>",
		Run:   pullImage,
	}

	pullCmd.Flags().StringP("image", "i", "image", "image name")

	return pullCmd
}

func pullImage(cmd *cobra.Command, args []string) {
	image, err := cmd.Flags().GetString("image")
	if err != nil {
		ExitWithError(128, err)
	}

	if err := app.DockerPull(image); err != nil {
		ExitWithError(1, err)
	}
}
