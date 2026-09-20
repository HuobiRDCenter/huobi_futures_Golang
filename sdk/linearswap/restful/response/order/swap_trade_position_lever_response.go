package order

type SwapTradePositionLeverResponse struct {
	Code int `json:"code"`

	// V5 设置杠杆：data 为对象 {"contract_code":"...","margin_mode":"...","lever_rate":13}
	Data struct {
		ContractCode string `json:"contract_code"`
		MarginMode   string `json:"margin_mode"`
		LeverRate    int64  `json:"lever_rate"`
	} `json:"data,omitempty"`

	Message string `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
