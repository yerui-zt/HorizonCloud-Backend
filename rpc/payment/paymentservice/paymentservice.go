// Code generated by goctl. DO NOT EDIT.
// Source: payment.proto

package paymentservice

import (
	"context"

	"HorizonX/rpc/payment/payment"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreatePaymentReq  = payment.CreatePaymentReq
	CreatePaymentResp = payment.CreatePaymentResp

	PaymentService interface {
		CreatePayment(ctx context.Context, in *CreatePaymentReq, opts ...grpc.CallOption) (*CreatePaymentResp, error)
	}

	defaultPaymentService struct {
		cli zrpc.Client
	}
)

func NewPaymentService(cli zrpc.Client) PaymentService {
	return &defaultPaymentService{
		cli: cli,
	}
}

func (m *defaultPaymentService) CreatePayment(ctx context.Context, in *CreatePaymentReq, opts ...grpc.CallOption) (*CreatePaymentResp, error) {
	client := payment.NewPaymentServiceClient(m.cli.Conn())
	return client.CreatePayment(ctx, in, opts...)
}