package copytrading

type CopytradingTraderCurrentPositionsResponse struct {
	Tid  string                            `json:"tid"`
	Data CopytradingTraderCurrentPositions `json:"data"`
	Code int64                             `json:"code"`
}

type CopytradingTraderCurrentPositions struct {
	Positions []CurrentPosition `json:"positions"`
	QueryID   int64             `json:"query_id"`
}

type CurrentPosition struct {
	SubPositionID    string `json:"sub_position_id"`
	MarginMode       string `json:"margin_mode"`
	PositionSide     string `json:"position_side"`
	Lever            string `json:"lever"`
	OpenOrderID      string `json:"open_order_id"`
	OpenAvgPrice     string `json:"open_avg_price"`
	OpenFee          string `json:"open_fee"`
	OpenTime         int64  `json:"open_time"`
	PositionMargin   string `json:"position_margin"`
	Volume           string `json:"volume"`
	LiquidationPrice string `json:"liquidation_price"`
	TpTriggerPrice   string `json:"tp_trigger_price"`
	SlTriggerPrice   string `json:"sl_trigger_price"`
}
