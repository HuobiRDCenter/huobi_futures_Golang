package copytrading

type CopytradingTraderPositionPositionsResponse struct {
	Tid  string                             `json:"tid"`
	Data CopytradingTraderPositionPositions `json:"data"`
	Code int64                              `json:"code"`
}

type CopytradingTraderPositionPositions struct {
	Positions []PositionPosition `json:"positions"`
	QueryID   int64              `json:"query_id"`
}

type PositionPosition struct {
	SubPositionID string `json:"sub_position_id"`
	MarginMode    string `json:"margin_mode"`
	PositionSide  string `json:"position_side"`
	Lever         string `json:"lever"`
	OpenOrderID   string `json:"open_order_id"`
	OpenAvgPrice  string `json:"open_avg_price"`
	OpenTime      int64  `json:"open_time"`
	Volume        string `json:"volume"`
	CloseTime     int64  `json:"close_time"`
	CloseAvgPrice string `json:"close_avg_price"`
	OpenFee       string `json:"open_fee"`
	CloseFee      string `json:"close_fee"`
	Profit        string `json:"profit"`
	ProfitRate    string `json:"profit_rate"`
	CloseType     int    `json:"close_type"`
	FollowTakes   string `json:"follow_takes"`
}
