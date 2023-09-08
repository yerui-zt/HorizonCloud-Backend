package vm

import (
	"HorizonX/common/xerr"
	"context"
	"github.com/pkg/errors"

	"HorizonX/api/internal/svc"
	"HorizonX/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVMGroupByRegionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVMGroupByRegionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVMGroupByRegionLogic {
	return &GetVMGroupByRegionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVMGroupByRegionLogic) GetVMGroupByRegion(req *types.GetVMGroupByRegionReq) (resp *types.GetVMGroupByRegionResp, err error) {
	whereBuilder := l.svcCtx.HypervisorGroupModel.SelectBuilder().Where(
		"region = ? and enable = 1 and visible = 1", req.Region)

	findGroups, err := l.svcCtx.HypervisorGroupModel.FindAll(l.ctx, whereBuilder, "id ASC")
	if err != nil {
		logx.WithContext(l.ctx).Errorf("find all vm groups by region failed: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(400, "no available vm group"), "get all vm groups by region failed: %v", err)
	}

	resp = &types.GetVMGroupByRegionResp{}
	for _, group := range findGroups {
		resp.Groups = append(resp.Groups, types.VMGroup{
			Id:     group.Id,
			Name:   group.Name,
			Region: group.Region,
		})
	}

	return resp, nil

}
