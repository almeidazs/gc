package cmd

import (
	"os"

	"github.com/almeidazs/gc/internal/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gc",
	SilenceUsage: true,
	SilenceErrors: false,
	Version: version.Version,
	Short:   "GC is an ergonomic AI-powered commit message generator",
}

func init() {
	rootCmd.SetVersionTemplate("GC {{.Version}}\n")
}

func Exec() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
