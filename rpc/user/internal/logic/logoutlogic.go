package logic

import (
	"HorizonX/common/xerr"
	"HorizonX/rpc/identity/identity"
	"context"
	"fmt"
	"github.com/pkg/errors"

	"HorizonX/rpc/user/internal/svc"
	"HorizonX/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *user.LogoutReq) (*user.LogoutResp, error) {
	fmt.Printf("%v\n", in)
	success, err := l.svcCtx.Identity.DeclineJWT(
		l.ctx,
		&identity.DeclineJWTReq{
			Token: in.AccessToken,
			Uid:   in.Uid,
		},
	)
	if err != nil || success.Success == false {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "decline jwt error [token: %s]", in.AccessToken)
	}

	return &user.LogoutResp{
		Success: true,
	}, nil
}
