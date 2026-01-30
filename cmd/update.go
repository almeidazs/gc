package cmd

import (
	"github.com/almeidazs/gc/internal/cmds/update"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"upgrade"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return update.Update()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
