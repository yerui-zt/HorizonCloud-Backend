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

type DeleteSSHKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSSHKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSSHKeyLogic {
	return &DeleteSSHKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSSHKeyLogic) DeleteSSHKey(req *types.DeleteSSHKeyReq) (resp *types.DeleteSSHKeyResp, err error) {
	uid := l.ctx.Value("uid").(string)
	sshKey, err := l.svcCtx.SshKeysModel.FindOne(l.ctx, nil, cast.ToInt64(req.KeyId))
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find ssh key error: %s", err.Error())
	}
	if sshKey == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.GENERAL_NOT_FOUND_ERROR), "ssh key not found [Uid:%s] [keyId:%s]", uid, req.KeyId)
	}

	// 验证用户是否有权限删除该 SSH Key
	if sshKey.UserId != cast.ToInt64(uid) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.UNAUTHORIZED), "unauthorized to delete this ssh key [Uid:%s] [keyId:%s]", uid, req.KeyId)
	}

	// 删除 SSH Key
	err = l.svcCtx.SshKeysModel.Delete(l.ctx, nil, cast.ToInt64(req.KeyId))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "delete ssh key error: %s", err.Error())
	}

	return &types.DeleteSSHKeyResp{
		Msg: "Success"}, nil

}
