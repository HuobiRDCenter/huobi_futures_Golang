package copytrading

type CopytradingTraderCloseAllPositionRequest struct {
	ContractCode string `json:"contract_code"`
	MarginMode   string `json:"margin_mode"`
	PositionSide string `json:"position_side"`
}
