package cmd

import (
	"github.com/spf13/cobra"
)

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Task management",
	Long:  `Manage your tasks with ease. Add, list, update, and complete tasks.`,
}

func init() {
	taskCmd.AddCommand(taskAddCmd)
	taskCmd.AddCommand(taskListCmd)
	taskCmd.AddCommand(taskDoneCmd)
	taskCmd.AddCommand(taskDeleteCmd)
}