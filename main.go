package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/lcaproni-pp/goword/cmd/generate"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		color.Red("%s", err)
		os.Exit(1)
	}
}

// NewRootCmd represents the base command when called without any sub commands
func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goword",
		Short: "tool for generating passwords quickly since doing it in 1password is clunky",
	}

	cmd.AddCommand(generate.NewCmd())

	return cmd
}
