package command

import (
	"context"
	"crmtest/internal/action"
)

type getOrders struct {
	dep *OrderDep
}

type OrderDep struct {
	A action.Action
}

func (o *getOrders) withArgs(args ...string) error {
	//TODO: реализовать передачу фильтров и тд
	return nil
}

func (o *getOrders) Exec(ctx context.Context) error {
	return o.dep.A.GetOrders(ctx, &action.GetOrdersDTO{})
}
