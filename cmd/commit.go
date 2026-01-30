package cmd

import (
	"github.com/almeidazs/gc/internal/cmds/commit"
	"github.com/spf13/cobra"
)

var branch string
var skip, coauthored bool

var commitCmd = &cobra.Command{
	Use:     "commit",
	Aliases: []string{"cmt"},
	Short:   "Generate commit messages and push-it",
	RunE: func(cmd *cobra.Command, files []string) error {
		return commit.Commit(commit.CommitOptions{Branch: branch, Coauthored: coauthored, Files: files, SkipPrompts: skip})
	},
}

func init() {
	commitCmd.Flags().BoolVarP(&skip, "yes", "y", false, "Skip all prompts or not")
	commitCmd.Flags().StringVarP(&branch, "branch", "b", "", "A specific branch to push to")
	commitCmd.Flags().BoolVarP(&coauthored, "coauthored", "c", false, "Whether the commit is coauthored or not")

	rootCmd.AddCommand(commitCmd)
}
