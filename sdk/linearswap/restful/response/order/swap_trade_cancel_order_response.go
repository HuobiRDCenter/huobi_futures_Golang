package order

type SwapTradeCancelOrderResponse struct {
	Code int `json:"code"`

	// V5 单笔撤单的 data 是**对象**：{"order_id":"...","client_order_id":"..."}
	// （批量撤单才是数组，见 SwapTradeCancelBatchOrdersResponse）
	Data struct {
		OrderId       string `json:"order_id"`
		ClientOrderId string `json:"client_order_id"`
	} `json:"data,omitempty"`

	Message string `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
