package order

// SwapTradePositionAllResponse POST /v5/trade/position_all（全平）的应答。
//
// 实测（2026-09-20）：
//
//	有持仓 → {"code":200,"data":[{"code":200,"message":"Success",
//	           "order_id":"1551218190814420992","client_order_id":"..."}]}
//	无持仓 → {"code":200,"data":[]}
//
// 即 data 是数组，元素自带 code/message（同批量下单的口径）。
type SwapTradePositionAllResponse struct {
	Code int `json:"code"`

	Data []struct {
		Code          int            `json:"code"`
		Message       string         `json:"message"`
		OrderId       FlexibleString `json:"order_id"`
		ClientOrderId string         `json:"client_order_id"`
	} `json:"data,omitempty"`

	Message string `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
