package sshkeys

import (
	"HorizonX/common/xerr"
	"HorizonX/model"
	"context"
	"github.com/pkg/errors"

	"HorizonX/api/internal/svc"
	"HorizonX/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserSSHKeysLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserSSHKeysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSSHKeysLogic {
	return &GetUserSSHKeysLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserSSHKeysLogic) GetUserSSHKeys(req *types.GetUserSSHKeysReq) (resp *types.GetUserSSHKeysResp, err error) {
	uid := l.ctx.Value("uid").(string)
	whereBuilder := l.svcCtx.SshKeysModel.SelectBuilder().Where("user_id = ?", uid)
	findKeys, err := l.svcCtx.SshKeysModel.FindAll(l.ctx, whereBuilder, "id ASC")
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find ssh keys error: %s", err.Error())
	}
	if findKeys == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.GENERAL_NOT_FOUND_ERROR), "ssh keys not found [Uid:%s]", uid)
	}

	sshKeys := make([]types.SSHKey, 0)
	for _, key := range findKeys {
		sshKeys = append(sshKeys, types.SSHKey{
			Id:        key.Id,
			Name:      key.Name,
			PublicKey: key.Content,
			CreateAt:  key.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	resp = &types.GetUserSSHKeysResp{
		SSHKeys: sshKeys,
	}

	return resp, nil
}
