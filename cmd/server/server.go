package server

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/team-dandelion/go-dandelion-cli/internal/build"
)

var (
	StartCmd = &cobra.Command{
		Use:          "server",
		Short:        "Generate service structure code",
		Example:      "go-dandelion-cli build -n example-application",
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

}

func setup() {
}

func run() error {
	fmt.Print("Type of service you want to create, enter a number（1-rpc 2-http）:")
	var serverType int
	if _, err := fmt.Scanln(&serverType); err != nil {
		fmt.Println("An error occurred while reading the input:", err)
		return nil
	}
	switch serverType {
	case 1:
		return build.Rpc()
	case 2:
		return build.Http()
	default:
		fmt.Println("This type is not supported!")
	}
	return nil
}
