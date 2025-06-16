package copytrading

type CopytradingTraderTotalProfitHistoryResponse struct {
	Tid  string                              `json:"tid"`
	Data CopytradingTraderTotalProfitHistory `json:"data"`
	Code int64                               `json:"code"`
}

type CopytradingTraderTotalProfitHistory struct {
	Items []Item `json:"items"`
}

type Item struct {
	ProfitCurrency    string `json:"profit_currency"`
	TotalProfitAmount string `json:"total_profit_amont"`
}
