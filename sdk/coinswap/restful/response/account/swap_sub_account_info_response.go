package account

type SwapSubAccountInfoResponse struct {
	Status string `json:"status"`
	Ts     int64  `json:"ts"`
	Data   []struct {
		Symbol            string  `json:"symbol"`
		ContractCode      string  `json:"contract_code"`
		MarginBalance     float64 `json:"margin_balance"`
		MarginPosition    float64 `json:"margin_position"`
		MarginFrozen      float64 `json:"margin_frozen"`
		MarginAvailable   float64 `json:"margin_available"`
		ProfitUnreal      float64 `json:"profit_unreal"`
		RiskRate          float64 `json:"risk_rate"`
		LiquidationPrice  float64 `json:"liquidation_price"`
		WithdrawAvailable float64 `json:"withdraw_available"`
		LeverRate         int     `json:"lever_rate"`
		AdjustFactor      float64 `json:"adjust_factor"`
		MarginStatic      float64 `json:"margin_static"`
		NewRiskRate       string  `json:"new_risk_rate"`
		TradePartition    string  `json:"trade_partition"`
	} `json:"data,omitempty"`
}
