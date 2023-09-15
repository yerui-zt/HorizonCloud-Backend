package logic

import (
	"HorizonX/common/cryptx"
	"HorizonX/common/xerr"
	"HorizonX/rpc/identity/identity"
	"HorizonX/rpc/user/internal/svc"
	"HorizonX/rpc/user/user"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// 1.数据库查找用户
	u, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find user by email error [email: %s]", in.Email)
	}

	// 2. bcrpt比对密码
	if cryptx.BcyptCheck(in.Password, u.Password) == false {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_PASSWORD_ERROR), "password error [email: %s]", in.Email)
	}

	// 3. 签发jwt
	rpcResp, err := l.svcCtx.Identity.IssueJWT(l.ctx, &identity.IssueJWTReq{
		Uid: u.Id,
		//Expire: timestamppb.Now(), // 此项目前无效
	})
	if err != nil {
		return nil, err
	}

	return &user.LoginResp{
		AccessToken: rpcResp.Token,
	}, nil
}
