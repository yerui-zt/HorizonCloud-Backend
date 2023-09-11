package logic

import (
	"context"

	"HorizonX/rpc/order/internal/svc"
	"HorizonX/rpc/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVMDeployOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateVMDeployOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVMDeployOrderLogic {
	return &CreateVMDeployOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建虚拟机部署订单
func (l *CreateVMDeployOrderLogic) CreateVMDeployOrder(in *order.CreateVMDeployOrderReq) (*order.CreateVMDeployOrderResp, error) {
	// todo: add your logic here and delete this line

	return &order.CreateVMDeployOrderResp{}, nil
}
