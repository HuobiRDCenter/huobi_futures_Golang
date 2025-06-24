package order

type SwapTradeOrderResponse struct {
	Code string `json:"code"`

	Data []struct {
		OrderId       string `json:"order_id"`
		ClientOrderId string `json:"client_order_id"`
		CancelReason  string `json:"cancel_reason"`
	} `json:"data,omitempty"`

	Message int `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
