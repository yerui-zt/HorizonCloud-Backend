package jobtype

import (
	"HorizonX/model"
	"encoding/json"
	"github.com/hibiken/asynq"
	"time"
)

const OrderExecActionMqPrefix = "order:action:exec"

type OrderExecActionMqPayload struct {
	Item *model.OrderItem
}

func NewOrderExecActionMq(item *model.OrderItem) (*asynq.Task, error) {
	payload, err := json.Marshal(OrderExecActionMqPayload{Item: item})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(OrderExecActionMqPrefix, payload, asynq.MaxRetry(0), asynq.Timeout(5*time.Minute)), nil
}
