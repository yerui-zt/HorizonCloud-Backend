package order

import (
	"HorizonX/rpc/order/order"
	"context"
	"github.com/jinzhu/copier"

	"HorizonX/api/internal/svc"
	"HorizonX/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDetailLogic {
	return &GetOrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderDetailLogic) GetOrderDetail(req *types.GetOrderDetailReq) (resp *types.GetOrderDetailResp, err error) {
	rpcResp, err := l.svcCtx.OrderRPC.GetOrderDetailItem(l.ctx, &order.GetOrderDetailItemReq{
		OrderNo: req.OrderNo,
	})
	if err != nil {
		return nil, err
	}
	items := make([]types.GetOrderDetailItem, 0)
	copier.Copy(&items, &rpcResp.Items)
	resp = &types.GetOrderDetailResp{
		Id:          rpcResp.Id,
		CreateTime:  rpcResp.CreateTime,
		UpdateTime:  rpcResp.UpdateTime,
		DueDate:     rpcResp.DueDate,
		OrderNo:     rpcResp.OrderNo,
		UserId:      rpcResp.UserId,
		TotalAmount: rpcResp.TotalAmount,
		Status:      rpcResp.Status,
		Items:       items,
	}

	return resp, nil

}
