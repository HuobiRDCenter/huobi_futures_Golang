package order

type GetTradePositionHistoryResponse struct {
	Code string `json:"code"`

	Data []struct {
		ID             string  `json:"id"`
		ContractCode   string  `json:"contract_code"`
		PositionSide   string  `json:"position_side"`
		MarginMode     string  `json:"margin_mode"`
		CostOpen       string  `json:"cost_open"`
		CloseAvgPrice  string  `json:"close_avg_price"`
		Volume         string  `json:"volume"`
		Available      string  `json:"available"`
		LeverRate      string  `json:"lever_rate"`
		Profit         string  `json:"profit"`
		ProfitRatio    string  `json:"profit_ratio"`
		State          string  `json:"state"`
		Fee            string  `json:"fee"`
		FundingFee     string  `json:"funding_fee"`
		MarginCurrency string  `json:"margin_currency"`
		Last           float64 `json:"last"`
		BusinessType   string  `json:"business_type"`
		ContractType   string  `json:"contract_type"`
		CreatedTime    string  `json:"created_time"`
		UpdatedTime    string  `json:"updated_time"`
	} `json:"data,omitempty"`

	Message int `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
