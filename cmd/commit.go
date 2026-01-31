package cmd

import (
	"github.com/almeidazs/gc/internal/cmds/commit"
	"github.com/spf13/cobra"
)

var branch, message string
var skip, push, emojis, upstream, coauthored bool

var commitCmd = &cobra.Command{
	Use:     "commit",
	Aliases: []string{"cmt"},
	Short:   "Generate ergonomic commit messages with AI",
	RunE: func(cmd *cobra.Command, files []string) error {
		return commit.Commit(commit.CommitOptions{
			Push:        push,
			Emojis:      emojis,
			Branch:      branch,
			Coauthored:  coauthored,
			Files:       files,
			SkipPrompts: skip,
			Message:     message,
			SetUpstream: upstream,
		})
	},
}

func init() {
	commitCmd.Flags().BoolVarP(&skip, "yes", "y", false, "Skip all prompts or not")
	commitCmd.Flags().BoolVarP(&push, "push", "p", false, "Automatically push the commit")
	commitCmd.Flags().BoolVarP(&emojis, "emojis", "e", false, "Use emojis when generating the message")
	commitCmd.Flags().StringVarP(&branch, "branch", "b", "", "A specific branch to push to")
	commitCmd.Flags().StringVarP(&message, "message", "m", "", "Use a specific message while commiting")
	commitCmd.Flags().BoolVarP(&coauthored, "coauthored", "c", false, "Whether the commit is coauthored or not")
	commitCmd.Flags().BoolVarP(&upstream, "upstream", "u", false, "Set the upstream with the branch used in the push")

	rootCmd.AddCommand(commitCmd)
}
