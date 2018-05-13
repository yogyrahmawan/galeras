package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/yogyrahmawan/galeras/app"
)

// BuildDockerCmd is command to build docker image
var BuildDockerCmd = &cobra.Command{
	Use:   "build-docker-image <dockerfilepath> <image_name>",
	Short: "build docker image",
	Run:   buildDockerImage,
}

func buildDockerImage(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		ExitWithError(128, errors.New("command not complete. use build-docker-image <dockerfilepath> <imagename>"))
	}

	err := app.NewImage(args[1], args[0])
	if err != nil {
		ExitWithError(1, errors.New("failed to build docker image, err = "+err.Error()))
	}
}
