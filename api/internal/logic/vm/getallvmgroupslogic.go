package vm

import (
	"HorizonX/common/xerr"
	"context"
	"github.com/pkg/errors"

	"HorizonX/api/internal/svc"
	"HorizonX/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllVMGroupsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllVMGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllVMGroupsLogic {
	return &GetAllVMGroupsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllVMGroupsLogic) GetAllVMGroups() (resp *types.GetAllVMGroupsResp, err error) {
	whereBuilder := l.svcCtx.HypervisorGroupModel.SelectBuilder().Where(
		"enable = ? and visible = ?", 1, 1)
	findGroups, err := l.svcCtx.HypervisorGroupModel.FindAll(l.ctx, whereBuilder, "id ASC")
	if err != nil {
		logx.WithContext(l.ctx).Errorf("find all vm groups failed: %v", err)
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(400, "no available vm group"), "get all vm groups failed: %v", err)
	}

	resp = &types.GetAllVMGroupsResp{}
	for _, group := range findGroups {
		resp.Groups = append(resp.Groups, types.VMGroup{
			Id:     group.Id,
			Name:   group.Name,
			Region: group.Region,
		})
	}
	return resp, nil
}
