package user

import (
	"HorizonX/rpc/user/user"
	"context"
	"github.com/pkg/errors"

	"HorizonX/api/internal/svc"
	"HorizonX/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	rpcResp, err := l.svcCtx.UserRPC.Register(l.ctx, &user.RegisterReq{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Country:   req.Country,
		Address:   req.Address,
		AffBy:     req.AffBy,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "register error [email: %s]", req.Email)
	}

	return &types.RegisterResp{
		Token: rpcResp.AccessToken,
	}, nil

}
