package order

type SwapTradeOrderResponse struct {
	Code int `json:"code"`

	// V5 单笔下单的 data 是**对象**：{"order_id":"...","client_order_id":"..."}
	// （批量下单才是数组，见 SwapTradeBatchOrderResponse）
	Data struct {
		OrderId       string `json:"order_id"`
		ClientOrderId string `json:"client_order_id"`
		CancelReason  string `json:"cancel_reason"`
	} `json:"data,omitempty"`

	Message string `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
