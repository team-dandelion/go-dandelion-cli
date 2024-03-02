package cmd

import (
	"errors"
	"example-project/gateway-http/cmd/api"
	"github.com/team-dandelion/go-dandelion/logger"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "example-project/gateway-http",
	Short: "gateway-http",
	SilenceUsage:true,
	Long: "authorize",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New(logger.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init(){
	rootCmd.AddCommand(api.StartCmd)
}

func Execute(){
	if err := rootCmd.Execute(); err != nil{
		os.Exit(-1)
	}
}
