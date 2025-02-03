package command

import (
	"context"
)

type Command interface {
	Exec(ctx context.Context) error
	withArgs(args ...string) error
}

func Parse(injection *OrderDep, args ...string) (Command, error) {
	//здесь в зависимости от аргументов, выбирается реализация command
	cmd := &getOrders{dep: injection}

	err := cmd.withArgs(args...)

	if err != nil {
		return nil, err
	}

	return cmd, nil
}
