package retail

type OrdersRequest struct {
	Filter OrdersFilter `url:"filter,omitempty"`
	Limit  int          `url:"limit,omitempty"`
	Page   int          `url:"page,omitempty"`
}
