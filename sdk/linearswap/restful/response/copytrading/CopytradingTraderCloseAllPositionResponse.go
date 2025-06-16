package copytrading

type CopytradingTraderCloseAllPositionResponse struct {
	Tid  string `json:"tid"`
	Data bool   `json:"data"`
	Code int64  `json:"code"`
}
