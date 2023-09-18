package sshkeys

import (
	"HorizonX/common/xerr"
	"HorizonX/model"
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"golang.org/x/crypto/ssh"

	"HorizonX/api/internal/svc"
	"HorizonX/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSSHKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSSHKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSSHKeyLogic {
	return &AddSSHKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSSHKeyLogic) AddSSHKey(req *types.AddSSHKeyReq) (resp *types.AddSSHKeyResp, err error) {
	// todo 限制用户最多添加 10 个 SSH Key

	// 1. 验证 SSH Key 合法性
	_, _, _, _, err = ssh.ParseAuthorizedKey([]byte(req.PublicKey))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.INVALID_SSH_KEY), "invalid ssh key: %s", err.Error())
	}

	// 2. 存入数据库
	uid := l.ctx.Value("uid").(string)
	insertKey := &model.SshKeys{
		UserId:  cast.ToInt64(uid),
		Name:    req.Name,
		Content: req.PublicKey,
	}
	res, err := l.svcCtx.SshKeysModel.Insert(l.ctx, nil, insertKey)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "insert ssh key error: %s", err.Error())
	}
	keyId, err := res.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find ssh key id error: %s", err.Error())
	}
	key, err := l.svcCtx.SshKeysModel.FindOne(l.ctx, nil, keyId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find ssh key error: %s", err.Error())
	}

	resp = &types.AddSSHKeyResp{
		SSHKey: types.SSHKey{
			Id:        key.Id,
			Name:      key.Name,
			PublicKey: key.Content,
			CreateAt:  key.CreateTime.Format("2006-01-02 15:04:05"),
		},
	}

	return resp, nil
}
