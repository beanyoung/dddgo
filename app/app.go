package app

import (
	"github.com/beanyoung/dddgo/rpc/order"
)

type orderServer struct {
	order.UnimplementedOrderServer
}

func NewOrderServer() order.OrderServer {
	return &orderServer{}
}

var _ order.OrderServer = (*orderServer)(nil)
