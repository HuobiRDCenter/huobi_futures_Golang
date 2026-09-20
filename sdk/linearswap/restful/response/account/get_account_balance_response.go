package account

// GetAccountBalanceResponse V5 GET /v5/account/balance。
//
// 实测（2026-09-20）：
//
//   - data 是**单个对象**（不是数组），details 才是数组
//
//   - 时间戳是**数字**（与推送帧里 account 主题的 created_time 是字符串不同）
//
//     {"data":{"state":"normal","equity":"0","details":[{"currency":"USDT","equity":"10.36",
//     "available":"10.36","withdraw_available":"10.36","isolated_available":"0",
//     "created_time":1742818758875,"updated_time":1789377880722,...}],"version":16575142}}
type GetAccountBalanceResponse struct {
	Code int `json:"code"`

	Data struct {
		State                 string `json:"state"`
		Equity                string `json:"equity"`
		InitialMargin         string `json:"initial_margin"`
		MaintenanceMargin     string `json:"maintenance_margin"`
		MaintenanceMarginRate string `json:"maintenance_margin_rate"`
		ProfitUnreal          string `json:"profit_unreal"`
		AvailableMargin       string `json:"available_margin"`
		VoucherValue          string `json:"voucher_value"`
		Version               int64  `json:"version"`
		CreatedTime           int64  `json:"created_time"`
		UpdatedTime           int64  `json:"updated_time"`
		Details               []struct {
			Currency              string `json:"currency"`
			Equity                string `json:"equity"`
			IsolatedEquity        string `json:"isolated_equity"`
			IsolatedAvailable     string `json:"isolated_available"`
			Available             string `json:"available"`
			WithdrawAvailable     string `json:"withdraw_available"`
			ProfitUnreal          string `json:"profit_unreal"`
			IsolatedProfitUnreal  string `json:"isolated_profit_unreal"`
			InitialMargin         string `json:"initial_margin"`
			MaintenanceMargin     string `json:"maintenance_margin"`
			MaintenanceMarginRate string `json:"maintenance_margin_rate"`
			InitialMarginRate     string `json:"initial_margin_rate"`
			Voucher               string `json:"voucher"`
			VoucherValue          string `json:"voucher_value"`
			CreatedTime           int64  `json:"created_time"`
			UpdatedTime           int64  `json:"updated_time"`
		} `json:"details"`
	} `json:"data,omitempty"`

	Message string `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
