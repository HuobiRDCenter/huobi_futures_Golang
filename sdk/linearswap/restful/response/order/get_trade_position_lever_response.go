package order

type GetTradePositionLeverResponse struct {
	Code string `json:"code"`

	Data []struct {
		ContractCode string `json:"contract_code"`
		ContractType string `json:"contract_type"`
		MarginMode   string `json:"margin_mode"`
		// 因为position_side是可选填的，所以使用指针类型，这样可以赋值为nil表示未填写
		PositionSide   *string `json:"position_side"`
		LeverRate      string  `json:"lever_rate"`
		AvailableLever string  `json:"available_lever"`
	} `json:"data,omitempty"`

	Message int `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
