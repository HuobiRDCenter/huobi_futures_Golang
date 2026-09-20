package order

type SwapTradeAllOrdersResponse struct {
	Code int `json:"code"`

	Data []struct {
		OrderId       string `json:"order_id"`
		ClientOrderId string `json:"client_order_id"`
	} `json:"data,omitempty"`

	Message string `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
