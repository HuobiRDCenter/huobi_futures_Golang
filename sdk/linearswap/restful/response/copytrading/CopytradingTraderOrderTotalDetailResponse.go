package copytrading

type CopytradingTraderOrderTotalDetailResponse struct {
	Tid  string                            `json:"tid"`
	Data CopytradingTraderOrderTotalDetail `json:"data"`
	Code int64                             `json:"code"`
}

type CopytradingTraderOrderTotalDetail struct {
	TotalFollowerNum   string              `json:"total_follower_num"`
	CurrentFollowerNum string              `json:"current_follower_num"`
	TotalPl            string              `json:"total_pl"`
	WinRate            string              `json:"win_rate"`
	Last24HoursProfit  []Last24HoursProfit `json:"last_24_hours_profit"`
	Last90DaysProfit   []Last90DaysProfit  `json:"last_90_days_profit"`
}

type Last24HoursProfit struct {
	Rate   string `json:"rate"`
	Amount string `json:"amount"`
	Ts     string `json:"ts"`
}

type Last90DaysProfit struct {
	Rate   string `json:"rate"`
	Amount string `json:"amount"`
	Ts     string `json:"ts"`
}
