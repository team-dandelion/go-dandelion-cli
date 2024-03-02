package build

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/gly-hub/toolbox/file"
	"log"
	"os"
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
	fmt.Println("Current path:", pwd)
	appDir := path.Join(pwd, projectName)
	if isExists, _ := file.CheckDirExist(appDir); isExists {
		fmt.Println("The same item already exists in the current path!")
		return
	}

	// Create an application folder
	fmt.Println("Creating a project home directory:", projectName)
	if err = file.CreateDir(appDir); err != nil {
		fmt.Println("Failed to create the project home directory，the err: ", err)
		return
	}

	// Create mod files
	fmt.Println("Creating mod files")
	cmd2 := exec.Command("go", "mod", "init", projectName)
	cmd2.Dir = appDir
	_ = cmd2.Run()

	// Create a common directory
	fmt.Println("Creating a common directory")
	if err = file.CreateDir(path.Join(appDir, "common")); err != nil {
		fmt.Println("Failed to create the common directory，the err: ", err)
		return
	}

	// Create the logic-rpc service
	fmt.Println("Creating the logic-rpc service.（logic service）")
	_ = buildRpc(appDir, projectName, "logic-rpc", true)

	// Create the logic-http service
	fmt.Println("Creating the gateway-http service.（gateway service）")
	_ = buildHttp(appDir, projectName, "gateway-http", true)
	// Pull the dependency configuration
	fmt.Println("Pulling the dependency configuration")
	cmd3 := exec.Command("go", "mod", "tidy")
	cmd3.Dir = appDir
	_ = cmd3.Run()
	fmt.Println(`Project creation completed!`)
}

func Rpc() error {
	pwd, err := file.GetPwd()
	if err != nil {
		return err
	}
	// Get project name
	projectName, err := GetProjectName()
	if err != nil {
		return err
	}
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

	return buildRpc(pwd, projectName, serverName, defaultConfig)
}

func buildRpc(pwd, projectName, serverName string, defaultConfig bool) error {
	var rpcBuilder RpcBuilder
	rpcBuilder.Tools.RpcServer = true
	rpcBuilder.PackageName = fmt.Sprintf("%s/%s", projectName, serverName)
	rpcBuilder.App = projectName
	rpcBuilder.ServerName = serverName
	rpcBuilder.Pwd = pwd

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
	pwd, err := file.GetPwd()
	if err != nil {
		return err
	}

	projectName, err := GetProjectName()
	if err != nil {
		return err
	}

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

	return buildHttp(pwd, projectName, serverName, defaultConfig)
}

func buildHttp(pwd, projectName, serverName string, defaultConfig bool) error {
	var httpBuilder HttpBuilder
	httpBuilder.Tools.Http = true
	httpBuilder.Tools.RpcClient = true
	httpBuilder.PackageName = fmt.Sprintf("%s/%s", projectName, serverName)
	httpBuilder.App = projectName
	httpBuilder.ServerName = serverName
	httpBuilder.Pwd = pwd
	var err error
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

func GetProjectName() (string, error) {
	pwd, err := file.GetPwd()
	if err != nil {
		return "", err
	}
	fmt.Println("Current road strength:", pwd)
	modFile := path.Join(pwd, "go.mod")
	if isExists, _ := file.CheckFileIsExist(modFile); !isExists {
		fmt.Println("no valid mod file found", isExists)
		return "", errors.New("no valid mod file found")
	}

	var projectName string
	var file *os.File
	file, err = os.Open(modFile)
	if err != nil {
		fmt.Printf("Cannot open text file: %s, err: [%v]", modFile, err)
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // or
		text := strings.TrimLeft(line, " ")
		if strings.HasPrefix(text, "module") {
			projectName = strings.TrimLeft(strings.TrimPrefix(text, "module"), " ")
			break
		}
	}

	if err = scanner.Err(); err != nil {
		log.Printf("Cannot scanner text file: %s, err: [%v]", modFile, err)
		return "", err
	}
	return projectName, nil
}
