## 🖥Using the go-dandelion-cli

### 1.Install
```
go get github.com/team-dandelion/go-dandelion-cli@latest
go install github.com/team-dandelion/go-dandelion-cli@latest
```

### 2.Create project
In the local directory, create a project as prompted.
```shell
# Create project
go-dandelion-cli project -n example-project
```
+ -n: project name

Project directory structure：
```shell
│
├── common  // Used to store public structures
│
│
├── logic-rpc
│   ├── boot
│   │   └── boot.go              // Register initialization methods here
│   ├── cmd
│   │   ├── api
│   │   │   └── server.go        // Service startup entry point
│   │   └── cobra.go             // Cobra command registration
│   ├── config                    // Service configuration folder
│   │   └── configs_local.yaml    // Local configuration file
│   ├── global
│   │   └── global.go            // Global variables
│   ├── internal
│   │   ├── dao                  // Database operations
│   │   ├── enum                 // Enums and constants
│   │   ├── logic                // Business logic
│   │   ├── model                // Data models
│   │   └── service              // Services
│   │       └── api.go           // Service interface
│   ├── static
│   │   └── rpc-server.txt       // Service name
│   ├── tools                    // Utility classes
│   │
│   └── main.go                  // Entry file
│
│
├── gateway-http
│   ├── cmd
│   │   ├── api
│   │   │   └── server.go        // Service startup entry point
│   │   └── cobra.go             // Cobra command registration
│   ├── config                    // Service configuration folder
│   │   └── configs_local.yaml    // Local configuration file
│   ├── internal
│   │   ├── middleware           // Custom middleware
│   │   ├── route                // Route management
│   │   │   └── route.go         // Provides basic routes
│   │   └── service              // Services
│   ├── static
│   │   └── http-server.txt       // Service name
│   │ 
│   └── main.go                  // Entry file
```

### 3.Run service
#### 3.1 Run rpc service
**Enter the service directory**

```go
cd example-project/logic-rpc
```

**Remove redundant configuration**

In `config/configs_local.yaml`, remove the mysql and redis configurations

**Added test rpc methods**
Add a Test Model to `common/test_model.go`
```go
type (
	TestParams struct{}
	TestResp   struct {
		Data string
	}
)

```
Add a Test rpc method to `service/api.go`
```go
import (
	"context"
	"example-project/common"
)

type RpcApi struct {
}

func (ra *RpcApi) Test(ctx context.Context, req common.TestParams, resp *common.TestResp) (err error) {
	resp.Data = "Hello, go-dandelion"
	return nil
}
```

**Start service**
```shell
go build -o logic-rpc
#运行
./logic-rpc server
```

#### 3.2 Run http service
**Enter the service directory**
```shell
cd example-project/gateway-http
```
**Add api interface**

Add a TestFunc method to `service/test_controller.go`
```shell
import (
	"example-project/common"
	routing "github.com/gly-hub/fasthttp-routing"
	"github.com/team-dandelion/go-dandelion/application"
	"github.com/team-dandelion/go-dandelion/server/http"
)

type AuthController struct {
	http.HttpController
}

func (a *AuthController) TestFunc(c *routing.Context) error {
	return application.SRpcCall(c, "logic-rpc", "Test", new(common.TestParams), new(common.TestResp))
}

```
Registered routes in `route/route.go`
```shell
testController := new(service.AuthController)
baseRouter.Get("/test", testController.TestFunc)
```

**Start service**
```shell
go build -o gateway-http
#运行
./gateway-http server
```

**test**
```shell
curl --location 'http://172.16.49.201:8080/api/test' --data ''
```

### 4.Customize the creation service
#### 4.1 Go to the application directory
cd example-project
#### 4.2 Build service
```shell
go-dandelion-cli server
```
output:
```shell
Type of service you want to create, enter a number（1-rpc 2-http）:1
RPC SERVICE NAME: example-server
Whether to initialize mysql?（y/n）:y
Whether to initialize redis?（y/n）:y
Whether to initialize the logger?（y/n）:y
Whether to initialize the trace link?（y/n）:y
```
