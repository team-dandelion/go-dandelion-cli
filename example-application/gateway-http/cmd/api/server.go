package api

import (
	"fmt"
	routing "github.com/gly-hub/fasthttp-routing"
	"example-application/gateway-http/internal/route"
	"github.com/team-dandelion/go-dandelion/application"
	"github.com/team-dandelion/go-dandelion/config"
	"github.com/team-dandelion/go-dandelion/logger"
	"github.com/gly-hub/toolbox/ip"
	"github.com/gly-hub/toolbox/stringx"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/signal"
)

var (
	env      string
	StartCmd = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "gateway-http server -e local",
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
	StartCmd.PersistentFlags().StringVarP(&env, "env", "e", "local", "Env")
}

func setup() {
	// 配置初始化
	config.InitConfig(env)
	// 应用初始化
	application.Init()
	// 路由初始化
	route.InitRoute()
	// 注册头部context链路
	application.RegisterHeaderFunc(HeaderFunc)
}

func HeaderFunc(ctx *routing.Context, data map[string]string) map[string]string {
	// 自定义头部链路。该方法能将需要的参数通过rpc进行传递 TODO

	return data
}

func run() error {
	// 启动http服务
	go func() {
		application.HttpServer().Server()
	}()
	content, _ := ioutil.ReadFile("./static/gateway-http.txt")
	fmt.Println(logger.Green(string(content)))
    fmt.Println(logger.Green("Server run at:"))
    fmt.Printf("-  Local:   http://localhost:%d/ \r\n", application.HttpServer().Port())
    fmt.Printf("-  Network: http://%s:%d/ \r\n", ip.GetLocalHost(), application.HttpServer().Port())
    fmt.Println()
    if config.GetEnv() != "production" {
        fmt.Println(logger.Green("Swagger run at:"))
        fmt.Printf("-  Local:   http://localhost:%d/api/swagger/index.html \r\n", application.HttpServer().Port())
        fmt.Printf("-  Network: http://%s:%d/api/swagger/index.html \r\n", ip.GetLocalHost(), application.HttpServer().Port())
    }

    quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt)
    <-quit
    fmt.Printf("%s Shutdown Server ... \r\n", stringx.GetCurrentTimeStr())
    logger.Info("Server exiting")
    return nil
}
