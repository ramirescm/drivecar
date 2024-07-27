package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ramirescm/drivecar/internal/folders"
	"github.com/ramirescm/drivecar/pkg/requests"
	"github.com/spf13/cobra"
)

func list() *cobra.Command {
	var id int32

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List folders",
		Run: func(cmd *cobra.Command, args []string) {
			path := "/folders"
			if id > 0 {
				path = fmt.Sprintf("/folders/%d", id)
			}

			data, err := requests.AuthenticatedGet(path)
			if err != nil {
				log.Printf("%x", err)
				os.Exit(1)
			}

			var fc folders.FolderContent
			err = json.Unmarshal(data, &fc)
			if err != nil {
				log.Printf("%x", err)
				os.Exit(1)
			}

			log.Println("= FOLDER ========")
			log.Println(fc.Folder.Name)
			log.Println("=========")
			for _, c := range fc.Content {
				log.Println(c.ID, " - ", c.Type, " - ", c.Name)
			}
		},
	}

	cmd.Flags().Int32VarP(&id, "id", "n", 0, "Folder id")

	return cmd
}
