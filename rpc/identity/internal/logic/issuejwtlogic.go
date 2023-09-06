package logic

import (
	"HorizonX/common/jwt"
	"HorizonX/rpc/identity/identity"
	"HorizonX/rpc/identity/internal/svc"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type IssueJWTLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIssueJWTLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IssueJWTLogic {
	return &IssueJWTLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IssueJWTLogic) IssueJWT(in *identity.IssueJWTReq) (*identity.IssueJWTResp, error) {
	token, err := jwt.NewJWT(l.svcCtx.Config.Jwt.AccessSecret).
		IssueGeneralUserToken(
			in.Uid,
			l.svcCtx.Config.Jwt.AccessExpire,
			l.svcCtx.Config.Jwt.Issuer,
		)
	if err != nil {
		logx.Errorf("IssueJWTLogic.IssueJWT err: %v", err)
	}

	return &identity.IssueJWTResp{
		Token: token,
	}, nil
}
