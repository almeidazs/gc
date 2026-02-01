package cmd

import (
	"github.com/almeidazs/gc/internal/cmds/profiles"
	"github.com/spf13/cobra"
)

var sweepCmd = &cobra.Command{
	Use:   "sweep",
	Short: "Sweep all profiles created",
	RunE: func(cmd *cobra.Command, args []string) error {
		return profiles.Sweep()
	},
}

func init() {
	rootCmd.AddCommand(sweepCmd)
}
