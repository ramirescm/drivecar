package main

import (
	"log"

	authCmd "github.com/ramirescm/drivecar/internal/auth/cmd"
	foldersCmd "github.com/ramirescm/drivecar/internal/folders/cmd"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{}

func main() {
	authCmd.Register(RootCmd)
	foldersCmd.Register(RootCmd)

	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
