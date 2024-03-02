package service

import (
	"context"
	"example-project/common"
)

type RpcApi struct {
}

func (ra *RpcApi) Test(ctx context.Context, req common.TestParams, resp *common.TestResp) error {
	resp.Data = "Hello, Go-dandelion"
	return nil
}
