package action

import (
	"context"
	"crmtest/internal/pkg/retail"
	logger "crmtest/internal/pkg/util"
	"crmtest/internal/repository"
	"crmtest/internal/service"
)

type Action interface {
	GetOrders(ctx context.Context, getOrdersDTO *GetOrdersDTO) error
}

type action struct {
	logger     logger.Logger
	repository repository.Repository
	service    service.Service
}

func NewGetOrdersAction(logger logger.Logger, repository repository.Repository, service service.Service) (Action, error) {
	return &action{
		logger:     logger,
		repository: repository,
		service:    service,
	}, nil
}

// GetOrdersDTO TODO: в формат запроса с учетом фильтров
type GetOrdersDTO struct {
}

func (getOrderAction *action) GetOrders(ctx context.Context, getOrdersDTO *GetOrdersDTO) error {
	//TODO: использовать пагинацию и качать все заказы
	orders, err := getOrderAction.service.GetOrders(ctx, &retail.OrdersRequest{
		Limit: 20,
	})

	if err != nil {
		return err
	}

	if err = getOrderAction.repository.CreateOrders(ctx, orders); err != nil {
		return err
	}

	return nil
}
