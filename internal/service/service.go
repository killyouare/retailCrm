package service

import (
	"context"
	"crmtest/internal/domain/order"
	"crmtest/internal/pkg/retail"
	logger "crmtest/internal/pkg/util"
	"time"
)

var dtLayout = "2006-01-02 15:04:05"

type Service interface {
	GetOrders(ctx context.Context, ordersRequest *retail.OrdersRequest) ([]order.Order, error)
}

type service struct {
	logger logger.Logger
	retail retail.Client
}

func NewGetOrderService(logger logger.Logger, retail retail.Client) (Service, error) {
	return &service{
		logger: logger,
		retail: retail,
	}, nil
}

func (s *service) GetOrders(ctx context.Context, ordersRequest *retail.OrdersRequest) ([]order.Order, error) {
	orderResponse, err := s.retail.GetOrders(ctx, ordersRequest)

	if err != nil {
		s.logger.Info("Get orders err %v", err)

		return nil, err
	}

	var orders []order.Order
	var orderIds []int
	for _, o := range orderResponse.Orders {
		createdAt, _ := time.Parse(dtLayout, o.CreatedAt)
		orders = append(orders, order.Order{
			OrderId: o.ID,
			//айдишник клиента не отдается на тестовом аккаунте
			ClientId:  o.Customer.ID,
			Number:    o.Number,
			CreatedAt: createdAt,
			TotalSum:  int64(o.TotalSumm * 100),
			PrepaySum: int64(o.PrepaySum * 100),
		})
		orderIds = append(orderIds, o.ID)
	}

	s.logger.Info("Получены айди заказов %v", orderIds)

	return orders, nil
}
