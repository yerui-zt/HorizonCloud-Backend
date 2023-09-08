package vm

import (
	"HorizonX/common/xerr"
	"context"
	"github.com/pkg/errors"

	"HorizonX/api/internal/svc"
	"HorizonX/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetImageLogic {
	return &GetImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetImageLogic) GetImage() (resp *types.GetImageResp, err error) {
	whereBuild := l.svcCtx.VmTemplateModel.SelectBuilder().Where(
		"enable = 1 and visible = 1")
	findImages, err := l.svcCtx.VmTemplateModel.FindAll(l.ctx, whereBuild, "id ASC")
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(400, "no available image"), "get all images failed: %v", err)
	}

	resp = &types.GetImageResp{}
	for _, image := range findImages {
		resp.Images = append(resp.Images, types.VMImage{
			Id:      image.Id,
			Name:    image.Name,
			Release: image.Release,
		})
	}

	return resp, nil
}
