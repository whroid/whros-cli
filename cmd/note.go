package cmd

import (
	"github.com/spf13/cobra"
)

var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "Note management",
	Long:  `Manage your notes and memos. Quick capture and search.`,
}

func init() {
	noteCmd.AddCommand(noteAddCmd)
	noteCmd.AddCommand(noteListCmd)
	noteCmd.AddCommand(noteSearchCmd)
	noteCmd.AddCommand(noteDeleteCmd)
}