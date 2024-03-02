package build

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/team-dandelion/go-dandelion-cli/internal/build"
)

var (
	appName  string
	StartCmd = &cobra.Command{
		Use:          "server",
		Short:        "生成服务结构代码",
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
	StartCmd.PersistentFlags().StringVarP(&appName, "name", "n", "example-server", "服务器名称")
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
		return build.Rpc(appName)
	case 2:
		return build.Http(appName)
	default:
		fmt.Println("不支持该类型！")
	}
	return nil
}
