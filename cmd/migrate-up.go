package cmd

import (
	"github.com/The-EpaG/Go-BasicServer/pkg/bootstrap/migrate"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateUpCmd)
}

var migrateUpCmd = &cobra.Command{
	Use:   "migrate-up",
	Short: "Database migration",
	Long:  "Database migration",
	Run: func(cmd *cobra.Command, args []string) {
		migrateUp()
	},
}

func migrateUp() {
	migrate.Up()
}
