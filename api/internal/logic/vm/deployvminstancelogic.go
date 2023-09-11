package vm

import (
	"context"

	"HorizonX/api/internal/svc"
	"HorizonX/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeployVMInstanceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeployVMInstanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeployVMInstanceLogic {
	return &DeployVMInstanceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeployVMInstanceLogic) DeployVMInstance(req *types.DeployVMInstanceReq) (resp *types.DeployVMInstanceResp, err error) {
	// todo: add your logic here and delete this line

	return
}
