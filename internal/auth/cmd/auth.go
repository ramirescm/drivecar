package cmd

import (
	"log"
	"os"

	"github.com/ramirescm/drivecar/pkg/requests"
	"github.com/spf13/cobra"
)

func authenticate() *cobra.Command {
	var (
		user     string
		password string
	)
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Authenticate user API",
		Run: func(cmd *cobra.Command, args []string) {
			if user == "" || password == "" {
				log.Println("Username and password are required")
				os.Exit(1)
			}

			err := requests.Auth("/auth", user, password)
			if err != nil {
				log.Printf("%x", err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&user, "user", "u", "", "username")
	cmd.Flags().StringVarP(&user, "password", "p", "", "password")

	return cmd
}
