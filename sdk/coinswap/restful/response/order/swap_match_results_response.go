package order

type SwapMatchResultsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ts   int64  `json:"ts"`
	Data []struct {
		QueryID          int64   `json:"query_id"`
		ID               string  `json:"id"`
		MatchID          int64   `json:"match_id"`
		OrderID          int64   `json:"order_id"`
		OrderIDStr       string  `json:"order_id_str"`
		Symbol           string  `json:"symbol"`
		ContractCode     string  `json:"contract_code"`
		Direction        string  `json:"direction"`
		Offset           string  `json:"offset"`
		TradeVolume      float64 `json:"trade_volume"`
		TradePrice       float64 `json:"trade_price"`
		TradeTurnover    float64 `json:"trade_turnover"`
		CreateDate       int64   `json:"create_date"`
		OffsetProfitLoss float64 `json:"offset_profitloss"`
		TradeFee         float64 `json:"trade_fee"`
		Role             string  `json:"role"`
		RealProfit       float64 `json:"real_profit"`
		FeeAsset         string  `json:"fee_asset"`
		OrderSource      string  `json:"order_source"`
	} `json:"data"`
}
