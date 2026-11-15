/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/AuraReaper/cli_task_manager/internal"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a CLI for managing your TODOs.",
	Long: `
task is a CLI for managing your TODOs.

Usage:
	task [command]

Available Commands:
	add         Add a new task to your TODO list
	do          Mark a task on your TODO list as complete
	list        List all of your incomplete tasks

Use "task [command] --help" for more information about a command.
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(db *internal.DB) error {
	rootCmd.AddCommand(addCmd(db))
	rootCmd.AddCommand(listCmd(db))
	rootCmd.AddCommand(doCmd(db))

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	return err
}
