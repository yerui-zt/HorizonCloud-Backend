package user

import (
	"HorizonX/rpc/user/user"
	"context"

	"HorizonX/api/internal/svc"
	"HorizonX/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	rpcResp, err := l.svcCtx.UserRPC.Login(l.ctx, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		// todo: 状态码处理
		return nil, err
	}

	return &types.LoginResp{
		Token: rpcResp.AccessToken,
	}, nil
}
