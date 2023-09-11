package logic

import (
	"HorizonX/common/xerr"
	"context"
	"github.com/pkg/errors"

	"HorizonX/rpc/order/internal/svc"
	"HorizonX/rpc/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderPaymentMethodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderPaymentMethodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderPaymentMethodLogic {
	return &GetOrderPaymentMethodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetOrderPaymentMethod 获取订单支付方式
func (l *GetOrderPaymentMethodLogic) GetOrderPaymentMethod(in *order.GetOrderPaymentMethodReq) (*order.GetOrderPaymentMethodResp, error) {
	_, err := l.svcCtx.OrderModel.FindOneByOrderNo(l.ctx, in.OrderNo)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_NOT_FOUND), "order not found %s [%v]", in.OrderNo, err)
	}

	whereBuilder := l.svcCtx.SystemPaymentMethodModel.SelectBuilder().Where("enable = 1")
	methods, err := l.svcCtx.SystemPaymentMethodModel.FindAll(l.ctx, whereBuilder, "id ASC")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(400, "No available payment method"), "get payment methods failed: %v", err)
	}

	var resp = &order.GetOrderPaymentMethodResp{}
	for _, method := range methods {
		resp.Methods = append(resp.Methods, &order.OrderPaymentMethod{
			Name: method.Name,
			Type: method.Type,
		})
	}

	return resp, nil
}
