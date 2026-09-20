package order

// SwapTradePositionResponse POST /v5/trade/position（按方向全平）的应答。
//
// 实测（2026-09-20）：data 是**单个对象**而非数组：
//
//	{"code":200,"message":"Success",
//	 "data":{"order_id":1551218221863182336,"client_order_id":"1551218221863182336"}}
//
// 注意 order_id 这里是数字（同族接口 /v5/trade/order 返回字符串），故用 FlexibleString 承接。
type SwapTradePositionResponse struct {
	Code int `json:"code"`

	Data struct {
		OrderId       FlexibleString `json:"order_id"`
		ClientOrderId string         `json:"client_order_id"`
	} `json:"data,omitempty"`

	Message string `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
