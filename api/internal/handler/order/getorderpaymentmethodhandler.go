package order

import (
	"net/http"

	"HorizonX/common/result"
	"HorizonX/common/validator"

	"HorizonX/api/internal/logic/order"
	"HorizonX/api/internal/svc"
	"HorizonX/api/internal/types"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetOrderPaymentMethodHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetOrderPaymentMethodReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		if errMsg, errCode := validator.Validate(req); errCode != 0 {
			result.ParamErrorResult(r, w, errors.New(errMsg))
			return
		}

		l := order.NewGetOrderPaymentMethodLogic(r.Context(), svcCtx)
		resp, err := l.GetOrderPaymentMethod(&req)
		result.HttpResult(r, w, resp, err)
	}
}
