package vm

import (
	"HorizonX/common/xerr"
	"context"
	"github.com/pkg/errors"

	"HorizonX/api/internal/svc"
	"HorizonX/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVMPlanByGroupIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVMPlanByGroupIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVMPlanByGroupIdLogic {
	return &GetVMPlanByGroupIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVMPlanByGroupIdLogic) GetVMPlanByGroupId(req *types.GetVMPlanByGroupIdReq) (resp *types.GetVMPlanByGroupIdResp, err error) {
	whereBuilder := l.svcCtx.VmPlanModel.SelectBuilder().Where(
		"hypervisor_group_id = ? and enable = 1 and visible = 1", req.GroupId)
	findPlan, err := l.svcCtx.VmPlanModel.FindAll(l.ctx, whereBuilder, "priority DESC")
	if err != nil {
		logx.WithContext(l.ctx).Errorf("find all vm plan by group id failed: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(400, "no available vm plan"), "get all vm plan by group id failed: %v", err)
	}

	resp = &types.GetVMPlanByGroupIdResp{}
	for _, plan := range findPlan {
		resp.Plans = append(resp.Plans, types.VMPlan{
			Id:                plan.Id,
			Name:              plan.Name,
			Stock:             plan.Stock,
			Vcpu:              plan.Vcpu,
			Memory:            plan.Memory,
			Disk:              plan.Disk,
			DataTransfer:      plan.DataTransfer,
			Bandwidth:         plan.Bandwidth,
			Ipv4Num:           plan.Ipv4Num,
			Ipv6Num:           plan.Ipv6Num,
			MonthlyPrice:      plan.MonthlyPrice,
			QuarterlyPrice:    plan.QuarterlyPrice,
			SemiAnnuallyPrice: plan.SemiAnnuallyPrice,
			AnnuallyPrice:     plan.AnnuallyPrice,
		})
	}

	return resp, nil

}
