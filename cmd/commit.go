package cmd

import (
	"github.com/almeidazs/gc/internal/cmds/commit"
	"github.com/spf13/cobra"
)

var branch string

var commitCmd = &cobra.Command{
	Use: "commit",
	Aliases: []string{"cmt"},
	Short: "Generate commit messages and push-it",
	RunE: func(cmd *cobra.Command, files []string) error {
		return commit.Commit(branch, files)
	},
}

func init() {
	commitCmd.Flags().StringVarP(&branch, "branch", "b", "", "A specific branch to push to")

	rootCmd.AddCommand(commitCmd)
}
