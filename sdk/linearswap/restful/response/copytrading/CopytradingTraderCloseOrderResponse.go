package copytrading

type CopytradingTraderCloseOrderResponse struct {
	Tid  string `json:"tid"`
	Data bool   `json:"data"`
	Code string `json:"code"`
}
