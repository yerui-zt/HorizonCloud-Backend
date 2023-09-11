package logic

import (
	"HorizonX/common/xerr"
	"HorizonX/model"
	"context"
	"github.com/pkg/errors"

	"HorizonX/rpc/order/internal/svc"
	"HorizonX/rpc/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderDetailItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderDetailItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDetailItemLogic {
	return &GetOrderDetailItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetOrderDetailItem 获取订单详情
func (l *GetOrderDetailItemLogic) GetOrderDetailItem(in *order.GetOrderDetailItemReq) (*order.GetOrderDetailItemResp, error) {
	findOrder, err := l.svcCtx.OrderModel.FindOneByOrderNo(l.ctx, in.OrderNo)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_NOT_FOUND), "order not found %s", in.OrderNo)
	}

	resp := &order.GetOrderDetailItemResp{
		Id:          findOrder.Id,
		CreateTime:  findOrder.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:  findOrder.UpdateTime.Format("2006-01-02 15:04:05"),
		DueDate:     findOrder.DueDate.Format("2006-01-02 15:04:05"),
		OrderNo:     findOrder.OrderNo,
		UserId:      findOrder.UserId,
		TotalAmount: findOrder.TotalAmount,
		Status:      findOrder.Status,
		Items:       nil,
	}

	whereBuilder := l.svcCtx.OrderItemModel.SelectBuilder().Where("order_id = ?", findOrder.Id)
	items, err := l.svcCtx.OrderItemModel.FindAll(l.ctx, whereBuilder, "id ASC")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get order item failed: %v", err)
	}
	if items != nil {
		for _, item := range items {
			resp.Items = append(resp.Items, &order.OrderDetailItem{
				Id:       item.Id,
				OrderId:  item.OrderId,
				Name:     item.Name,
				Content:  item.Content,
				Quantity: item.Quantity,
				Amount:   item.Amount,
			})
		}
	}

	return resp, nil
}
