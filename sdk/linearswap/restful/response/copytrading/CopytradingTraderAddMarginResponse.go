package copytrading

type CopytradingTraderAddMarginResponse struct {
	Tid  string `json:"tid"`
	Data struct {
		ContractCode string `json:"contract_code"`
		Amount       string `json:"amount"`
		Type         int    `json:"type"`
		PositionSide string `json:"position_side"`
	} `json:"data"`
	Code int64 `json:"code"`
}
