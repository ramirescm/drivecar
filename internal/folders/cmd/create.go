package cmd

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	"github.com/ramirescm/drivecar/internal/folders"
	"github.com/ramirescm/drivecar/pkg/requests"
	"github.com/spf13/cobra"
)

func create() *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create new folder",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" {
				log.Println("Folder name is required")
				os.Exit(1)
			}

			folder := folders.Folder{Name: name}
			var body bytes.Buffer
			err := json.NewEncoder(&body).Encode(folder)
			if err != nil {
				log.Printf("%x", err)
				os.Exit(1)
			}

			_, err = requests.AuthenticatedPost("/folders", &body)
			if err != nil {
				log.Printf("%x", err)
				os.Exit(1)
			}

			log.Println("Folder created")
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Folder name")

	return cmd
}
