package copytrading

type CopytradingTraderAddMarginRequest struct {
	ContractCode string `json:"contract_code"`
	Amount       string `json:"amount"`
	Type         int    `json:"type"`
	PositionSide string `json:"position_side"`
}
