package order

type GetTradePositionOpensResponse struct {
	Code string `json:"code"`

	Data []struct {
		ContractCode      string  `json:"contract_code"`
		PositionSide      string  `json:"position_side"`
		MarginMode        string  `json:"margin_mode"`
		CostOpen          string  `json:"cost_open"`
		Volume            string  `json:"volume"`
		Available         string  `json:"available"`
		LeverRate         string  `json:"lever_rate"`
		AdlRiskPercent    string  `json:"adl_risk_percent"`
		LiquidationPrice  string  `json:"liquidation_price"`
		InitialMargin     string  `json:"initial_margin"`
		MaintenanceMargin string  `json:"maintenance_margin"`
		Margin            string  `json:"margin"`
		ProfitUnreal      string  `json:"profit_unreal"`
		ProfitRate        string  `json:"profit_rate"`
		MarginRate        string  `json:"margin_rate"`
		MarginCurrency    string  `json:"margin_currency"`
		PositionMode      string  `json:"position_mode"`
		Last              float64 `json:"last"`
		ContractType      string  `json:"contract_type"`
		CreatedTime       string  `json:"created_time"`
		UpdatedTime       string  `json:"updated_time"`
	} `json:"data,omitempty"`

	Message int `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
