package logic

import (
	"HorizonX/common/aqueue/jobtype"
	"HorizonX/common/xerr"
	"HorizonX/model"
	"HorizonX/rpc/mqueue/worker/internal/svc"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
)

type OrderExecActionMqHandler struct {
	svcCtx *svc.ServiceContext
}

func NewOrderExecActionMqHandler(svcCtx *svc.ServiceContext) *OrderExecActionMqHandler {
	return &OrderExecActionMqHandler{
		svcCtx: svcCtx,
	}
}

func (l *OrderExecActionMqHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p jobtype.OrderExecActionMqPayload
	err := json.Unmarshal(t.Payload(), &p)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.JSON_UNMARSHAL_ERROR), "unmarshal payload error [payload: %s]", string(t.Payload()))
	}

	findOrder, err := l.svcCtx.OrderModel.FindOne(ctx, nil, p.OrderId)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.ORDER_NOT_FOUND), "find order error [orderId: %d]", p.OrderId)
	}

	whereBuilder := l.svcCtx.OrderItemModel.SelectBuilder().Where("order_id = ?", findOrder.Id)
	items, err := l.svcCtx.OrderItemModel.FindAll(ctx, whereBuilder, "id ASC")
	if err != nil && err != model.ErrNotFound {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "get order item failed: %v", err)
	}

	if items == nil {
		return nil
	}

	for _, item := range items {
		switch item.ActionType {
		case "vm_instance_create":
			fmt.Println("vm_instance_create", p.OrderId)

		case "addFunds":
		default:
			return nil
		}
	}

	return nil
}
