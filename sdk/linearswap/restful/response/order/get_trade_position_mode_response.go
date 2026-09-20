package order

type GetTradePositionModeResponse struct {
	Code int `json:"code"`

	// V5 查询持仓模式：data 为**对象** {"position_mode":"dual_side"}
	Data struct {
		PositionMode string `json:"position_mode"`
	} `json:"data,omitempty"`

	Message string `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
