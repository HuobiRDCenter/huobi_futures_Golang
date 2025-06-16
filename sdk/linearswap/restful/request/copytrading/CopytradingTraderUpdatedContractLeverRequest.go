package copytrading

type CopytradingTraderUpdatedContractLeverRequest struct {
	ContractCode string `json:"contract_code"`
	Lever        int    `json:"lever"`
	MarginMode   string `json:"margin_mode"`
}
