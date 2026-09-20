package order

// GetTradeOrderTradesResponse V5 /v5/trade/order/details 查询订单成交明细。
// data 为成交数组，无成交时可为 null/[]，故用切片承接。
// 口径对齐 Java TradeOrderTradesResponse（List<DataBean>）与
// C++ response/order/GetOrderTradesResponse.hpp（unique_ptr<vector<Data>>）。
//
// 实测样例（2026-09-20）：
//
//	{"code":200,"data":[{"id":"1551218221931384832","order_id":"1551218221863182336",
//	 "trade_id":"652482","trade_price":"7632.2","trade_volume":"1","created_time":"1789880844338",...}]}
type GetTradeOrderTradesResponse struct {
	Code int `json:"code"`

	Data []GetTradeOrderTradesData `json:"data,omitempty"`

	Message string `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}

type GetTradeOrderTradesData struct {
	Id            string `json:"id,omitempty"`
	ContractCode  string `json:"contract_code,omitempty"`
	OrderId       string `json:"order_id,omitempty"`
	TradeId       string `json:"trade_id,omitempty"`
	Side          string `json:"side,omitempty"`
	PositionSide  string `json:"position_side,omitempty"`
	OrderType     string `json:"order_type,omitempty"`
	MarginMode    string `json:"margin_mode,omitempty"`
	Type          string `json:"type,omitempty"`
	ClientOrderId string `json:"client_order_id,omitempty"`
	Role          string `json:"role,omitempty"`
	TradePrice    string `json:"trade_price,omitempty"`
	TradeVolume   string `json:"trade_volume,omitempty"`
	TradeTurnover string `json:"trade_turnover,omitempty"`
	// 线上返回的是字符串毫秒时间戳："created_time":"1789880844338"（同订单详情接口）
	CreatedTime       string `json:"created_time,omitempty"`
	UpdatedTime       string `json:"updated_time,omitempty"`
	OrderSource       string `json:"order_source,omitempty"`
	FeeCurrency       string `json:"fee_currency,omitempty"`
	TradeFee          string `json:"trade_fee,omitempty"`
	DeductionCurrency string `json:"deduction_currency,omitempty"`
	DeductionAmount   string `json:"deduction_amount,omitempty"`
	DeductionPrice    string `json:"deduction_price,omitempty"`
	Profit            string `json:"profit,omitempty"`
	ContractType      string `json:"contract_type,omitempty"`
}
