package logic

import (
	"HorizonX/common/cryptx"
	"HorizonX/common/xerr"
	"HorizonX/model"
	"HorizonX/rpc/identity/identity"
	"HorizonX/rpc/user/internal/svc"
	"HorizonX/rpc/user/user"
	"context"
	"github.com/pkg/errors"
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
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "find user by email error [email: %s]", in.Email)
	}
	if find != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_EXIST_ERROR), "email already exists [email: %s]", in.Email)
	}

	// 2. 检查affBy是否真实
	affUser, err := l.svcCtx.UserModel.FindOne(l.ctx, in.AffBy)
	if err != nil && err != model.ErrNotFound {
		logx.WithContext(l.ctx).Errorf("find aff_by user by id error [id: %d]", in.AffBy)
	}
	if affUser == nil {
		in.AffBy = 0
	}

	// 3. 创建用户
	cPass, err := cryptx.BcryptHash(in.Password)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "create user [email: %s] failed [BcryptHash failed]", in.Email)
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
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "create user [email: %s] failed [Insert failed]", in.Email)
	}
	uid, err := result.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "create user [email: %s] failed [LastInsertId failed]", in.Email)
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
