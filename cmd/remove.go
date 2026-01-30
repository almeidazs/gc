package cmd

import (
	"github.com/almeidazs/gc/internal/cmds/profiles"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove <name>",
	Aliases: []string{"rm"},
	Args:    cobra.ExactArgs(1),
	Short:   "Remove a profile created before",
	RunE: func(cmd *cobra.Command, args []string) error {
		return profiles.Remove(args[0])
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
