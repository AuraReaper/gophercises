/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/AuraReaper/cli_task_manager/internal"
	"github.com/spf13/cobra"
)

func listCmd(db *internal.DB) *cobra.Command {
	return &cobra.Command{
		Use:  "list",
		Long: `List all of your incomplete tasks`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listAllTasks(db, args)
		},
	}
}

func listAllTasks(db *internal.DB, args []string) error {
	tasks, err := db.ListAllTasks()
	if err != nil {
		return err
	}

	fmt.Println("You have the following tasks:")

	for _, task := range tasks {
		status := " "
		if task.Done {
			status = "x"
		}

		fmt.Printf("%s [%s] %s\n", task.ID, status, task.Description)
	}

	return nil
}
