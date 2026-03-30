package cmd

import (
	"whros-cli/internal/task"

	"github.com/spf13/cobra"
)

var taskAddCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		priority, _ := cmd.Flags().GetString("priority")
		tags, _ := cmd.Flags().GetStringSlice("tag")
		dueDate, _ := cmd.Flags().GetString("due")

		t := task.NewTask(args[0], priority, dueDate, tags)
		return task.AddTask(t)
	},
}

var taskListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		all, _ := cmd.Flags().GetBool("all")
		return task.ListTasks(all)
	},
}

var taskDoneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark a task as done",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return task.DoneTask(args[0])
	},
}

var taskDeleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return task.DeleteTask(args[0])
	},
}

func init() {
	taskAddCmd.Flags().StringP("priority", "p", "medium", "Task priority (low, medium, high)")
	taskAddCmd.Flags().StringSliceP("tag", "t", []string{}, "Task tags")
	taskAddCmd.Flags().String("due", "", "Due date (YYYY-MM-DD)")
	taskListCmd.Flags().BoolP("all", "a", false, "Show all tasks including completed")
}