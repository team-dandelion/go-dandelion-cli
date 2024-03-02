package service

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
