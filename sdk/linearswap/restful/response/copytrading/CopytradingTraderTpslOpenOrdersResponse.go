package copytrading

type CopytradingTraderTpslOpenOrdersResponse struct {
	Tid  string                           `json:"tid"`
	Data []CopytradingTraderTpslOpenOrders `json:"data"`
	Code int64                            `json:"code"`
}

type CopytradingTraderTpslOpenOrders struct {
	// 这里根据实际data字段内容需要补充
	ContractCode  string `json:"contract_code"`
	Volume        string `json:"volume"`
	MarginMode    string `json:"margin_mode"`
	PositionSide  string `json:"position_side"`
	TriggerType   string `json:"trigger_type"`
	TpslOrderType string `json:"tpsl_order_type"`
	Price         string `json:"price"`
}
