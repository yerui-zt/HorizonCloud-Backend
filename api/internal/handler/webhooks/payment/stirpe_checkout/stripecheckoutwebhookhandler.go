package stirpe_checkout

import (
	"bytes"
	"io"
	"net/http"

	"HorizonX/common/result"
	"HorizonX/common/validator"

	"HorizonX/api/internal/logic/webhooks/payment/stirpe_checkout"
	"HorizonX/api/internal/svc"
	"HorizonX/api/internal/types"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func StripeCheckoutWebhookHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var bodyBytes []byte
		bodyBytes, _ = io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		bodyCpy := io.NopCloser(bytes.NewBuffer(bodyBytes))
		const MaxBodyBytes = int64(65536)
		bodyCpy = http.MaxBytesReader(w, bodyCpy, MaxBodyBytes)
		sessionObj, err := io.ReadAll(bodyCpy)
		if err != nil {
			stripeErr := errors.New("StripeObj err: " + err.Error())
			result.ParamErrorResult(r, w, stripeErr)
			return
		}
		var req types.StripeCheckoutWebhookReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		if errMsg, errCode := validator.Validate(req); errCode != 0 {
			result.ParamErrorResult(r, w, errors.New(errMsg))
			return
		}

		l := stirpe_checkout.NewStripeCheckoutWebhookLogic(r.Context(), svcCtx)
		resp, err := l.StripeCheckoutWebhook(&req, sessionObj)
		result.HttpResult(r, w, resp, err)
	}
}
