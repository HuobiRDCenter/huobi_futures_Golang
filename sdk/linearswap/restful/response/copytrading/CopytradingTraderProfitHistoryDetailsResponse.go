package copytrading

type CopytradingTraderProfitHistoryDetailsResponse struct {
	Tid  string                                `json:"tid"`
	Data CopytradingTraderProfitHistoryDetails `json:"data"`
	Code int64                                 `json:"code"`
}

type CopytradingTraderProfitHistoryDetails struct {
	ProfitItems []ProfitItem `json:"profit_items"`
	QueryID     int64        `json:"query_id"`
}

type ProfitItem struct {
	UserID         string `json:"user_id"`
	ProfitCurrency string `json:"profit_currency"`
	ProfitAmount   string `json:"profit_amount"`
	ProfitTime     string `json:"profit_time"`
}
