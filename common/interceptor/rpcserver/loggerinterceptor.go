package rpcserver

import (
	"HorizonX/common/xerr"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	resp, err = handler(ctx, req)
	if err != nil {
		causeErr := errors.Cause(err) // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok {
			// 如果是 xerr 的自定义错误，则返回自定义的错误内容
			// 如果是 一般错误，则只记录错误日志
			// 一般的，推荐使用一般错误，不要使用自定义错误

			logx.WithContext(ctx).Errorf("[RPC-SRV-ERR] %+v", err)

			err = status.Error(codes.Code(e.GetErrCode()), e.GetErrMsg())
		} else {
			logx.WithContext(ctx).Errorf("[RPC-SRV-ERR] %+v", err)
		}

	}

	return resp, err
}
