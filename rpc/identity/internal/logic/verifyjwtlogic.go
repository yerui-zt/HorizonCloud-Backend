package logic

import (
	"HorizonX/rpc/identity/identity"
	"HorizonX/rpc/identity/internal/svc"
	"context"
	"fmt"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyJWTLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyJWTLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyJWTLogic {
	return &VerifyJWTLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyJWTLogic) VerifyJWT(in *identity.VerifyJWTReq) (*identity.VerifyJWTResp, error) {
	// 采用黑名单策略，如果redis中存在此token，则说明token已经失效

	// key: jwt-blacklist:token value: uid
	exist, err := l.svcCtx.Redis.ExistsCtx(l.ctx, fmt.Sprintf("jwt-blacklist:%s", in.Token))
	if err != nil {
		return nil, errors.Wrapf(err, "redis exists failed [key: %s]", fmt.Sprintf("jwt-blacklist:%s", in.Token))
	}

	return &identity.VerifyJWTResp{
		Valid: !exist,
	}, nil
}
