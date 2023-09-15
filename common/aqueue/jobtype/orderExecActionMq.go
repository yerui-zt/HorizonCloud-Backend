package jobtype

import (
	"encoding/json"
	"github.com/hibiken/asynq"
)

const OrderExecActionMqPrefix = "order:action:exec"

type OrderExecActionMqPayload struct {
	OrderId int64 `json:"order_id"`
}

func NewOrderExecActionMq(orderId int64) (*asynq.Task, error) {
	payload, err := json.Marshal(OrderExecActionMqPayload{OrderId: orderId})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(OrderExecActionMqPrefix, payload), nil
}
