package cmd

import (
	"whros-cli/internal/calendar"

	"github.com/spf13/cobra"
)

var calendarAddCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new calendar event",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		time, _ := cmd.Flags().GetString("time")
		desc, _ := cmd.Flags().GetString("desc")
		duration, _ := cmd.Flags().GetInt("duration")

		event := calendar.NewEvent(args[0], time, desc, duration)
		return calendar.AddEvent(event)
	},
}

var calendarListCmd = &cobra.Command{
	Use:   "list",
	Short: "List calendar events",
	RunE: func(cmd *cobra.Command, args []string) error {
		date, _ := cmd.Flags().GetString("date")
		return calendar.ListEvents(date)
	},
}

var calendarDeleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a calendar event",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return calendar.DeleteEvent(args[0])
	},
}

func init() {
	calendarAddCmd.Flags().StringP("time", "t", "", "Event time (YYYY-MM-DD HH:MM)")
	calendarAddCmd.Flags().StringP("desc", "d", "", "Event description")
	calendarAddCmd.Flags().IntP("duration", "u", 60, "Duration in minutes")
	calendarListCmd.Flags().StringP("date", "D", "", "Date to list events (YYYY-MM-DD)")
}