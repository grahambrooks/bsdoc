package main

import (
	"fmt"
	"os"

	"github.com/grahambrooks/bsdoc/internal/commands"
	"github.com/spf13/cobra"
)

var version = "dev"

func main() {
	rootCmd := &cobra.Command{
		Use:   "bsdoc",
		Short: "A CLI tool for generating detailed backstage documentation",
		Long:  `bsdoc uses GitHub Copilot to generate comprehensive backstage documentation for your projects.`,
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("bsdoc v%s\n", version)
		},
	}

	generateCmd := commands.NewGenerateCommand()

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(generateCmd)

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
