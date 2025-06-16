package copytrading

type CopytradingTraderQueryContractResponse struct {
	Tid  string                           `json:"tid"`
	Data []CopytradingTraderQueryContract `json:"data"`
	Code int64                            `json:"code"`
}

type CopytradingTraderQueryContract struct {
	ContractCode string `json:"contract_code"`
}
