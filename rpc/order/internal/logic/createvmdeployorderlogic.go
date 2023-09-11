package logic

import (
	"HorizonX/common/tools"
	"HorizonX/common/xerr"
	"HorizonX/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"

	"HorizonX/rpc/order/internal/svc"
	"HorizonX/rpc/order/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVMDeployOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateVMDeployOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVMDeployOrderLogic {
	return &CreateVMDeployOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateVMDeployOrder 创建虚拟机部署订单
func (l *CreateVMDeployOrderLogic) CreateVMDeployOrder(in *order.CreateVMDeployOrderReq) (*order.CreateVMDeployOrderResp, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, nil, in.Uid)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_FOUND_ERROR), "find user by id error [id: %d]", in.Uid)
	}

	plan, err := l.svcCtx.VmPlanModel.FindOne(l.ctx, nil, in.PlanId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_PLAN_NOT_FOUND), "find plan by id error [id: %d]", in.PlanId)
	}

	osImage, err := l.svcCtx.VmTemplateModel.FindOneByName(l.ctx, in.Image)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_VM_IMAGE_NOT_FOUND), "find vm template by name error [name: %s]", in.Image)
	}

	vmGroup, err := l.svcCtx.HypervisorGroupModel.FindOne(l.ctx, nil, in.VmGroupId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_VM_GROUP_NOT_FOUND), "find vm group by id error [id: %d]", in.VmGroupId)
	}

	var newOrder *model.Order
	// 开启事务 - 创建订单
	if err := l.svcCtx.OrderModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		// 1. 创建主订单
		now := time.Now()
		insertOrder := model.Order{
			DueDate: now.AddDate(0, 0, 1),
			OrderNo: tools.GenerateOrderNo(now),
			UserId:  user.Id,
			Type:    "general",
			Status:  "unpaid",
		}
		res, err := l.svcCtx.OrderModel.Insert(l.ctx, session, &insertOrder)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "create order failed [Insert failed] [err: %v]", err)
		}
		oid, err := res.LastInsertId()
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "create order failed [LastInsertId failed] [err: %v]", err)
		}
		newOrder, err = l.svcCtx.OrderModel.FindOne(l.ctx, session, oid)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "create order failed [FindOne failed] [err: %v]", err)
		}

		// 2. 创建订单Item
		amount := calculatePrice(plan, in.BillingCycle)
		if amount == 0 {
			return errors.Wrapf(xerr.NewErrCodeMsg(400, "invalid billing cycle"), "invalid billing cycle [billing cycle: %s]", in.BillingCycle)
		}
		// 创建content
		content, err := json.Marshal(&order.OrderItemContentVMCreateContent{
			Plan:    plan.Name,
			VMGroup: vmGroup.Name,
			OSImage: osImage.Name,
			ServicePeriod: fmt.Sprintf("%s ~ %s",
				now.Format("2006-01-02"),
				calculateDueDate(now, in.BillingCycle).Format("2006-01-02")),
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "marshal order item content failed [err: %v]", err)
		}
		// 创建action
		action, err := json.Marshal(&order.OrderItemActionVmInstanceCreateAction{
			HypervisorGroupId: plan.HypervisorGroupId,
			PlanID:            plan.Id,
			OSImageID:         osImage.Id,
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "marshal order item action failed [err: %v]", err)
		}

		insertOrderItem := model.OrderItem{
			OrderId:    oid,
			Name:       fmt.Sprintf("Cloud Instance  %s-%s", vmGroup.Name, plan.Name),
			ActionType: "vm_instance_create",
			Action:     string(action),
			Content:    string(content),
			Quantity:   1,
			Amount:     amount,
		}
		res, err = l.svcCtx.OrderItemModel.Insert(l.ctx, session, &insertOrderItem)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "create order item failed [Insert failed] [err: %v]", err)
		}

		// 3. 更新订单总价
		newOrder.TotalAmount += amount
		_, err = l.svcCtx.OrderModel.Update(l.ctx, session, newOrder)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "update order failed [Update failed] [err: %v]", err)
		}

		return nil

	}); err != nil {
		return nil, err
	}

	return &order.CreateVMDeployOrderResp{
		OrderNo: newOrder.OrderNo,
	}, nil

}

// calculateDueDate 计算订单的到期时间
func calculateDueDate(createTime time.Time, billingCycle string) time.Time {
	dueDate := createTime
	switch billingCycle {
	case "monthly":
		return dueDate.AddDate(0, 1, 0)
	case "quarterly":
		return dueDate.AddDate(0, 3, 0)
	case "semiAnnually":
		return dueDate.AddDate(0, 6, 0)
	case "annually":
		return dueDate.AddDate(1, 0, 0)
	default:
		return dueDate
	}
}

// calculatePrice 计算订单的价格
func calculatePrice(plan *model.VmPlan, cycle string) int64 {
	switch cycle {
	case "monthly":
		return plan.MonthlyPrice
	case "quarterly":
		return plan.QuarterlyPrice
	case "semiAnnually":
		return plan.SemiAnnuallyPrice
	case "annually":
		return plan.AnnuallyPrice
	default:
		return 0
	}
}
