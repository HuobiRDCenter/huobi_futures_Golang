package copytrading

type CopytradingTraderQueryContractLeverResponse struct {
	Tid  string `json:"tid"`
	Data struct {
		ContractCode     string `json:"contract_code"`
		MarginMode       string `json:"margin_mode"`
		LeverRateRange   int    `json:"lever_rate_range"`
		CurrentLeverRate int    `json:"current_lever_rate"`
	} `json:"data"`
	Code int64 `json:"code"`
}
