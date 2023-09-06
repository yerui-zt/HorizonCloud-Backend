package logic

import (
	"HorizonX/common/cryptx"
	"HorizonX/model"
	"HorizonX/rpc/identity/identity"
	"HorizonX/rpc/user/internal/svc"
	"HorizonX/rpc/user/user"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// 1. 检查邮箱是否存在
	find, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if find != nil {
		return nil, err
	}

	// 2. 检查affBy是否真实
	find, _ = l.svcCtx.UserModel.FindOne(l.ctx, in.AffBy)
	if find == nil {
		in.AffBy = 0
	}

	// 3. 创建用户
	cPass, err := cryptx.BcryptHash(in.Password)
	if err != nil {
		return nil, err
	}
	insertUser := &model.User{
		Email:     in.Email,
		Password:  cPass,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Country:   in.Country,
		Address:   in.Address,
		AffBy:     in.AffBy,
	}
	result, err := l.svcCtx.UserModel.Insert(l.ctx, insertUser)
	if err != nil {
		return nil, err
	}
	uid, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// 4. 签发jwt
	rpcResp, err := l.svcCtx.Identity.IssueJWT(l.ctx, &identity.IssueJWTReq{
		Uid:    uid,
		Expire: timestamppb.Now(), // 此项目前无效
	})
	if err != nil {
		return nil, err
	}

	//todo: 发送验证码

	return &user.RegisterResp{
		AccessToken: rpcResp.Token,
	}, nil

}
