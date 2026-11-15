/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/AuraReaper/cli_task_manager/internal"
	"github.com/spf13/cobra"
)

func addCmd(db *internal.DB) *cobra.Command {
	return &cobra.Command{
		Use: "add",
		// Short: "Add a new task to your TODO list",
		Long: `Add a new task to your TODO list`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return addTask(db, args)
		},
	}
}

func addTask(db *internal.DB, args []string) error {
	description := strings.Join(args, " ")
	task := internal.Task{
		Description: description,
		Done:        false,
	}

	fmt.Printf("Added \"%s\" to your task list", task.Description)

	if err := db.CreateTask(task); err != nil {
		return fmt.Errorf("error in creating a task, %w", err)
	}

	return nil
}
