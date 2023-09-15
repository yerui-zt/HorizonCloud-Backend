package logic

import (
	"HorizonX/common/aqueue/jobtype"
	"HorizonX/common/xerr"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

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

	err = l.svcCtx.OrderModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		// 1. 记录支付信息
		findOrder.CallbackNo.String = in.CallbackNo
		findOrder.CallbackNo.Valid = true
		findOrder.Status = "paid"
		findOrder.PaymentMethod = in.Method
		_, err = l.svcCtx.OrderModel.Update(l.ctx, session, findOrder)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "update order failed: %v", err)
		}

		// 2. 发布订单支付成功事件
		// 使用消息队列，异步处理订单
		task, err := jobtype.NewOrderExecActionMq(findOrder.Id)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.MQ_PUBLISH_ERROR), "publish order exec action mq failed: %v", err)
		}
		_, err = l.svcCtx.AsynqClient.Enqueue(task)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.MQ_PUBLISH_ERROR), "publish order exec action mq failed: %v", err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &order.FullFillOrderResp{
		Success: true,
	}, nil
}
