package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ramirescm/drivecar/internal/folders"
	"github.com/ramirescm/drivecar/pkg/requests"
	"github.com/spf13/cobra"
)

func update() *cobra.Command {
	var name string
	var id int32

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update folder name",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" || id <= 0 {
				log.Println("Folder name and id are required")
				os.Exit(1)
			}

			folder := folders.Folder{Name: name}

			var body bytes.Buffer
			err := json.NewEncoder(&body).Encode(folder)
			if err != nil {
				log.Printf("%x", err)
				os.Exit(1)
			}

			path := fmt.Sprintf("/folders/%d", id)
			_, err = requests.AuthenticatedPut(path, &body)
			if err != nil {
				log.Printf("%x", err)
				os.Exit(1)
			}

			log.Println("Folder updated")
		},
	}

	cmd.Flags().Int32VarP(&id, "id", "n", 0, "Folder id")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Folder name")

	return cmd
}
