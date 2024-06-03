package cmd

import (
	"log/slog"
	"os"

	"github.com/gkwa/greengossip/core"
	"github.com/spf13/cobra"
)

var markdownDir string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		core.Run(markdownDir, verbose)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	var err error
	runCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "enable verbose logging")
	runCmd.PersistentFlags().StringVar(&markdownDir, "obsidian-vault-path", "", "path to Obsidian vault")
	err = runCmd.MarkPersistentFlagRequired("obsidian-vault-path")
	if err != nil {
		slog.Error("error marking flag as required", "flag", "obsidian-vault-path", "error", err)
		os.Exit(1)
	}
}
