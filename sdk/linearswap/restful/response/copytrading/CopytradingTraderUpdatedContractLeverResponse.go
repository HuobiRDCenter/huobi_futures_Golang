package copytrading

type CopytradingTraderUpdatedContractLeverResponse struct {
	Tid  string                               `json:"tid"`
	Data CopytradingTraderUpdatedContractLeve `json:"data"`
	Code int64                                `json:"code"`
}

type CopytradingTraderUpdatedContractLeve struct {
	ContractCode string `json:"contract_code"`
	Lever        int    `json:"lever"`
	MarginMode   string `json:"margin_mode"`
}
