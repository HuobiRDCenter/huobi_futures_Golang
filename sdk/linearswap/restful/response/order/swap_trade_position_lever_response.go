package order

type SwapTradePositionLeverResponse struct {
	Code string `json:"code"`

	Data []struct {
		ContractCode string `json:"contract_code"`
		MarginMode   string `json:"margin_mode"`
		LeverRate    string `json:"lever_rate"`
	} `json:"data,omitempty"`

	Message int `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
