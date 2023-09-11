package order

import (
	"HorizonX/api/internal/svc"
	"HorizonX/api/internal/types"
	"HorizonX/rpc/order/order"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderPaymentMethodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderPaymentMethodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderPaymentMethodLogic {
	return &GetOrderPaymentMethodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderPaymentMethodLogic) GetOrderPaymentMethod(req *types.GetOrderPaymentMethodReq) (resp *types.GetOrderPaymentMethodResp, err error) {
	rpcResp, err := l.svcCtx.OrderRPC.GetOrderPaymentMethod(l.ctx, &order.GetOrderPaymentMethodReq{
		OrderNo: req.OrderNo,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.GetOrderPaymentMethodResp{}
	for _, method := range rpcResp.Methods {
		resp.Methods = append(resp.Methods, types.OrderPaymentMethod{
			Name: method.Name,
			Type: method.Type,
		})
	}

	return resp, nil
}
