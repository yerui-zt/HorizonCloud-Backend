package stirpe_checkout

import (
	"HorizonX/common/xerr"
	"HorizonX/rpc/order/order"
	"HorizonX/rpc/payment/payment"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/webhook"

	"HorizonX/api/internal/svc"
	"HorizonX/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StripeCheckoutWebhookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStripeCheckoutWebhookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StripeCheckoutWebhookLogic {
	return &StripeCheckoutWebhookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StripeCheckoutWebhookLogic) StripeCheckoutWebhook(req *types.StripeCheckoutWebhookReq, sessionObj []byte) (resp *types.StripeCheckoutWebhookResp, err error) {
	method, err := l.svcCtx.SystemPaymentMethodModel.FindOneByPath(l.ctx, req.UniquePath)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(400, "invalid payment method"), "FindOneByPath failed: %v", err)
	}

	credentialStr := method.Credential
	credential := new(payment.StripeCheckoutCredential)
	err = json.Unmarshal([]byte(credentialStr), &credential)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(500, "invalid stripe credential"), "json.Unmarshal failed: %v", err)
	}

	endpointSecret := credential.WebhookSecKey
	var event stripe.Event
	if l.svcCtx.Config.Mode == "dev" {
		event, err = webhook.ConstructEventWithOptions(sessionObj, req.StripeSignature, endpointSecret,
			webhook.ConstructEventOptions{
				IgnoreAPIVersionMismatch: true,
			})
	} else {
		event, err = webhook.ConstructEvent(sessionObj, req.StripeSignature, endpointSecret)
	}
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(400, "invalid stripe signature"), "ConstructEvent failed: %v", err)
	}

	switch event.Type {
	case "checkout.session.completed":
		var session stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &session)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCodeMsg(500, "invalid stripe session"), "json.Unmarshal failed: %v", err)
		}
		orderNo := session.ClientReferenceID
		callBackNo := session.PaymentIntent.ID

		_, err = l.svcCtx.OrderRPC.FullFillOrder(l.ctx, &order.FullFillOrderReq{
			OrderNo:    orderNo,
			CallbackNo: callBackNo,
			Method:     method.Name,
		})
		if err != nil {
			return nil, err
		}
	}

	return &types.StripeCheckoutWebhookResp{
		Msg: "success",
	}, nil
}
