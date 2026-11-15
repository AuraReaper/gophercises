/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/AuraReaper/cli_task_manager/cmd"
	"github.com/AuraReaper/cli_task_manager/internal"
)

const (
	dbPath     = "tasks.db"
	bucketName = "tasks"
)

func main() {
	conn, err := internal.OpenDB(dbPath, bucketName)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err = cmd.Execute(conn); err != nil {
		log.Fatal(err)
	}
}
