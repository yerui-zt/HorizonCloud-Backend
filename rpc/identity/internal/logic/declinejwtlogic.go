package logic

import (
	"context"
	"fmt"
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
	_, err := l.svcCtx.Redis.SetnxExCtx(
		l.ctx,
		fmt.Sprintf("jwt-blacklist:%s", in.Token),
		cast.ToString(in.Uid),
		cast.ToInt(l.svcCtx.Config.Jwt.AccessExpire),
	)
	if err != nil {
		return nil, err
	}

	return &identity.DeclineJWTResp{
		Success: true,
	}, nil
}
