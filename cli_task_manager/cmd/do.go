/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/AuraReaper/cli_task_manager/internal"
	"github.com/spf13/cobra"
)

func doCmd(db *internal.DB) *cobra.Command {
	return &cobra.Command{
		Use:  "do",
		Long: `Mark a task on your TODO list as complete`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return doTask(db, args)
		},
	}
}

func doTask(db *internal.DB, args []string) error {
	id := args[0]
	return db.DoTask(id)
}
