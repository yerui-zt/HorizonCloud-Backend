package result

import (
	"HorizonX/common/xerr"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
)

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		// 默认错误返回，如果错误类型不是 xerr.CodeError 时返回
		errcode := xerr.Code(xerr.SERVER_COMMON_ERROR)
		errmsg := xerr.SERVER_COMMON_ERROR

		// 如果错误类型是 xerr.CodeError 时返回
		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				errcode = grpcCode
				errmsg = gstatus.Message()
			}
		}

		logx.WithContext(r.Context()).Errorf("[API-ERR] : %+v ", err)

		httpx.WriteJson(w, int(errcode), Error(errcode, errmsg))
	}
}

// JWT鉴权失败返回
func AuthHttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		// 默认错误返回，如果错误类型不是 xerr.CodeError 时返回
		errcode := xerr.Code(xerr.UNAUTHORIZED)
		errmsg := xerr.UNAUTHORIZED

		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*xerr.CodeError); ok {
			//自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				errcode = grpcCode
				errmsg = gstatus.Message()
			}
		}

		logx.WithContext(r.Context()).Errorf("[GATEWAY-ERR] : %+v ", err)

		httpx.WriteJson(w, http.StatusUnauthorized, Error(errcode, errmsg))
	}
}

// ParamErrorResult http 参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s: %s", xerr.Code(xerr.REQUEST_PARAM_ERROR), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(xerr.Code(xerr.REQUEST_PARAM_ERROR), errMsg))
}
