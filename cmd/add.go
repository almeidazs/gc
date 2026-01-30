package cmd

import (
	"github.com/almeidazs/gc/internal/cmds/profiles"
	"github.com/spf13/cobra"
)

var key string

var addCmd = &cobra.Command{
	Use:   "add <name>",
	Args:  cobra.ExactArgs(1),
	Short: "Add a new profile to use",
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		return profiles.Add(name, key)
	},
}

func init() {
	addCmd.Flags().StringVarP(&key, "key", "k", "", "Enter the API key to use")

	rootCmd.AddCommand(addCmd)
}
