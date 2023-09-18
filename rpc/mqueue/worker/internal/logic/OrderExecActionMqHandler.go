package logic

import (
	"HorizonX/common/xerr"
	"HorizonX/rpc/mqueue/worker/internal/svc"
	"HorizonX/rpc/mqueue/worker/jobtype"
	"HorizonX/rpc/order/order"
	"HorizonX/rpc/vm/vmservice"
	"context"
	"encoding/json"
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

	item := p.Item
	findOrder, err := l.svcCtx.OrderModel.FindOne(ctx, nil, item.OrderId)
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find order error [order_id: %d]", item.OrderId)
	}
	// 开始执行
	switch item.ActionType {
	case "vm_instance_create":
		actionStr := item.Action
		var action order.OrderItemActionVmInstanceCreateAction
		err := json.Unmarshal([]byte(actionStr), &action)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.JSON_UNMARSHAL_ERROR), "unmarshal action error [action: %s]", actionStr)
		}
		_, err = l.svcCtx.VmRPC.DeployVMInstance(ctx, &vmservice.DeployVMInstanceReq{
			Hostname:     action.HostName,
			BillingCycle: action.BillingCycle,
			ImageId:      action.OSImageID,
			GroupId:      action.HypervisorGroupId,
			PlanId:       action.PlanID,
			UserId:       findOrder.UserId,
		})
		if err != nil {
			return err
		}
		return nil
	case "addFunds":
	default:
		return nil
	}

	return nil
}
