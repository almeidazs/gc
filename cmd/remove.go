package cmd

import (
	"github.com/almeidazs/gc/internal/cmds/profiles"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove <name>",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"rm", "delete", "del"},
	Short:   "Remove a profile created before",
	RunE: func(cmd *cobra.Command, args []string) error {
		return profiles.Remove(args[0])
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
