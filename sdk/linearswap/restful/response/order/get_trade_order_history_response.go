package order

type GetTradeOrderHistoryResponse struct {
	Code int `json:"code"`

	Data []struct {
		ID                 string `json:"id"`
		ContractCode       string `json:"contract_code"`
		Side               string `json:"side"`
		PositionSide       string `json:"position_side"`
		Type               string `json:"type"`
		PriceMatch         string `json:"price_match"`
		OrderID            string `json:"order_id"`
		ClientOrderID      string `json:"client_order_id"`
		MarginMode         string `json:"margin_mode"`
		Price              string `json:"price"`
		CancelReason       string `json:"cancel_reason"`
		Volume             string `json:"volume"`
		LeverRate          int64  `json:"lever_rate"`
		State              string `json:"state"`
		CanceledSource     string `json:"canceled_source"`
		OrderSource        string `json:"order_source"`
		ReduceOnly         bool   `json:"reduce_only"`
		TimeInForce        string `json:"time_in_force"`
		TpTriggerPrice     string `json:"tp_trigger_price"`
		TpOrderPrice       string `json:"tp_order_price"`
		TpType             string `json:"tp_type"`
		TpTriggerPriceType string `json:"tp_trigger_price_type"`
		SlTriggerPrice     string `json:"sl_trigger_price"`
		SlOrderPrice       string `json:"sl_order_price"`
		SlType             string `json:"sl_type"`
		SlTriggerPriceType string `json:"sl_trigger_price_type"`
		TradeAvgPrice      string `json:"trade_avg_price"`
		TradeVolume        string `json:"trade_volume"`
		TradeTurnover      string `json:"trade_turnover"`
		FeeCurrency        string `json:"fee_currency"`
		Fee                string `json:"fee"`
		DeductionCurrency  string `json:"deduction_currency"`
		DeductionAmount    string `json:"deduction_amount"`
		Profit             string `json:"profit"`
		ContractType       string `json:"contract_type"`
		CreatedTime        string `json:"created_time"`
		UpdatedTime        string `json:"updated_time"`
	} `json:"data,omitempty"`

	Message string `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
