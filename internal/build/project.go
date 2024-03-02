package build

import (
	"errors"
	"fmt"
	"github.com/gly-hub/toolbox/file"
	"os/exec"
	"path"
	"strings"
)

// 构建应用基础架构

// NewProject Create an application
func NewProject(projectName string) {
	pwd, err := file.GetPwd()
	if err != nil {
		return
	}
	appDir := path.Join(pwd, projectName)
	if isExists, _ := file.CheckDirExist(appDir); isExists {
		fmt.Println("The same item already exists in the current path!")
		return
	}

	// Create an application folder
	if err = file.CreateDir(appDir); err != nil {
		fmt.Println("Failed to create the project home directory，the err: ", err)
		return
	}

	// Create mod files
	cmd2 := exec.Command("go", "mod", "init", projectName)
	_ = cmd2.Run()

	// Create a common directory

	// Create the logic-rpc service
	_ = buildRpc(projectName, "logic-rpc", true)

	// Create the logic-http service
	_ = buildRpc(projectName, "logic-rpc", true)
	// Pull the dependency configuration
	cmd3 := exec.Command("go", "mod", "tidy")
	_ = cmd3.Run()
}

func Rpc() error {
	// Get project name

	var serverName string
	fmt.Print("RPC Server Name:")
	if _, err := fmt.Scanln(&serverName); err != nil {
		fmt.Println("An error occurred while reading the input:", err)
		return nil
	}
	defaultConfig, err := EnterBool("Whether to use the default initial configuration?", false)
	if err != nil {
		fmt.Println("An error occurred while reading the input:", err)
		return nil
	}

	return buildRpc(projectName, serverName, defaultConfig)
}

func buildRpc(projectName, serverName string, defaultConfig bool) error {
	var rpcBuilder RpcBuilder
	rpcBuilder.Tools.RpcServer = true
	rpcBuilder.PackageName = fmt.Sprintf("%s/%s", projectName, serverName)
	rpcBuilder.App = projectName
	rpcBuilder.ServerName = serverName

	var err error
	rpcBuilder.Tools.DB, err = EnterBool("Whether to initialize mysql?", defaultConfig)
	if err != nil {
		return err
	}

	rpcBuilder.Tools.Redis, err = EnterBool("Whether to initialize redis?", defaultConfig)
	if err != nil {
		return err
	}

	rpcBuilder.Tools.Logger, err = EnterBool("Whether to initialize logger?", defaultConfig)
	if err != nil {
		return err
	}

	rpcBuilder.Tools.Trace, err = EnterBool("Whether to initialize trace?", defaultConfig)
	if err != nil {
		return err
	}

	rpcBuilder.BuildRpcServer()
	return nil
}

func Http() error {
	var serverName string
	fmt.Print("HTTP Server Name:")
	if _, err := fmt.Scanln(&serverName); err != nil {
		fmt.Println("An error occurred while reading the input:", err)
		return nil
	}

	defaultConfig, err := EnterBool("Whether to use the default initial configuration?", false)
	if err != nil {
		fmt.Println("An error occurred while reading the input:", err)
		return nil
	}

	return buildHttp(projectName, serverName, defaultConfig)
}

func buildHttp(projectName, serverName string, defaultConfig bool) error {
	var httpBuilder HttpBuilder
	httpBuilder.Tools.Http = true
	httpBuilder.Tools.RpcClient = true
	httpBuilder.PackageName = fmt.Sprintf("%s/%s", projectName, serverName)
	httpBuilder.App = projectName
	httpBuilder.ServerName = serverName
	var err error
	httpBuilder.Tools.DB, err = EnterBool("Whether to initialize mysql?", defaultConfig)
	if err != nil {
		return err
	}

	httpBuilder.Tools.Redis, err = EnterBool("Whether to initialize redis?", defaultConfig)
	if err != nil {
		return err
	}

	httpBuilder.Tools.Logger, err = EnterBool("Whether to initialize logger?", defaultConfig)
	if err != nil {
		return err
	}

	httpBuilder.Tools.Trace, err = EnterBool("Whether to initialize trace?", defaultConfig)
	if err != nil {
		return err
	}
	httpBuilder.BuildHttpServer()
	return nil
}

func EnterBool(text string, defaultConfig bool) (bool, error) {
	if defaultConfig {
		return true, nil
	}
	var need string
	fmt.Print(fmt.Sprintf("%s（y/n）:", text))
	if _, err := fmt.Scanln(&need); err != nil {
		fmt.Println()
		return false, errors.New(fmt.Sprintf("An error occurred while reading the input:%s", err))
	}
	need = strings.ToLower(need)
	if need == "y" || need == "yes" {
		return true, nil
	}
	return false, nil
}
