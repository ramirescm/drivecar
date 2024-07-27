package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/ramirescm/drivecar/pkg/requests"
	"github.com/spf13/cobra"
)

func delete() *cobra.Command {
	var id int32

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete folder",
		Run: func(cmd *cobra.Command, args []string) {
			if id <= 0 {
				log.Println("Folder id is required")
				os.Exit(1)
			}

			path := fmt.Sprintf("/folders/%d", id)
			err := requests.AuthenticatedDelete(path)
			if err != nil {
				log.Printf("%x", err)
				os.Exit(1)
			}

			log.Println("Folder deleted")
		},
	}

	cmd.Flags().Int32VarP(&id, "id", "n", 0, "Folder id")

	return cmd
}
