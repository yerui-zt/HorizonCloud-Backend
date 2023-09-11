package vm

import (
	"HorizonX/rpc/order/orderservice"
	"context"
	"github.com/spf13/cast"

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
	uid := l.ctx.Value("uid").(string)
	rpcResp, err := l.svcCtx.OrderRPC.CreateVMDeployOrder(l.ctx, &orderservice.CreateVMDeployOrderReq{
		Uid:          cast.ToInt64(uid),
		VmGroupId:    req.GroupId,
		PlanId:       req.PlanID,
		Image:        req.Image,
		BillingCycle: req.BillingCycle,
	})
	if err != nil {
		return nil, err
	}

	return &types.DeployVMInstanceResp{
		OrderNo: rpcResp.OrderNo,
	}, nil
}
