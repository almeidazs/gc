package cmd

import (
	"github.com/almeidazs/gc/internal/cmds/profiles"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch <name>",
	Args:  cobra.ExactArgs(1),
	Short: "Switch to other profile quickly",
	RunE: func(cmd *cobra.Command, args []string) error {
		return profiles.Switch(args[0])
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
}
