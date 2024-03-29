package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/team-dandelion/go-dandelion-cli/cmd/project"
	"github.com/team-dandelion/go-dandelion-cli/cmd/server"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "github.com/team-dandelion/go-dandelion-cli",
	Short:        "go-dandelion-cli",
	SilenceUsage: true,
	Long:         "go-dandelion-cli",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(project.StartCmd)
	rootCmd.AddCommand(server.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
