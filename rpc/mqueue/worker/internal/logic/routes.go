package logic

import (
	"HorizonX/common/aqueue/jobtype"
	"HorizonX/rpc/mqueue/worker/internal/svc"
	"context"
	"github.com/hibiken/asynq"
)

type MqWorker struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMqWorker(ctx context.Context, svcCtx *svc.ServiceContext) *MqWorker {
	return &MqWorker{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MqWorker) Register() *asynq.ServeMux {

	mux := asynq.NewServeMux()

	mux.Handle(jobtype.OrderExecActionMqPrefix, NewOrderExecActionMqHandler(l.svcCtx))

	return mux
}
