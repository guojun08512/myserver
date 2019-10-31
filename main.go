package main

import (
	"fmt"
	"os"

	"myserver/pkg/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		if err != cmd.ErrUsage {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error()) // #nosec
			os.Exit(1)
		}
	}
}
