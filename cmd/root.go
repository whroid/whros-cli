package cmd

import (
	"fmt"
	"os"

	"whros-cli/internal/config"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "whros",
	Short: "whros - Personal affairs management CLI",
	Long: `whros is a personal affairs management CLI tool.
It helps you manage tasks, calendar events, notes, and more.

Quick Start:
  whros task add "Complete the project" --priority high
  whros task list
  whros calendar add "Meeting" --time "2024-01-15 14:00"
  whros note add "Quick idea" --content "Something important"

Use "whros help" for more information.`,
	SilenceUsage: true,
}

func Execute() {
	config.InitConfig()
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(taskCmd)
	RootCmd.AddCommand(calendarCmd)
	RootCmd.AddCommand(noteCmd)
	RootCmd.AddCommand(configCmd)
}