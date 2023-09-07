package user

import (
	"HorizonX/api/internal/svc"
	"HorizonX/api/internal/types"
	"HorizonX/rpc/user/user"
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout() (resp *types.LoginOutResp, err error) {
	uid := l.ctx.Value("uid").(string)
	jwtToken := l.ctx.Value("jwtToken").(string)

	rpcResp, err := l.svcCtx.UserRPC.Logout(l.ctx, &user.LogoutReq{
		Uid:         cast.ToInt64(uid),
		AccessToken: jwtToken,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "rpc user logout error")
	}

	return &types.LoginOutResp{
		Success: rpcResp.Success,
	}, nil
}
