package order

import (
	"HorizonX/rpc/order/order"
	"context"

	"HorizonX/api/internal/svc"
	"HorizonX/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayOrderLogic {
	return &PayOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayOrderLogic) PayOrder(req *types.PayOrderReq) (resp *types.PayOrderResp, err error) {
	rpcResp, err := l.svcCtx.OrderRPC.PayOrder(l.ctx, &order.PayOrderReq{
		OrderNo: req.OrderNo,
		Method:  req.Method,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.PayOrderResp{
		Url: rpcResp.Url,
	}
	return resp, nil
}
