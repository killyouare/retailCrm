package order

import "time"

type Order struct {
	OrderId   int
	ClientId  int
	Number    string
	CreatedAt time.Time
	SendAt    *time.Time
	//Операции с интами надежнее, поэтому кастим к минимальным юнитамы
	TotalSum  int64
	PrepaySum int64
}
