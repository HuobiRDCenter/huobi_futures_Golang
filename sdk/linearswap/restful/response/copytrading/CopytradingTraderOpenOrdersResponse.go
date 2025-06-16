package copytrading

type CopytradingTraderOpenOrdersResponse struct {
	Tid  string                        `json:"tid"`
	Data []CopytradingTraderOpenOrders `json:"data"`
	Code int64                         `json:"code"`
}

type CopytradingTraderOpenOrders struct {
	// 根据实际数据结构添加字段
	ContractCode   string `json:"contract_code"`
	Price          string `json:"price"`
	Volume         string `json:"volume"`
	MarginMode     string `json:"margin_mode"`
	PositionSide   string `json:"position_side"`
	OrderDirection int    `json:"order_direction"`
	Lever          int    `json:"lever"`
	AvgPrice       string `json:"avg_price"`
	AvgVolume      string `json:"avg_volume"`
	Fee            string `json:"fee"`
}
