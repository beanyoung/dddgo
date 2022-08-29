package app

import (
	"context"

	"github.com/google/uuid"

	"github.com/beanyoung/dddgo/rpc/order"
)

func (*orderServer) Consume(context.Context, *order.ConsumeRequest) (*order.ConsumeReply, error) {
	resp := &order.ConsumeReply{}
	resp.OrderId = uuid.New().String()
	return resp, nil
}
