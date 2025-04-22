package order

type SwapTradeCancelOrderResponse struct {
	Code string `json:"code"`

	Data []struct {
		OrderId       string `json:"order_id"`
		ClientOrderId string `json:"client_order_id"`
	} `json:"data,omitempty"`

	Message int `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
