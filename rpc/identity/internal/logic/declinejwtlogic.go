package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cast"

	"HorizonX/rpc/identity/identity"
	"HorizonX/rpc/identity/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeclineJWTLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeclineJWTLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeclineJWTLogic {
	return &DeclineJWTLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeclineJWTLogic) DeclineJWT(in *identity.DeclineJWTReq) (*identity.DeclineJWTResp, error) {
	// 将token加入黑名单
	// key: jwt-blacklist:token value: uid
	key := fmt.Sprintf("jwt-blacklist:%v", in.Token)
	_, err := l.svcCtx.Redis.SetnxExCtx(
		l.ctx,
		key,
		cast.ToString(in.Uid),
		cast.ToInt(l.svcCtx.Config.Jwt.AccessExpire),
	)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("redis setnxex failed [key: %s]", key)
		return nil, errors.Wrapf(err, "redis setnxex failed [key: %s]", key)
	}

	return &identity.DeclineJWTResp{
		Success: true,
	}, nil
}
