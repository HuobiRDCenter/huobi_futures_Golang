package account

type GetAccountBalanceResponse struct {
	Code string `json:"code"`

	Data []struct {
		State                 string `json:"state"`
		Equity                string `json:"equity"`
		InitialMargin         string `json:"initial_margin"`
		MaintenanceMargin     string `json:"maintenance_margin"`
		MaintenanceMarginRate string `json:"maintenance_margin_rate"`
		ProfitUnreal          string `json:"profit_unreal"`
		AvailableMargin       string `json:"available_margin"`
		CreatedTime           string `json:"created_time"`
		UpdatedTime           string `json:"updated_time"`
		Details               []struct {
			Currency              string `json:"currency"`
			Equity                string `json:"equity"`
			IsolatedEquity        string `json:"isolated_equity"`
			Available             string `json:"available"`
			ProfitUnreal          string `json:"profit_unreal"`
			InitialMargin         string `json:"initial_margin"`
			MaintenanceMargin     string `json:"maintenance_margin"`
			MaintenanceMarginRate string `json:"maintenance_margin_rate"`
			InitialMarginRate     string `json:"initial_margin_rate"`
			CreatedTime           string `json:"created_time"`
			UpdatedTime           string `json:"updated_time"`
		} `json:"details"`
	} `json:"data,omitempty"`

	Message int `json:"message,omitempty"`

	Ts int64 `json:"ts"`
}
