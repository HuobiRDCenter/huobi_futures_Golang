package order

type GetTradeOrderResponse struct {
	Code int `json:"code"`

	// V5 查询订单：data 为**单个订单对象**，不是数组
	Data struct {
		ID                 string `json:"id,omitempty"`
		ContractCode       string `json:"contract_code,omitempty"`
		Side               string `json:"side,omitempty"`
		PositionSide       string `json:"position_side,omitempty"`
		Type               string `json:"type,omitempty"`
		PriceMatch         string `json:"price_match,omitempty"`
		OrderID            string `json:"order_id,omitempty"`
		ClientOrderID      string `json:"client_order_id,omitempty"`
		MarginMode         string `json:"margin_mode,omitempty"`
		Price              string `json:"price,omitempty"`
		Volume             string `json:"volume,omitempty"`
		LeverRate          int64  `json:"lever_rate,omitempty"`
		State              string `json:"state,omitempty"`
		OrderSource        string `json:"order_source,omitempty"`
		ReduceOnly         bool   `json:"reduce_only,omitempty"`
		TimeInForce        string `json:"time_in_force,omitempty"`
		TpTriggerPrice     string `json:"tp_trigger_price,omitempty"`
		TpOrderPrice       string `json:"tp_order_price,omitempty"`
		TpType             string `json:"tp_type,omitempty"`
		TpTriggerPriceType string `json:"tp_trigger_price_type,omitempty"`
		SlTriggerPrice     string `json:"sl_trigger_price,omitempty"`
		SlOrderPrice       string `json:"sl_order_price,omitempty"`
		SlType             string `json:"sl_type,omitempty"`
		SlTriggerPriceType string `json:"sl_trigger_price_type,omitempty"`
		TradeAvgPrice      string `json:"trade_avg_price,omitempty"`
		TradeVolume        string `json:"trade_volume,omitempty"`
		TradeTurnover      string `json:"trade_turnover,omitempty"`
		FeeCurrency        string `json:"fee_currency,omitempty"`
		Fee                string `json:"fee,omitempty"`
		PriceProtect       bool   `json:"price_protect,omitempty"`
		Profit             string `json:"profit,omitempty"`
		ContractType       string `json:"contract_type,omitempty"`
		// 线上返回的是字符串毫秒时间戳："created_time":"1789880078388"
		CreatedTime string `json:"created_time,omitempty"`
		UpdatedTime string `json:"updated_time,omitempty"`
	} `json:"data,omitempty"`

	Message string `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
