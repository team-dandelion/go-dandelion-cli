package application

import (
	"github.com/spf13/cobra"
	"github.com/team-dandelion/go-dandelion-cli/internal/build"
)

var (
	appName  string
	StartCmd = &cobra.Command{
		Use:          "create",
		Short:        "Create a go-dandelion project",
		Example:      "go-dandelion-cli create -n example-application",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&appName, "name", "n", "example-application", "Project name")
}

func setup() {
}

func run() error {
	build.BuildApplication(appName)
	return nil
}
