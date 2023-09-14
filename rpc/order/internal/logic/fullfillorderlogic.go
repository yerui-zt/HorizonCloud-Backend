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

type FullFillOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFullFillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FullFillOrderLogic {
	return &FullFillOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FullFillOrder 订单支付成功后处理回调
func (l *FullFillOrderLogic) FullFillOrder(in *order.FullFillOrderReq) (*order.FullFillOrderResp, error) {
	findOrder, err := l.svcCtx.OrderModel.FindOneByOrderNo(l.ctx, in.OrderNo)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_NOT_FOUND), "order not found %s", in.OrderNo)
	}
	if findOrder.Status != "unpaid" {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_STATUS_ERROR), "order status error %s [%s]", in.OrderNo, err.Error())
	}

	// 1. 记录支付信息
	// todo: 拆分到 paymentRPC
	findOrder.CallbackNo.String = in.CallbackNo
	findOrder.CallbackNo.Valid = true
	findOrder.Status = "paid"
	findOrder.PaymentMethod = in.Method
	_, err = l.svcCtx.OrderModel.Update(l.ctx, nil, findOrder)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "update order failed: %v", err)
	}

	return &order.FullFillOrderResp{}, nil

	// 发布订单支付成功事件

}

func (l *FullFillOrderLogic) handleAction(o *model.Order) error {
	whereBuilder := l.svcCtx.OrderItemModel.SelectBuilder().Where("order_id = ?", o.Id)
	findItems, err := l.svcCtx.OrderItemModel.FindAll(l.ctx, whereBuilder, "id ASC")
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.ORDER_ITEM_PARSE_ERROR), "get order item [order %d] failed: %v ", o.Id, err)
	}
	if findItems == nil {
		return nil
	}

	for _, item := range findItems {
		switch item.ActionType {
		case "vm_instance_create":

		case "addFunds":
		default:
			return nil
		}
	}

	return nil
}
