package main

import (
	"fmt"
	"os"

	"github.com/yogyrahmawan/galeras/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
