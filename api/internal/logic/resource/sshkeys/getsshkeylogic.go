package sshkeys

import (
	"HorizonX/common/xerr"
	"HorizonX/model"
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/cast"

	"HorizonX/api/internal/svc"
	"HorizonX/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSSHKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSSHKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSSHKeyLogic {
	return &GetSSHKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSSHKeyLogic) GetSSHKey(req *types.GetSSHKeyReq) (resp *types.GetSSHKeyResp, err error) {
	// 判断用户是否有权限获取该 SSH Key
	uid := l.ctx.Value("uid").(string)
	sshKey, err := l.svcCtx.SshKeysModel.FindOne(l.ctx, nil, cast.ToInt64(req.KeyId))
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find ssh key error: %s", err.Error())
	}
	if sshKey == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.GENERAL_NOT_FOUND_ERROR), "ssh key not found [Uid:%s] [keyId:%s]", uid, req.KeyId)
	}

	if sshKey.UserId != cast.ToInt64(uid) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.UNAUTHORIZED), "unauthorized to get this ssh key [Uid:%s] [keyId:%s]", uid, req.KeyId)
	}

	resp = &types.GetSSHKeyResp{
		SSHKey: types.SSHKey{
			Id:        sshKey.Id,
			Name:      sshKey.Name,
			PublicKey: sshKey.Content,
			CreateAt:  sshKey.CreateTime.Format("2006-01-02 15:04:05"),
		},
	}

	return resp, nil
}
