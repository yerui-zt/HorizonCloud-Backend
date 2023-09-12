package logic

import (
	"HorizonX/common/xerr"
	"HorizonX/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"

	"HorizonX/rpc/payment/internal/svc"
	"HorizonX/rpc/payment/payment"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePaymentLogic {
	return &CreatePaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreatePayment 创建支付, 生成支付链接
func (l *CreatePaymentLogic) CreatePayment(in *payment.CreatePaymentReq) (*payment.CreatePaymentResp, error) {
	findOrder, err := l.svcCtx.OrderModel.FindOneByOrderNo(l.ctx, in.OrderNo)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ORDER_NOT_FOUND), "Order not found [%s] [%s]", in.OrderNo, err.Error())
	}
	findUser, err := l.svcCtx.UserModel.FindOne(l.ctx, nil, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_FOUND_ERROR), "unknown user [%d] [s]", in.UserId, err.Error())
	}
	findMethod, err := l.svcCtx.SystemPaymentMethodModel.FindOneByName(l.ctx, in.Method)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.PAYMENT_METHOD_NOT_FOUND), "payment method not found [%s] [%s]", in.Method, err.Error())
	}

	var url string
	switch findMethod.Type {
	case "stripe_checkout":
		url, err = l.createStripeCheckoutPayment(findOrder, findUser, findMethod)
	case "alipay":
	}

	if err != nil {
		return nil, err
	} else {
		resp := &payment.CreatePaymentResp{
			Url: url,
		}
		return resp, nil
	}
}

// createStripeCheckoutPayment 创建Stripe Checkout支付
func (l *CreatePaymentLogic) createStripeCheckoutPayment(o *model.Order, u *model.User, method *model.SystemPaymentMethod) (string, error) {
	credentialStr := method.Credential
	credential := payment.StripeCheckoutCredential{}
	err := json.Unmarshal([]byte(credentialStr), &credential)
	if err != nil {
		return "", errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "unmarshal credential failed [%s] [%s]", credentialStr, err.Error())
	}

	stripe.Key = credential.SecKey
	// todo 设置成功和取消的url
	domain := "https://horizonx.app"
	params := &stripe.CheckoutSessionParams{
		CustomerEmail: stripe.String(u.Email),
		//InvoiceCreation: &stripe.CheckoutSessionInvoiceCreationParams{
		//	Enabled: stripe.Bool(true),
		//},
		ClientReferenceID: stripe.String(o.OrderNo),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
			"alipay",
		}),
		//PaymentMethodOptions: &stripe.CheckoutSessionPaymentMethodOptionsParams{
		//	WeChatPay: &stripe.CheckoutSessionPaymentMethodOptionsWeChatPayParams{
		//		Client: stripe.String("web"),
		//	},
		//},
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("USD"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(fmt.Sprintf("%s - Order #%s", "HorizonX", o.OrderNo)),
						//Description: stripe.String(fmt.Sprintf("Order #%s", s.Order.OrderNo)),
					},
					UnitAmount: stripe.Int64(o.TotalAmount),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		// /api/payment/:payment_method_type/:unique/success

		SuccessURL: stripe.String(fmt.Sprintf("%s/api/payment/%s/%s/success", domain, method.Type, method.Path)),
		CancelURL:  stripe.String(fmt.Sprintf("%s/api/payment/%s/%s/cancel", domain, method.Type, method.Path)),
	}
	stripeSession, err := session.New(params)
	if err != nil {
		return "", errors.Wrapf(xerr.NewErrCode(xerr.PAYMENT_CREATE_ERROR), "create stripe checkout session failed [%s]", err.Error())
	}
	return stripeSession.URL, nil

}
