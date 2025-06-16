package copytrading

type CopytradingTraderPlaceOrderRequest struct {
	// 交易对
	ContractCode string `json:"contract_code"`
	// 委托价格
	Price string `json:"price,omitempty"`
	// 委托数量
	Amount string `json:"amount"`
	// 保证金模式 isolated：逐仓模式，cross：全仓模式
	MarginMode string `json:"margin_mode"`
	// 下单价格类型 1:限价、2:限价-对手价、6:限价-最优20档、8:fok、13:fok-对手价、16:最优20档-FOK下单、17:市价单
	OrderPriceType int `json:"order_price_type"`
	// 买卖方向类型 1-buy,2-sell
	OrderDirection int `json:"order_direction"`
	// 杠杆倍数
	Lever int `json:"lever"`
	// 止盈价格
	TpTriggerPrice string `json:"tp_trigger_price,omitempty"`
	// 止损价格
	SlTriggerPrice string `json:"sl_trigger_price,omitempty"`
}
