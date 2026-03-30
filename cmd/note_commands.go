package cmd

import (
	"whros-cli/internal/note"

	"github.com/spf13/cobra"
)

var noteAddCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new note",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		content, _ := cmd.Flags().GetString("content")
		tag, _ := cmd.Flags().GetString("tag")

		n := note.NewNote(args[0], content, tag)
		return note.AddNote(n)
	},
}

var noteListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all notes",
	RunE: func(cmd *cobra.Command, args []string) error {
		return note.ListNotes()
	},
}

var noteSearchCmd = &cobra.Command{
	Use:   "search [keyword]",
	Short: "Search notes",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return note.SearchNotes(args[0])
	},
}

var noteDeleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a note",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return note.DeleteNote(args[0])
	},
}

func init() {
	noteAddCmd.Flags().StringP("content", "c", "", "Note content")
	noteAddCmd.Flags().StringP("tag", "t", "", "Note tag")
}