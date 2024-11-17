package order

type GetTradePositionRiskLimitResponse struct {
	Code string `json:"code"`

	Data []struct {
		ContractCode          string `json:"contract_code"`
		ContractType          string `json:"contract_type"`
		MarginMode            string `json:"margin_mode"`
		PositionSide          string `json:"position_side"`
		MaxLever              string `json:"max_lever"`
		MaintenanceMarginRate string `json:"maintenance_margin_rate"`
		MaxVolume             string `json:"max_volume"`
		MinVolume             string `json:"min_volume"`
	} `json:"data,omitempty"`

	Message int `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
