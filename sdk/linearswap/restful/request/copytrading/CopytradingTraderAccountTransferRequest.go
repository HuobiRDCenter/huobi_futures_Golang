package copytrading

type CopytradingTraderAccountTransferRequest struct {
	Amount   string `json:"amount"`
	Type     int    `json:"type"`
	Currency string `json:"currency"`
}
