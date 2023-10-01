package logic

import (
	"HorizonX/common/xerr"
	"HorizonX/model"
	"HorizonX/rpc/order/internal/svc"
	"HorizonX/rpc/order/order"
	"HorizonX/rpc/payment/payment"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayOrderLogic {
	return &PayOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PayOrder 支付订单
func (l *PayOrderLogic) PayOrder(in *order.PayOrderReq) (*order.PayOrderResp, error) {
	findOrder, err := l.svcCtx.OrderModel.FindOneByOrderNo(l.ctx, in.OrderNo)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_NOT_FOUND), "Order not found [%s] [%s]", in.OrderNo, err.Error())
	}

	if findOrder.Status != "unpaid" {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_HAS_PAID), "Order has been paid [%s]", in.OrderNo)
	}
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, nil, findOrder.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_FOUND_ERROR), "unknown user [%s]", findOrder.UserId)
	}

	//判断余额是否足够，如果足够调用 payByBalance
	//如果不够，调用 支付接口获取支付url
	if user.Balance >= findOrder.TotalAmount && findOrder.Type != "addFunds" {
		err := l.payByBalance(user, findOrder)
		if err != nil {
			return nil, err
		}
		return &order.PayOrderResp{
			Url: "nil",
		}, nil
	}

	// 调用支付接口，创建付款url
	rpcResp, err := l.svcCtx.PaymentRPC.CreatePayment(l.ctx, &payment.CreatePaymentReq{
		OrderNo: findOrder.OrderNo,
		UserId:  user.Id,
		Method:  in.Method,
	})
	if err != nil {
		return nil, err
	}

	resp := &order.PayOrderResp{
		Url: rpcResp.Url,
	}
	return resp, nil

}

func (l *PayOrderLogic) payByBalance(user *model.User, o *model.Order) error {

	err := l.svcCtx.OrderModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		user.Balance -= o.TotalAmount
		_, err := l.svcCtx.UserModel.Update(l.ctx, session, user)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "update user balance failed [%s]", err.Error())
		}

		fullFillOrderLogic := NewFullFillOrderLogic(l.ctx, l.svcCtx)
		_, err = fullFillOrderLogic.FullFillOrder(&order.FullFillOrderReq{
			OrderNo:    o.OrderNo,
			CallbackNo: "",
			Method:     "balance",
		})
		if err != nil {
			return err
		}
		return nil
	})

	return err
}
