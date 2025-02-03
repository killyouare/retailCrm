package retail

type OrdersResponse struct {
	Success    bool        `json:"success"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Orders     []Order     `json:"orders,omitempty"`
}
