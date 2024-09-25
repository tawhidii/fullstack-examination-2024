/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zuu-development/fullstack-examination-2024/internal/db"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	Run: func(cmd *cobra.Command, args []string) {
		dbInstance, err := db.New(cfg.SQLite.DBFilename)
		if err != nil {
			fmt.Println("Error connecting to database")
			return
		}
		if err := db.Migrate(dbInstance); err != nil {
			fmt.Println("Error migrating database")
			return
		}
		fmt.Println("Migration completed. SQLite.DBFilename: ", cfg.SQLite.DBFilename)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}