package cmd

import (
	"github.com/spf13/cobra"
)

var calendarCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Calendar management",
	Long:  `Manage your calendar events. Add, list, and manage events.`,
}

func init() {
	calendarCmd.AddCommand(calendarAddCmd)
	calendarCmd.AddCommand(calendarListCmd)
	calendarCmd.AddCommand(calendarDeleteCmd)
}