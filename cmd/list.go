package cmd

import (
	"github.com/almeidazs/gc/internal/cmds/profiles"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all profiles that you created",
	RunE: func(cmd *cobra.Command, args []string) error {
		return profiles.List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
