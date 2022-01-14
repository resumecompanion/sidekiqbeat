package main

import (
	"os"

	"github.com/resumecompanion/sidekiqbeat/cmd"

	_ "github.com/resumecompanion/sidekiqbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
