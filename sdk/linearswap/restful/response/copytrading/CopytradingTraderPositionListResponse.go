package copytrading

type CopytradingTraderPositionListResponse struct {
	Tid  string                        `json:"tid"`
	Data CopytradingTraderPositionList `json:"data"`
	Code int64                         `json:"code"`
}

type CopytradingTraderPositionList struct {
	Positions []Position `json:"positions"`
}

type Position struct {
	MarginMode       string `json:"margin_mode"`
	PositionSide     string `json:"position_side"`
	Lever            string `json:"lever"`
	OpenAvgPrice     string `json:"open_avg_price"`
	PositionMargin   string `json:"position_margin"`
	MarginRate       string `json:"margin_rate"`
	Volume           string `json:"volume"`
	LiquidationPrice string `json:"liquidation_price"`
	UnrealProfit     string `json:"unreal_profit"`
	Profit           string `json:"profit"`
	ProfitRate       string `json:"profit_rate"`
	ContractCode     string `json:"contract_code"`
}
